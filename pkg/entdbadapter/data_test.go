package entdbadapter

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"entgo.io/ent/dialect"
	dialectSql "entgo.io/ent/dialect/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/fastschema/fastschema/app"
	"github.com/fastschema/fastschema/pkg/utils"
	"github.com/fastschema/fastschema/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockTestCreateData struct {
	Name        string
	Schema      string
	InputJSON   string
	Run         func(model app.Model, entity *schema.Entity) error
	Expect      func(sqlmock.Sqlmock)
	ExpectError string
	Transaction bool
}

type MockTestUpdateData struct {
	Name         string
	Schema       string
	InputJSON    string
	Run          func(model app.Model, entity *schema.Entity) (int, error)
	Expect       func(sqlmock.Sqlmock)
	Predicates   []*app.Predicate
	WantErr      bool
	WantAffected int
	Transaction  bool
}

type MockTestDeleteData struct {
	Name         string
	Schema       string
	Run          func(model app.Model) (int, error)
	Expect       func(sqlmock.Sqlmock)
	Predicates   []*app.Predicate
	WantErr      bool
	WantAffected int
	Transaction  bool
}

type MockTestCountData struct {
	Name   string
	Schema string
	Filter string
	Column string
	Unique bool
	Expect func(sqlmock.Sqlmock)
	Run    func(
		model app.Model,
		predicates []*app.Predicate,
		unique bool,
		column string,
	) (int, error)
	ExpectError string
	ExpectCount int
}
type MockTestQueryData struct {
	Name    string
	Schema  string
	Filter  string
	Limit   uint
	Offset  uint
	Columns []string
	Order   []string
	Expect  func(sqlmock.Sqlmock)
	Run     func(
		model app.Model,
		predicates []*app.Predicate,
		limit, offset uint,
		order []string,
		columns ...string,
	) ([]*schema.Entity, error)
	ExpectError    string
	ExpectEntities []*schema.Entity
}

func createSchemaBuilderFromDir(schemaDir string) *schema.Builder {
	var err error
	var sb *schema.Builder

	if sb, err = schema.NewBuilderFromDir(schemaDir); err != nil {
		panic(err)
	}

	return sb
}

func createSchemaBuilder() *schema.Builder {
	return createSchemaBuilderFromDir("../../tests/data/schemas")
}

func createMockAdapter(t *testing.T) EntAdapter {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)

	tmpDir, err := os.MkdirTemp("", "migrations")
	assert.NoError(t, err)

	sb := createSchemaBuilder()
	client := utils.Must(NewEntClient(&app.DBConfig{
		Driver:       "sqlmock",
		MigrationDir: tmpDir,
	}, sb, dialectSql.OpenDB(dialect.MySQL, db)))

	adapter, ok := client.(EntAdapter)
	assert.True(t, ok)

	return adapter
}

var testUserSchemaJSON = `{
	"name": "user",
	"namespace": "users",
	"label_field": "name",
	"fields": [
		{
			"name": "name",
			"label": "Name",
			"type": "string",
			"sortable": true
		},
		{
			"name": "json_field",
			"label": "JSON Field",
			"type": "json"
		},
		{
			"name": "bytes_field",
			"label": "Bytes Field",
			"type": "bytes"
		},
		{
			"name": "bool_field",
			"label": "Bool Field",
			"type": "bool"
		},
		{
			"name": "float32_field",
			"label": "Float32 Field",
			"type": "float32"
		},
		{
			"name": "float64_field",
			"label": "Float64 Field",
			"type": "float64"
		},
		{
			"name": "int8_field",
			"label": "int8 Field",
			"type": "int8"
		},
		{
			"name": "int16_field",
			"label": "int16 Field",
			"type": "int16"
		},
		{
			"name": "int32_field",
			"label": "int32 Field",
			"type": "int32"
		},
		{
			"name": "int_field",
			"label": "int Field",
			"type": "int"
		},
		{
			"name": "int64_field",
			"label": "int64 Field",
			"type": "int64"
		},
		{
			"name": "uint8_field",
			"label": "uint8 Field",
			"type": "uint8"
		},
		{
			"label": "uint16 Field",
			"name": "uint16_field",
			"type": "uint16"
		},
		{
			"label": "uint32 Field",
			"name": "uint32_field",
			"type": "uint32"
		},
		{
			"label": "uint Field",
			"name": "uint_field",
			"type": "uint"
		},
		{
			"label": "uint64 Field",
			"name": "uint64_field",
			"type": "uint64"
		},
		{
			"label": "time Field",
			"name": "time_field",
			"type": "time"
		},
		{
			"label": "uuid Field",
			"name": "uuid_field",
			"type": "uuid"
		},
		{
			"label": "enum Field",
			"name": "enum_field",
			"type": "enum",
			"enums": [
				{
					"label": "Enum1",
					"value": "enum1"
				},
				{
					"label": "Enum2",
					"value": "enum2"
				}
			]
		},
		{
			"label": "string Field",
			"name": "string_field",
			"type": "string"
		},
		{
			"label": "text Field",
			"name": "text_field",
			"type": "text"
		}
	]
}`

