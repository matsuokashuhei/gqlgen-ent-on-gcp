package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Instructor holds the schema definition for the Instructor entity.
type Instructor struct {
	ent.Schema
}

func (Instructor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Instructor.
func (Instructor) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("syllabic_characters"),
		field.String("biography").Optional(),
		field.String("phone_number").Optional(),
		field.String("email").Optional(),
	}
}

// Edges of the Instructor.
func (Instructor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("classes", Class.Type),
	}
}
