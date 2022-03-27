package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// MembersClass holds the schema definition for the MembersClass entity.
type MembersClass struct {
	ent.Schema
}

func (MembersClass) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the MembersClass.
func (MembersClass) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date_of_admission"),
		field.Time("date_of_withdrawal").Optional(),
	}
}

// Edges of the MembersClass.
func (MembersClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("member", Member.Type).Ref("members_classes").Unique().Required(),
		edge.From("class", Class.Type).Ref("members_classes").Unique().Required(),
	}
}
