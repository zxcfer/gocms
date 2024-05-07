package entdbadapter

import (
	"fmt"
	"testing"

	atlasSchema "ariga.io/atlas/sql/schema"
	entSchema "entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
	"github.com/fastschema/fastschema/app"
	"github.com/fastschema/fastschema/schema"
	"github.com/stretchr/testify/assert"
)

func TestCreateEntColumn(t *testing.T) {
	type args struct {
		name   string
		field  *schema.Field
		column *entSchema.Column
	}

	tests := []args{
		{
			name: "testIDColumn",
			field: &schema.Field{
				Name:   "id",
				Type:   schema.TypeUint64,
				Unique: true,
				DB: &schema.FieldDB{
					Increment: true,
				},
			},
			column: &entSchema.Column{
				Name:      "id",
				Type:      field.TypeUint64,
				Increment: true,
				Unique:    true,
			},
		},
		{
			name: "testTextColumn",
			field: &schema.Field{
				Name: "content",
				Type: schema.TypeText,
				Size: 100,
				DB: &schema.FieldDB{
					Collation: "utf8mb4_unicode_ci",
					Key:       "MUL",
					Attr:      "UNIQUE",
				},
			},
			column: &entSchema.Column{
				Name:      "content",
				Type:      field.TypeString,
				Size:      100,
				Collation: "utf8mb4_unicode_ci",
				Key:       "MUL",
				Attr:      "UNIQUE",
			},
		},
		{
			name: "testNormalColumn",
			field: &schema.Field{
				Name:     "name",
				Type:     schema.TypeString,
				Default:  "test",
				Optional: true,
			},
			column: &entSchema.Column{
				Name:     "name",
				Type:     field.TypeString,
				Default:  "test",
				Nullable: true,
			},
		},
		{
			name: "testEnumColumn",
			field: &schema.Field{
				Name: "status",
				Type: schema.TypeEnum,
				Enums: []*schema.FieldEnum{
					{
						Label: "Active",
						Value: "active",
					},
					{
						Label: "Inactive",
						Value: "inactive",
					},
				},
			},
			column: &entSchema.Column{
				Name:  "status",
				Type:  field.TypeEnum,
				Enums: []string{"active", "inactive"},
			},
		},
		{
			name: "testTimeColumn",
			field: &schema.Field{
				Name: "created_at",
				Type: schema.TypeTime,
			},
			column: &entSchema.Column{
				Name: "created_at",
				Type: field.TypeTime,
				SchemaType: map[string]string{
					"mysql": "datetime",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			column := createEntColumn(tt.field)
			assert.Equal(t, tt.column, column)
		})
	}
}

func TestCreateDBDSN(t *testing.T) {
	config := &app.DBConfig{
		Driver: "mysql",
		User:   "user",
		Pass:   "pass",
		Host:   "localhost",
		Port:   "3306",
		Name:   "database",
	}

	expectedMySQLDSN := "user:pass@tcp(localhost:3306)/database?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true"
	assert.Equal(t, expectedMySQLDSN, CreateDBDSN(config))

	config.Driver = "pgx"
	expectedPGXDSN := "host=localhost port=3306 user=user dbname=database password=pass sslmode=disable"
	assert.Equal(t, expectedPGXDSN, CreateDBDSN(config))

	config.Driver = "sqlite"
	expectedSQLiteDSN := "file:database?cache=shared&_fk=1&_pragma=foreign_keys(1)"
	assert.Equal(t, expectedSQLiteDSN, CreateDBDSN(config))
	config.Name = ":memory:"
	expectedSQLiteMemoryDSN := ":memory:?cache=shared&_fk=1&_pragma=foreign_keys(1)"
	assert.Equal(t, expectedSQLiteMemoryDSN, CreateDBDSN(config))
}

func TestGetEntDialect(t *testing.T) {
	tests := []struct {
		name            string
		config          *app.DBConfig
		expectedDialect string
		expectedError   error
	}{
		{
			name: "Supported driver",
			config: &app.DBConfig{
				Driver: "mysql",
			},
			expectedDialect: "mysql",
			expectedError:   nil,
		},
		{
			name: "Unsupported driver",
			config: &app.DBConfig{
				Driver: "mongodb",
			},
			expectedDialect: "",
			expectedError:   fmt.Errorf("unsupported driver: mongodb"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dialect, err := GetEntDialect(tt.config)
			assert.Equal(t, tt.expectedDialect, dialect)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestCreateRenameColumnsHook(t *testing.T) {
	// Define sample rename tables and rename columns
	renameTables := []*app.RenameItem{
		{From: "old_table", To: "new_table"},
		{From: "another_table", To: "renamed_table"},
	}
	renameColumns := []*app.RenameItem{
		{From: "old_column", To: "new_column"},
		{From: "another_column", To: "renamed_column"},
	}

	// Create a sample current schema
	currentSchema := &atlasSchema.Schema{
		Tables: []*atlasSchema.Table{
			{
				Name: "old_table",
				Columns: []*atlasSchema.Column{
					{Name: "old_column"},
					{Name: "another_column"},
				},
			},
			{
				Name: "another_table",
				Columns: []*atlasSchema.Column{
					{Name: "some_column"},
				},
			},
		},
	}

	// Create a sample desired schema
	desiredSchema := &atlasSchema.Schema{
		Tables: []*atlasSchema.Table{
			{
				Name: "new_table",
				Columns: []*atlasSchema.Column{
					{Name: "new_column"},
				},
			},
			{
				Name: "renamed_table",
				Columns: []*atlasSchema.Column{
					{Name: "some_column"},
					{Name: "renamed_column"},
				},
			},
		},
	}

	// Create the diff hook
	diffHook := createRenameColumnsHook(renameTables, renameColumns)
	expectedChanges := []atlasSchema.Change{
		&atlasSchema.AddTable{
			T: desiredSchema.Tables[0],
		},
		&atlasSchema.RenameColumn{
			From: currentSchema.Tables[0].Columns[0],
			To:   desiredSchema.Tables[0].Columns[0],
		},
		&atlasSchema.RenameColumn{
			From: currentSchema.Tables[0].Columns[1],
			To:   desiredSchema.Tables[1].Columns[1],
		},
	}

	// Create the differ with the diff hook
	var next entSchema.Differ = entSchema.DiffFunc(func(current, desired *atlasSchema.Schema) ([]atlasSchema.Change, error) {
		return expectedChanges, nil
	})

	differ := diffHook(next)

	// Calculate the diff between current and desired schema
	changes, err := differ.Diff(currentSchema, desiredSchema)
	assert.NoError(t, err)

	assert.Equal(t, expectedChanges[1:], changes)
}

func TestNOW(t *testing.T) {
	// Test for MySQL dialect
	mysqlResult := NOW("mysql")
	assert.NotNil(t, mysqlResult)
	// Add assertions for the expected MySQL result

	// Test for PostgreSQL dialect
	pgxResult := NOW("pgx")
	assert.NotNil(t, pgxResult)
	// Add assertions for the expected PostgreSQL result

	// Test for SQLite dialect
	sqliteResult := NOW("sqlite")
	assert.NotNil(t, sqliteResult)
	// Add assertions for the expected SQLite result

	// Test for unsupported dialect
	unsupportedResult := NOW("unsupported")
	assert.NotNil(t, unsupportedResult)
	// Add assertions for the expected unsupported result
}
