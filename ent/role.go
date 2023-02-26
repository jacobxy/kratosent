// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratosent/ent/role"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted int8 `json:"is_deleted,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldID, role.FieldIsDeleted:
			values[i] = new(sql.NullInt64)
		case role.FieldName, role.FieldDescription:
			values[i] = new(sql.NullString)
		case role.FieldCreateAt, role.FieldUpdateAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Role", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = int64(value.Int64)
			}
		case role.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				r.IsDeleted = int8(value.Int64)
			}
		case role.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				r.CreateAt = value.Time
			}
		case role.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				r.UpdateAt = value.Time
			}
		case role.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case role.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = new(string)
				*r.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return NewRoleClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", r.IsDeleted))
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(r.CreateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(r.UpdateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	if v := r.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role
