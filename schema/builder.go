package schema

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fastschema/fastschema/pkg/utils"
)

// Builder holds the schema of the database.
type Builder struct {
	dir       string
	schemas   map[string]*Schema
	relations []*Relation
}

func GetSchemasFromDir(dir string) (map[string]*Schema, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, err
	}

	jsonFiles, err := filepath.Glob(path.Join(dir, "*.json"))

	if err != nil {
		return nil, err
	}

	schemas := make(map[string]*Schema)
	coreSchemas := []string{UserSchema, PermissionSchema, RoleSchema, MediaSchema}

	for _, coreSchema := range coreSchemas {
		schema, err := NewSchemaFromJSON(coreSchema)
		if err != nil {
			return nil, err
		}

		schemas[schema.Name] = schema
	}

	for _, jsonFile := range jsonFiles {
		schema, err := NewSchemaFromJSONFile(jsonFile)

		if err != nil {
			return nil, err
		}

		if _, ok := schemas[schema.Name]; ok {
			schemas[schema.Name].Fields = append(schemas[schema.Name].Fields, schema.Fields...)
		} else {
			schemas[schema.Name] = schema
		}
	}

	return schemas, nil
}

// NewBuilderFromSchemas creates a new schema from a map of schemas.
func NewBuilderFromSchemas(dir string, schemas map[string]*Schema) (*Builder, error) {
	b := &Builder{dir: dir, schemas: map[string]*Schema{}}

	for _, schema := range schemas {
		if err := schema.Init(false); err != nil {
			return nil, err
		}

		b.schemas[schema.Name] = schema
	}

	return b, b.Init()
}

// NewBuilderFromDir creates a new schema builder from a directory.
func NewBuilderFromDir(dir string) (*Builder, error) {
	schemas, err := GetSchemasFromDir(dir)
	if err != nil {
		return nil, err
	}

	return NewBuilderFromSchemas(dir, schemas)
}

// Clone clones the builder.
func (b *Builder) Clone() *Builder {
	clone := &Builder{
		dir:       b.dir,
		schemas:   map[string]*Schema{},
		relations: []*Relation{},
	}

	for _, schema := range b.schemas {
		clone.schemas[schema.Name] = schema.Clone()
	}

	for _, relation := range b.relations {
		clone.relations = append(clone.relations, relation.Clone())
	}

	return clone
}

// SaveToDir saves all the schemas to a directory.
func (b *Builder) SaveToDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	for _, schema := range b.schemas {
		schemaFile := path.Join(dir, fmt.Sprintf("%s.json", schema.Name))
		if err := schema.SaveToFile(schemaFile); err != nil {
			return err
		}
	}

	return nil
}

// Dir returns the directory of the builder.
// If dirs is not empty, it will set the dir to the first element of dirs.
func (b *Builder) Dir(dirs ...string) string {
	if len(dirs) > 0 {
		b.dir = dirs[0]
	}

	return b.dir
}

// Init initializes the builder.
func (b *Builder) Init() (err error) {
	if b.schemas == nil {
		b.schemas = map[string]*Schema{}
	}

	if err = b.CreateRelations(); err != nil {
		return err
	}

	if err = b.CreateFKs(); err != nil {
		return err
	}

	return nil
}

// SchemaFile returns the json file path of a schema
func (b *Builder) SchemaFile(name string) string {
	return path.Join(b.dir, fmt.Sprintf("%s.json", name))
}

// Schemas returns all schemas
func (b *Builder) Schemas() []*Schema {
	schemas := make([]*Schema, 0)
	for _, schema := range b.schemas {
		schemas = append(schemas, schema)
	}

	return schemas
}

// AddSchema adds a new schema
func (b *Builder) AddSchema(schema *Schema) {
	if b.schemas == nil {
		b.schemas = map[string]*Schema{}
	}

	b.schemas[schema.Name] = schema
}

// ReplaceSchema replaces a schema
func (b *Builder) ReplaceSchema(name string, schema *Schema) {
	b.schemas[name] = schema
}

// Relations returns all relations
func (b *Builder) Relations() []*Relation {
	return b.relations
}

