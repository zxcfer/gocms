package schema

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const FieldID = "id"
const FieldCreatedAt = "created_at"
const FieldUpdatedAt = "updated_at"
const FieldDeletedAt = "deleted_at"

// FieldEnum define the data struct for an enum field
type FieldEnum struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// FieldRenderer define the renderer of a field
type FieldRenderer struct {
	Class    string         `json:"class,omitempty"`    // renderer class name
	Settings map[string]any `json:"settings,omitempty"` // renderer settings.
}

// FieldDB define the db config for a field
type FieldDB struct {
	Attr      string `json:"attr,omitempty"`      // extra attributes.
	Collation string `json:"collation,omitempty"` // collation type (utf8mb4_unicode_ci, utf8mb4_general_ci)
	Increment bool   `json:"increment,omitempty"` // auto increment
	Key       string `json:"key,omitempty"`       // key definition (PRI, UNI or MUL).
}

func (f *FieldEnum) Clone() *FieldEnum {
	return &FieldEnum{
		Value: f.Value,
		Label: f.Label,
	}
}

func (f *FieldDB) Clone() *FieldDB {
	if f == nil {
		return nil
	}

	return &FieldDB{
		Attr:      f.Attr,
		Collation: f.Collation,
		Increment: f.Increment,
		Key:       f.Key,
	}
}

// FieltType define the data type of a field
type FieldType int

const (
	TypeInvalid FieldType = iota
	TypeBool
	TypeTime
	TypeJSON
	TypeUUID
	TypeBytes
	TypeEnum
	TypeString
	TypeText
	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt
	TypeInt64
	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint
	// TypeUintptr
	TypeUint64
	TypeFloat32
	TypeFloat64
	TypeRelation
	TypeMedia
	endFieldTypes
)

var (
	fieldTypeToStrings = [...]string{
		TypeInvalid: "invalid",
		TypeBool:    "bool",
		TypeTime:    "time",
		TypeJSON:    "json",
		TypeUUID:    "uuid",
		TypeBytes:   "bytes",
		TypeEnum:    "enum",
		TypeString:  "string",
		TypeText:    "text",
		TypeInt:     "int",
		TypeInt8:    "int8",
		TypeInt16:   "int16",
		TypeInt32:   "int32",
		TypeInt64:   "int64",
		TypeUint:    "uint",
		// TypeUintptr:  "uintptr",
		TypeUint8:    "uint8",
		TypeUint16:   "uint16",
		TypeUint32:   "uint32",
		TypeUint64:   "uint64",
		TypeFloat32:  "float32",
		TypeFloat64:  "float64",
		TypeRelation: "relation",
		TypeMedia:    "media",
	}

	stringToFieldTypes = map[string]FieldType{
		"invalid": TypeInvalid,
		"bool":    TypeBool,
		"time":    TypeTime,
		"json":    TypeJSON,
		"uuid":    TypeUUID,
		"bytes":   TypeBytes,
		"enum":    TypeEnum,
		"string":  TypeString,
		"text":    TypeText,
		"int":     TypeInt,
		"int8":    TypeInt8,
		"int16":   TypeInt16,
		"int32":   TypeInt32,
		"int64":   TypeInt64,
		"uint":    TypeUint,
		// "uintptr":  TypeUintptr,
		"uint8":    TypeUint8,
		"uint16":   TypeUint16,
		"uint32":   TypeUint32,
		"uint64":   TypeUint64,
		"float32":  TypeFloat32,
		"float64":  TypeFloat64,
		"relation": TypeRelation,
		"media":    TypeMedia,
	}
)

func (t FieldType) IsRelationType() bool {
	return t == TypeRelation || t == TypeMedia
}

// String returns the string representation of a type.
func (t FieldType) String() string {
	if t < endFieldTypes {
		return fieldTypeToStrings[t]
	}
	return fieldTypeToStrings[TypeInvalid]
}

// Valid reports if the given type if known type.
func (t FieldType) Valid() bool {
	return t > TypeInvalid && t < endFieldTypes
}

// MarshalJSON marshal an enum value to the quoted json string value
func (t FieldType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(fieldTypeToStrings[t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (t *FieldType) UnmarshalJSON(b []byte) error {
	var fieldType string
	if err := json.Unmarshal(b, &fieldType); err != nil {
		return err
	}
	*t = stringToFieldTypes[fieldType] // If the string can't be found, it will be set to the zero value: 'invalid'

	if *t == TypeInvalid {
		return fmt.Errorf("invalid field type %q", fieldType)
	}

	return nil
}

// RelationType define the relation type of a field
type RelationType int

const (
	RelationInvalid RelationType = iota
	O2O
	O2M
	M2M
	endRelationTypes
)

var (
	relationTypeToStrings = [...]string{
		RelationInvalid: "invalid",
		O2O:             "o2o",
		O2M:             "o2m",
		M2M:             "m2m",
	}

	stringToRelationTypes = map[string]RelationType{
		"invalid": RelationInvalid,
		"o2o":     O2O,
		"o2m":     O2M,
		"m2m":     M2M,
	}
)

func (t RelationType) IsO2O() bool {
	return t == O2O
}

func (t RelationType) IsO2M() bool {
	return t == O2M
}

func (t RelationType) IsM2M() bool {
	return t == M2M
}

// String returns the string representation of a type.
func (t RelationType) String() string {
	if t < endRelationTypes {
		return relationTypeToStrings[t]
	}
	return relationTypeToStrings[RelationInvalid]
}

// Valid reports if the given type if known type.
func (t RelationType) Valid() bool {
	return t > RelationInvalid && t < endRelationTypes
}

// MarshalJSON marshal an enum value to the quoted json string value
func (t RelationType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(relationTypeToStrings[t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (t *RelationType) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*t = stringToRelationTypes[j] // If the string can't be found, it will be set to the zero value: 'invalid'
	return nil
}
