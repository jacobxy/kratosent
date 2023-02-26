package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

func (Department) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DBMixin{},
	}
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int64("parent_department_id").Optional().Nillable(),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return nil
}
