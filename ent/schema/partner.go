package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Partner struct {
	ent.Schema
}

func (Partner) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").NotEmpty(),
		field.Int("rating"),
		field.Float("radiusOfOperation"),
	}
}

func (Partner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("materials", Material.Type),
	}
}
