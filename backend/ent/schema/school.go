package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// School holds the schema definition for the School entity.
type School struct {
	ent.Schema
}

func (School) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the School.
func (School) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the School.
func (School) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("studios", Studio.Type),
	}
}