func NewMockExpectClient(
	createMockClient func(db *sql.DB) app.DBClient,
	s *schema.Builder,
	beforeCreateClient func(m sqlmock.Sqlmock),
	expectTransaction bool,
) (app.DBClient, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	if expectTransaction {
		mock.ExpectBegin()
	}

	if beforeCreateClient != nil {
		beforeCreateClient(mock)
	}

	if expectTransaction {
		mock.ExpectCommit()
	}

	driver := createMockClient(db)

	return driver, nil
}

func MockRunCreateTests(
	createMockClient func(db *sql.DB) app.DBClient,
	sb *schema.Builder,
	t *testing.T,
	tests []MockTestCreateData,
) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			fmt.Printf("Running test: %s\n", tt.Name)
			entity, err := schema.NewEntityFromJSON(tt.InputJSON)
			require.NoError(t, err)

			client, err := NewMockExpectClient(createMockClient, sb, tt.Expect, tt.Transaction)
			require.NoError(t, err)

			model, err := client.Model(tt.Schema)
			require.NoError(t, err)

			runFn := tt.Run
			if runFn == nil {
				runFn = func(model app.Model, entity *schema.Entity) error {
					_, err := model.Mutation().Create(entity)
					return err
				}
			}

			err = runFn(model, entity)
			if err != nil {
				assert.Equal(t, tt.ExpectError, err.Error())
			}

			fmt.Printf("\n\n\n")
		})
	}
}

func MockRunUpdateTests(
	createMockClient func(db *sql.DB) app.DBClient,
	sb *schema.Builder,
	t *testing.T,
	tests []MockTestUpdateData,
	extended ...bool,
) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			fmt.Printf("Running test: %s\n", tt.Name)
			client, err := NewMockExpectClient(createMockClient, sb, tt.Expect, tt.Transaction)
			require.NoError(t, err)
			entity, err := schema.NewEntityFromJSON(tt.InputJSON)
			require.NoError(t, err)

			model, err := client.Model(tt.Schema)
			require.NoError(t, err)
			runFn := tt.Run
			if runFn == nil {
				runFn = func(model app.Model, entity *schema.Entity) (int, error) {
					mut := model.Mutation()
					if len(tt.Predicates) > 0 {
						mut = mut.Where(tt.Predicates...)
					}
					return mut.Update(entity)
				}
			}

			affected, err := runFn(model, entity)
			require.Equal(t, tt.WantErr, err != nil, err)
			if len(extended) > 0 && extended[0] {
				require.Equal(t, tt.WantAffected, affected)
			}
			fmt.Printf("\n\n\n")
		})
	}
}

func MockRunDeleteTests(
	createMockClient func(db *sql.DB) app.DBClient,
	sb *schema.Builder,
	t *testing.T,
	tests []MockTestDeleteData,
	extended ...bool,
) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			fmt.Printf("Running test: %s\n", tt.Name)
			client, err := NewMockExpectClient(createMockClient, sb, tt.Expect, tt.Transaction)
			require.NoError(t, err)

			model, err := client.Model(tt.Schema)
			require.NoError(t, err)
			runFn := tt.Run
			if runFn == nil {
				runFn = func(model app.Model) (int, error) {
					mut := model.Mutation()
					if len(tt.Predicates) > 0 {
						mut = mut.Where(tt.Predicates...)
					}
					return mut.Delete()
				}
			}

			affected, err := runFn(model)
			require.Equal(t, tt.WantErr, err != nil, err)
			if len(extended) > 0 && extended[0] {
				require.Equal(t, tt.WantAffected, affected)
			}
			fmt.Printf("\n\n\n")
		})
	}
}

