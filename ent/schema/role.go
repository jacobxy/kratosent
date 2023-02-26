package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DBMixin{},
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional().Nillable(),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
