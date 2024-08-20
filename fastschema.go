package fastschema

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/fastschema/fastschema/app"
	"github.com/fastschema/fastschema/pkg/entdbadapter"
	"github.com/fastschema/fastschema/pkg/errors"
	"github.com/fastschema/fastschema/pkg/rclonefs"
	"github.com/fastschema/fastschema/pkg/restresolver"
	"github.com/fastschema/fastschema/pkg/utils"
	"github.com/fastschema/fastschema/pkg/zaplogger"
	"github.com/fastschema/fastschema/schema"
	cs "github.com/fastschema/fastschema/services/content"
	ms "github.com/fastschema/fastschema/services/media"
	rs "github.com/fastschema/fastschema/services/role"
	ss "github.com/fastschema/fastschema/services/schema"
	ts "github.com/fastschema/fastschema/services/tool"
	us "github.com/fastschema/fastschema/services/user"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

//go:embed all:dash/*
var embedDashStatic embed.FS
var CWD string

func init() {
	var err error
	if CWD, err = os.Getwd(); err != nil {
		log.Fatal(err)
	}
}

type AppConfig struct {
	Dir               string
	AppKey            string
	Port              string
	BaseURL           string
	DashURL           string
	APIBaseName       string
	DashBaseName      string
	Logger            app.Logger
	DB                app.DBClient
	StorageConfig     *app.StorageConfig
	HideResourcesInfo bool
}

func (ac *AppConfig) Clone() *AppConfig {
	return &AppConfig{
		Dir:               ac.Dir,
		AppKey:            ac.AppKey,
		Port:              ac.Port,
		BaseURL:           ac.BaseURL,
		DashURL:           ac.DashURL,
		APIBaseName:       ac.APIBaseName,
		DashBaseName:      ac.DashBaseName,
		Logger:            ac.Logger,
		DB:                ac.DB,
		StorageConfig:     ac.StorageConfig.Clone(),
		HideResourcesInfo: ac.HideResourcesInfo,
	}
}

type App struct {
	config          *AppConfig
	dir             string
	envFile         string
	dataDir         string
	logDir          string
	publicDir       string
	schemasDir      string
	migrationDir    string
	schemaBuilder   *schema.Builder
	resources       *app.ResourcesManager
	api             *app.Resource
	hooks           *app.Hooks
	roles           []*app.Role
	disks           []app.Disk
	defaultDisk     app.Disk
	setupToken      string
	startupMessages []string
	statics         []*app.StaticFs
}

