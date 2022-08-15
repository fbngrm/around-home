package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fbngrm/around-home/pkg/materials"
)

type Material struct {
	ent.Schema
}

func (Material) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("material").
			GoType(materials.Material("")),
	}
}

func (Material) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("partners", Partner.Type).
			Ref("materials"),
	}
}