func MockRunCountTests(
	createMockClient func(db *sql.DB) app.DBClient,
	sb *schema.Builder,
	t *testing.T,
	tests []MockTestCountData,
) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			fmt.Printf("Running test: %s\n", tt.Name)
			client, err := NewMockExpectClient(createMockClient, sb, tt.Expect, false)
			require.NoError(t, err)

			model, err := client.Model(tt.Schema)
			require.NoError(t, err)

			runFn := tt.Run
			if runFn == nil {
				runFn = func(
					model app.Model,
					predicates []*app.Predicate,
					unique bool,
					column string,
				) (int, error) {
					query := model.Query()
					if len(predicates) > 0 {
						query = query.Where(predicates...)
					}
					countOpts := &app.CountOption{
						Unique: unique,
						Column: column,
					}

					return query.Count(countOpts)
				}
			}

			var predicates []*app.Predicate
			if tt.Filter != "" {
				predicates, err = app.CreatePredicatesFromFilterObject(sb, model.Schema(), tt.Filter)
				require.NoError(t, err)
			}

			count, err := runFn(model, predicates, tt.Unique, tt.Column)
			if tt.ExpectError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tt.ExpectCount, count)
			} else {
				assert.Error(t, err)
				require.Equal(t, tt.ExpectError, err.Error(), err)
			}

			fmt.Printf("\n\n\n")
		})
	}
}

func MockDefaultQueryRunFn(
	model app.Model,
	predicates []*app.Predicate,
	limit, offset uint,
	order []string,
	columns ...string,
) ([]*schema.Entity, error) {
	query := model.Query()
	if len(predicates) > 0 {
		query = query.Where(predicates...)
	}

	if len(order) > 0 {
		query = query.Order(order...)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	if offset > 0 {
		query = query.Offset(offset)
	}

	if len(columns) > 0 {
		query = query.Select(columns...)
	}

	return query.Get()
}

func MockRunQueryTests(
	createMockClient func(db *sql.DB) app.DBClient,
	sb *schema.Builder,
	t *testing.T,
	tests []MockTestQueryData,
	extended ...bool,
) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			fmt.Printf("Running test: %s\n", tt.Name)
			client, err := NewMockExpectClient(createMockClient, sb, tt.Expect, false)
			require.NoError(t, err)

			model, err := client.Model(tt.Schema)
			require.NoError(t, err)

			runFn := tt.Run
			if runFn == nil {
				runFn = MockDefaultQueryRunFn
			}

			var predicates []*app.Predicate
			if tt.Filter != "" {
				predicates, err = app.CreatePredicatesFromFilterObject(sb, model.Schema(), tt.Filter)
				require.NoError(t, err)
			}

			entities, err := runFn(model, predicates, tt.Limit, tt.Offset, tt.Order, tt.Columns...)
			if tt.ExpectError == "" {
				assert.NoError(t, err)
				expectEntititiesJSONs := make([]string, len(tt.ExpectEntities))
				entitiesJSONs := make([]string, len(entities))
				for i, e := range tt.ExpectEntities {
					expectEntitityJSON, err := e.ToJSON()
					require.NoError(t, err)
					expectEntititiesJSONs[i] = expectEntitityJSON
				}

				for i, e := range entities {
					expectJSON, err := e.ToJSON()
					require.NoError(t, err)
					entitiesJSONs[i] = expectJSON
				}

				if !assert.Equal(t, expectEntititiesJSONs, entitiesJSONs) {
					fmt.Println("------------WANT-----------")
					for _, we := range expectEntititiesJSONs {
						fmt.Println(we)
					}
					fmt.Println("------------GOT-----------")
					for _, e := range entitiesJSONs {
						fmt.Println(e)
					}
				}
			} else {
				assert.Error(t, err)
				require.Equal(t, tt.ExpectError, err.Error(), err)
			}

			fmt.Printf("\n\n\n")
		})
	}
}
