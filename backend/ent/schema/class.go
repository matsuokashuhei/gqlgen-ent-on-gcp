package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

func (Class) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("tuition"),
		field.Time("start_date"),
		field.Time("end_date").Optional(),
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("schedule", Schedule.Type).Ref("classes").Unique().Required(),
		edge.From("instructor", Instructor.Type).Ref("classes").Unique().Required(),
	}
}