func New(config *AppConfig) (_ *App, err error) {
	a := &App{
		config: config.Clone(),
		disks:  []app.Disk{},
		roles:  []*app.Role{},
		hooks:  &app.Hooks{},
	}

	if err := a.prepareConfig(); err != nil {
		return nil, err
	}

	if err = a.init(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) init() (err error) {
	if err = a.getDefaultDisk(); err != nil {
		return err
	}

	if err = a.getDefaultLogger(); err != nil {
		return err
	}

	if err = a.createSchemaBuilder(); err != nil {
		return err
	}

	if err = a.createResources(); err != nil {
		return err
	}

	if err = a.getDefaultDBClient(); err != nil {
		return err
	}

	// if a local disk has a public path, then add it to the statics
	for _, disk := range a.disks {
		publicPath := disk.LocalPublicPath()

		if publicPath != "" {
			a.startupMessages = append(
				a.startupMessages,
				fmt.Sprintf("Serving files from disk [%s:%s] at %s",
					disk.Name(), publicPath, disk.Root()),
			)

			a.statics = append(a.statics, &app.StaticFs{
				BasePath: publicPath,
				Root:     http.Dir(disk.Root()),
			})
		}
	}

	setupToken, err := a.SetupToken()
	if err != nil {
		return err
	}

	if setupToken != "" {
		type setupData struct {
			Token    string `json:"token"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		a.api.Add(app.NewResource("setup", func(c app.Context, setupData *setupData) (bool, error) {
			if setupToken == "" {
				return false, errors.BadRequest("Setup token is not available")
			}

			if setupData == nil {
				return false, errors.BadRequest("Invalid setup data")
			}

			if setupData.Token != setupToken {
				return false, errors.Forbidden("Invalid setup token")
			}

			if err := ts.Setup(
				a.DB(),
				a.Logger(),
				setupData.Username, setupData.Email, setupData.Password,
			); err != nil {
				return false, err
			}

			if err := a.UpdateCache(); err != nil {
				return false, err
			}

			setupToken = ""
			a.setupToken = ""

			return true, nil
		}, app.Meta{app.POST: "/setup"}, true))

		setupURL := fmt.Sprintf(
			"%s/setup/?token=%s\033[0m",
			a.config.DashURL,
			setupToken,
		)

		a.startupMessages = append(a.startupMessages, fmt.Sprintf(
			"Visit the following URL to setup the app: %s",
			setupURL,
		))
	}

	a.statics = append(a.statics, &app.StaticFs{
		BasePath:   "/" + a.config.DashBaseName,
		Root:       http.FS(embedDashStatic),
		PathPrefix: "dash",
	})

	return nil
}

func (a *App) Config() *AppConfig {
	return a.config
}

func (a *App) Key() string {
	return a.config.AppKey
}

func (a *App) Dir() string {
	return a.dir
}

func (a *App) Reload(migration *app.Migration) (err error) {
	if a.DB() != nil {
		if err = a.DB().Close(); err != nil {
			return err
		}
	}

	if err = a.createSchemaBuilder(); err != nil {
		return err
	}

	newDB, err := a.DB().Reload(a.schemaBuilder, migration)
	if err != nil {
		return err
	}

	a.config.DB = newDB

	return nil
}

// UpdateCache updates the application cache.
// It fetches all roles from the database and stores them in the cache.
func (a *App) UpdateCache() error {
	a.roles = []*app.Role{}
	roleModel, err := a.DB().Model("role")
	if err != nil {
		return err
	}

	roles, err := roleModel.Query().Select(
		"id",
		"name",
		"description",
		"root",
		"permissions",
		schema.FieldCreatedAt,
		schema.FieldUpdatedAt,
		schema.FieldDeletedAt,
	).Get(context.Background())

	if err != nil {
		return err
	}

	for _, r := range roles {
		role := app.EntityToRole(r)
		a.roles = append(a.roles, role)
	}

	return nil
}

func (a *App) Start() error {
	addr := fmt.Sprintf(":%s", a.config.Port)
	if err := a.resources.Init(); err != nil {
		return err
	}

	if !a.config.HideResourcesInfo {
		a.resources.Print()
	}
	restResolver := restresolver.NewRestResolver(a.resources, a.Logger(), a.statics...)

	fmt.Printf("\n")
	for _, msg := range a.startupMessages {
		color.Green("> %s", msg)
	}
	fmt.Printf("\n")

	return restResolver.Start(addr)
}

func (a *App) SchemaBuilder() *schema.Builder {
	return a.schemaBuilder
}

func (a *App) DB() app.DBClient {
	return a.config.DB
}

func (a *App) API() *app.Resource {
	return a.api
}

func (a *App) Logger() app.Logger {
	return a.config.Logger
}

func (a *App) Resources() *app.ResourcesManager {
	return a.resources
}

func (a *App) Roles() []*app.Role {
	return a.roles
}

func (a *App) Hooks() *app.Hooks {
	return a.hooks
}

func (a *App) Disks() []app.Disk {
	return a.disks
}

func (a *App) Disk(names ...string) app.Disk {
	if len(names) == 0 {
		return a.defaultDisk
	}

	for _, disk := range a.disks {
		if disk.Name() == names[0] {
			return disk
		}
	}

	return nil
}

func (a *App) AddResource(resource *app.Resource) {
	a.resources.Add(resource)
}

func (a *App) AddMiddlewares(middlewares ...app.Middleware) {
	a.resources.Middlewares = append(a.resources.Middlewares, middlewares...)
}

func (a *App) OnPreResolve(middlewares ...app.Middleware) {
	a.hooks.PreResolve = append(a.hooks.PreResolve, middlewares...)
}

func (a *App) OnPostResolve(middlewares ...app.Middleware) {
	a.hooks.PostResolve = append(a.hooks.PostResolve, middlewares...)
}

func (a *App) OnPostDBGet(hooks ...app.PostDBGetHook) {
	a.hooks.PostDBGet = append(a.hooks.PostDBGet, hooks...)
}

func (a *App) SetupToken() (string, error) {
	// If there is no roles and users, then the app is not setup
	// we need to setup the app.
	// Generate a random token and return it to enable the setup.
	// If there are roles or users, then the app is already setup.
	// Return an empty string to disable the setup.
	needSetup, err := a.needSetup()
	if err != nil {
		return "", err
	}

	if !needSetup {
		return "", nil
	}

	if a.setupToken == "" {
		a.setupToken = utils.RandomString(32)
	}

	return a.setupToken, nil
}

func (a *App) needSetup() (bool, error) {
	// If there is no roles and users, then the app is not setup
	// we need to setup the app.
	var err error
	var userCount int
	var roleCount int
	ctx := context.Background()
	countOption := &app.CountOption{
		Column: "id",
		Unique: true,
	}

	userModel, err := a.DB().Model("user")
	if err != nil {
		return false, err
	}

	roleModel, err := a.DB().Model("role")
	if err != nil {
		return false, err
	}

	if userCount, err = userModel.Query().Count(countOption, ctx); err != nil {
		return false, err
	}

	if roleCount, err = roleModel.Query().Count(countOption, ctx); err != nil {
		return false, err
	}

	return userCount == 0 && roleCount == 0, nil
}

func (a *App) getDefaultDBClient() (err error) {
	if a.DB() != nil {
		return nil
	}

	dbConfig := &app.DBConfig{
		Driver:       utils.Env("DB_DRIVER", "sqlite"),
		Name:         utils.Env("DB_NAME"),
		User:         utils.Env("DB_USER"),
		Pass:         utils.Env("DB_PASS"),
		Host:         utils.Env("DB_HOST"),
		Port:         utils.Env("DB_PORT"),
		LogQueries:   utils.Env("DB_LOGGING", "false") == "true",
		Logger:       a.Logger(),
		MigrationDir: a.migrationDir,
		Hooks: func() *app.Hooks {
			return a.hooks
		},
	}

	// If driver is sqlite and the DB_NAME (file path) is not set,
	// Set the DB_NAME to the default sqlite db file path.
	if dbConfig.Driver == "sqlite" && dbConfig.Name == "" {
		dbConfig.Name = path.Join(a.dataDir, "fastschema.db")
		a.startupMessages = append(
			a.startupMessages,
			fmt.Sprintf("Using default sqlite db file: %s", dbConfig.Name),
		)
	}

	if a.config.DB, err = entdbadapter.NewClient(dbConfig, a.schemaBuilder); err != nil {
		return err
	}

	if err := a.UpdateCache(); err != nil {
		return err
	}

	return nil
}

func (a *App) getDefaultLogger() (err error) {
	if a.config.Logger == nil {
		a.config.Logger, err = zaplogger.NewZapLogger(&zaplogger.ZapConfig{
			Development: true,
			LogFile:     path.Join(a.dir, "data/logs/app.log"),
		})
	}

	return err
}

func (a *App) getDefaultDisk() (err error) {
	if a.config.StorageConfig == nil {
		a.config.StorageConfig = &app.StorageConfig{}
	}

	defaultDiskName := a.config.StorageConfig.DefaultDisk
	if defaultDiskName == "" {
		defaultDiskName = utils.Env("STORAGE_DEFAULT_DISK")
	}

	storageDisksConfig := a.config.StorageConfig.DisksConfig
	if utils.Env("STORAGE_DISKS") != "" && storageDisksConfig == nil {
		if err := json.Unmarshal([]byte(utils.Env("STORAGE_DISKS")), &storageDisksConfig); err != nil {
			return err
		}
	}

	// if threre is no disk config, add a default disk
	if storageDisksConfig == nil {
		storageDisksConfig = []*app.DiskConfig{{
			Name:       "local_public",
			Driver:     "local",
			PublicPath: "/files",
			BaseURL:    fmt.Sprintf("%s/files", a.config.BaseURL),
			Root:       a.publicDir,
		}}
	}

	if a.disks, err = rclonefs.NewFromConfig(storageDisksConfig, a.dataDir); err != nil {
		return err
	}

	foundDefaultDisk := false
	for _, disk := range a.disks {
		if disk.Name() == defaultDiskName {
			a.defaultDisk = disk
			foundDefaultDisk = true
			break
		}
	}

	if defaultDiskName != "" && !foundDefaultDisk {
		return fmt.Errorf("default disk [%s] not found", defaultDiskName)
	}

	if a.defaultDisk == nil && len(a.disks) > 0 {
		a.defaultDisk = a.disks[0]
	}

	return nil
}

func (a *App) createSchemaBuilder() (err error) {
	if a.schemaBuilder, err = schema.NewBuilderFromDir(a.schemasDir); err != nil {
		return err
	}

	return nil
}

func (a *App) createResources() error {
	userService := us.New(a)
	roleService := rs.New(a)
	mediaService := ms.New(a)
	schemaService := ss.New(a)
	contentService := cs.New(a)
	toolService := ts.New(a)

	a.hooks.PostDBGet = append(a.hooks.PostDBGet, mediaService.MediaListHook)
	a.hooks.PreResolve = append(a.hooks.PreResolve, roleService.Authorize)

	a.resources = app.NewResourcesManager()
	a.resources.Middlewares = append(a.resources.Middlewares, roleService.ParseUser)
	a.resources.Hooks = func() *app.Hooks {
		return a.hooks
	}

	a.api = a.resources.Group(a.config.APIBaseName)
	a.api.Group("user").
		Add(app.NewResource("logout", userService.Logout, app.Meta{app.POST: "/logout"}, true)).
		Add(app.NewResource("me", userService.Me, true)).
		Add(app.NewResource("login", userService.Login, app.Meta{app.POST: "/login"}, true))

	a.api.Group("schema").
		Add(app.NewResource("list", schemaService.List, app.Meta{app.GET: ""})).
		Add(app.NewResource("create", schemaService.Create, app.Meta{app.POST: ""})).
		Add(app.NewResource("detail", schemaService.Detail, app.Meta{app.GET: "/:name"})).
		Add(app.NewResource("update", schemaService.Update, app.Meta{app.PUT: "/:name"})).
		Add(app.NewResource("delete", schemaService.Delete, app.Meta{app.DELETE: "/:name"}))

	a.api.Group("content").
		Add(app.NewResource("list", contentService.List, app.Meta{app.GET: "/:schema"})).
		Add(app.NewResource("detail", contentService.Detail, app.Meta{app.GET: "/:schema/:id"})).
		Add(app.NewResource("create", contentService.Create, app.Meta{app.POST: "/:schema"})).
		Add(app.NewResource("update", contentService.Update, app.Meta{app.PUT: "/:schema/:id"})).
		Add(app.NewResource("delete", contentService.Delete, app.Meta{app.DELETE: "/:schema/:id"}))

	a.api.Group("role").
		Add(app.NewResource("list", roleService.List, app.Meta{app.GET: ""})).
		Add(app.NewResource("resources", roleService.ResourcesList, app.Meta{app.GET: "/resources"})).
		Add(app.NewResource("detail", roleService.Detail, app.Meta{app.GET: "/:id"})).
		Add(app.NewResource("create", roleService.Create, app.Meta{app.POST: ""})).
		Add(app.NewResource("update", roleService.Update, app.Meta{app.PUT: "/:id"})).
		Add(app.NewResource("delete", roleService.Delete, app.Meta{app.DELETE: "/:id"}))

	a.api.Group("media").
		Add(app.NewResource("upload", mediaService.Upload, app.Meta{app.POST: "/upload"})).
		Add(app.NewResource("delete", mediaService.Delete, app.Meta{app.DELETE: ""}))

	a.api.Group("tool").
		Add(app.NewResource("stats", toolService.Stats, app.Meta{app.GET: "/stats"}, true))

	return nil
}

func (a *App) getAppDir() {
	defer func() {
		a.startupMessages = append(a.startupMessages, fmt.Sprintf("Using app directory: %s", a.dir))
	}()

	if a.config.Dir == "" {
		a.dir = CWD
		return
	}

	if strings.HasPrefix(a.config.Dir, "/") {
		a.dir = a.config.Dir
		return
	}

	a.dir = path.Join(CWD, a.config.Dir)
}

func (a *App) prepareConfig() (err error) {
	a.getAppDir()
	a.dataDir = path.Join(a.dir, "data")
	a.logDir = path.Join(a.dataDir, "logs")
	a.publicDir = path.Join(a.dataDir, "public")
	a.schemasDir = path.Join(a.dataDir, "schemas")
	a.migrationDir = path.Join(a.dataDir, "migrations")
	envFile := path.Join(a.dataDir, ".env")

	// pretty print a in console

	if err = utils.MkDirs(
		a.logDir,
		a.publicDir,
		a.schemasDir,
		a.migrationDir,
	); err != nil {
		return err
	}

	if utils.IsFileExists(envFile) {
		a.envFile = envFile
		if err := godotenv.Load(envFile); err != nil {
			return err
		}
	}

	if a.config.AppKey == "" {
		a.config.AppKey = utils.Env("APP_KEY")
	}

	if a.config.Port == "" {
		a.config.Port = utils.Env("APP_PORT", "8000")
	}

	if a.config.BaseURL == "" {
		a.config.BaseURL = utils.Env("APP_BASE_URL")
	}

	if a.config.DashURL == "" {
		a.config.DashURL = utils.Env("APP_DASH_URL")
	}

	if a.config.APIBaseName == "" {
		a.config.APIBaseName = utils.Env("APP_API_BASE_NAME", "api")
	}

	if a.config.DashBaseName == "" {
		a.config.DashBaseName = utils.Env("APP_DASH_BASE_NAME", "dash")
	}

	if a.config.BaseURL == "" {
		a.config.BaseURL = fmt.Sprintf("http://localhost:%s", a.config.Port)
	}

	if a.config.DashURL == "" {
		a.config.DashURL = fmt.Sprintf("%s/%s", a.config.BaseURL, a.config.DashBaseName)
	}

	if a.config.AppKey == "" {
		a.config.AppKey = utils.RandomString(32)
		if err := utils.AppendFile(
			envFile,
			fmt.Sprintf("APP_KEY=%s\n", a.config.AppKey),
		); err != nil {
			return err
		}

		a.startupMessages = append(
			a.startupMessages,
			fmt.Sprintf("APP_KEY is not set. A new key is generated and saved to %s", envFile),
		)
	}

	return nil
}
