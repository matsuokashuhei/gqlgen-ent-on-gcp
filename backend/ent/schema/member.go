package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number").Unique().Annotations(entgql.OrderField("NUMBER")),
		field.String("name"),
		field.String("kana").Annotations(entgql.OrderField("KANA")),
		field.Enum("gender").Values("MALE", "FEMALE", "OTHER"),
		field.Time("date_of_birth").Optional(),
		field.String("phone_number").Optional(),
		field.String("email").Optional(),
		field.Time("date_of_admission"),
		field.Time("date_of_withdrawal").Optional(),
		field.String("memo").Optional(),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members_classes", MembersClass.Type),
	}
}
