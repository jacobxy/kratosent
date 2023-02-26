// Code generated by ent, DO NOT EDIT.

package department

import (
	"time"
)

const (
	// Label holds the string label denoting the department type in the database.
	Label = "department"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldParentDepartmentID holds the string denoting the parent_department_id field in the database.
	FieldParentDepartmentID = "parent_department_id"
	// Table holds the table name of the department in the database.
	Table = "departments"
)

// Columns holds all SQL columns for department fields.
var Columns = []string{
	FieldID,
	FieldIsDeleted,
	FieldCreateAt,
	FieldUpdateAt,
	FieldName,
	FieldParentDepartmentID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted int8
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