// CreateRelations creates all relations between nodes
func (b *Builder) CreateRelations() (err error) {
	for _, s := range b.schemas {
		for _, field := range s.Fields {
			if !field.Type.IsRelationType() {
				continue
			}

			relationSchema, err := b.Schema(field.Relation.TargetSchemaName)
			if err != nil {
				return NewRelationNodeError(s, field)
			}

			b.relations = append(b.relations, field.Relation.Init(s, relationSchema, field))
		}
	}

	for _, r := range b.relations {
		if r.Type == M2M {
			currentSchema, err := b.Schema(r.SchemaName)
			if err != nil {
				return err
			}

			junctionSchema, exists, err := b.CreateM2mJunctionSchema(currentSchema, r)
			if err != nil {
				return err
			}

			r.JunctionSchema = junctionSchema

			if !exists {
				b.AddSchema(junctionSchema)
			}
		}

		if r.BackRef == nil {
			r.BackRef = b.Relation(r.GetBackRefName())
			if r.BackRef == nil {
				return NewRelationBackRefError(r)
			}
		}
	}

	return nil
}

// CreateFKs creates all foreign keys for relations
func (b *Builder) CreateFKs() error {
	for _, relation := range b.relations {
		schema, err := b.Schema(relation.SchemaName)
		if err != nil {
			return err
		}

		// O2O and O2M relations
		if relation.Type.IsO2O() || relation.Type.IsO2M() {
			relation.CreateFKFields()

			if relation.FKFields != nil {
				schema.Fields = utils.SliceInsertBeforeElement(schema.Fields, relation.FKFields[0], func(f *Field) bool {
					return f.Name == FieldCreatedAt
				})
				schema.DBColumns = utils.SliceInsertBeforeElement(schema.DBColumns, relation.GetTargetFKColumn(), func(c string) bool {
					return c == FieldCreatedAt
				})
			}
		}
	}

	return nil
}

// Schema returns a node by it's name
func (b *Builder) Schema(name string) (*Schema, error) {
	for _, schema := range b.schemas {
		if schema.Name == name {
			return schema, nil
		}
	}

	return nil, fmt.Errorf("schema %s not found", name)
}

// Relation returns a relation by it's name
func (b *Builder) Relation(name string) *Relation {
	for _, relation := range b.relations {
		if relation.Name == name {
			return relation
		}
	}

	return nil
}

func (b *Builder) CreateM2mJunctionSchema(currentSchema *Schema, r *Relation) (*Schema, bool, error) {
	if r == nil || !r.Type.IsM2M() {
		return nil, false, fmt.Errorf("field %s is not a m2m relation", r.Name)
	}

	targetSchema, err := b.Schema(r.TargetSchemaName)
	if err != nil {
		return nil, false, err
	}

	// firstFKName hold the relation information for the current schema
	// secondFKName hold the relation information for the target schema
	// If the relation is bidi, use the schema name as the first fk name to avoid conflicts
	firstFKName := utils.If(r.IsBidi(), r.SchemaName, r.FieldName)
	secondFKName := r.TargetFieldName

	// The firstFKName is connected to the target schema
	// The secondFKName is connected to the current schema
	fKColumnNames := []string{firstFKName, secondFKName}
	r.RelationSchemas = []*Schema{targetSchema, currentSchema}
	r.FKColumns = &RelationFKColumns{
		CurrentColumn: secondFKName,
		TargetColumn:  firstFKName,
	}

	tableNameParts := []string{firstFKName, secondFKName}
	sort.Strings(tableNameParts)
	r.JunctionTable = strings.Join(tableNameParts, "_")

	// If the junction schema already exists, skip creating it
	junctionSchema, _ := b.Schema(r.JunctionTable)
	if junctionSchema != nil {
		return junctionSchema, true, nil
	}

	junctionSchema = &Schema{
		Name:             r.JunctionTable,
		Namespace:        r.JunctionTable,
		LabelFieldName:   fKColumnNames[0],
		IsJunctionSchema: true,
		IsSystemSchema:   true,
		Fields: []*Field{
			CreateUint64Field(fKColumnNames[0]),
			CreateUint64Field(fKColumnNames[1]),
		},
	}

	if err := junctionSchema.Init(true); err != nil {
		return nil, false, err
	}

	return junctionSchema, false, nil
}
