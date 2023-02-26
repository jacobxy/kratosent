package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type DBMixin struct {
	mixin.Schema
}

func (DBMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().StorageKey("id").GoType(int64(0)),
		field.Int8("is_deleted").Default(0).GoType(int8(0)),
		field.Time("create_at").Immutable().Default(time.Now),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
	}
}
