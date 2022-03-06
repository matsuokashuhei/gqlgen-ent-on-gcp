package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

func (Room) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("capacity"),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("studio", Studio.Type).Ref("rooms").Unique().Required(),
		edge.To("schedules", Schedule.Type),
	}
}
