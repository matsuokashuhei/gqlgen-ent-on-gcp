package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Studio holds the schema definition for the Studio entity.
type Studio struct {
	ent.Schema
}

func (Studio) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Studio.
func (Studio) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("location"),
	}
}

// Edges of the Studio.
func (Studio) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("school", School.Type).Ref("studios").Unique(),
		edge.To("rooms", Room.Type),
	}
}
