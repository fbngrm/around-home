// Code generated by ent, DO NOT EDIT.

package partner

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/fbngrm/around-home/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// Rating applies equality check predicate on the "rating" field. It's identical to RatingEQ.
func Rating(v int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRating), v))
	})
}

// RadiusOfOperation applies equality check predicate on the "radiusOfOperation" field. It's identical to RadiusOfOperationEQ.
func RadiusOfOperation(v float64) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRadiusOfOperation), v))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Partner {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Partner {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// RatingEQ applies the EQ predicate on the "rating" field.
func RatingEQ(v int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRating), v))
	})
}

// RatingNEQ applies the NEQ predicate on the "rating" field.
func RatingNEQ(v int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRating), v))
	})
}

// RatingIn applies the In predicate on the "rating" field.
func RatingIn(vs ...int) predicate.Partner {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRating), v...))
	})
}

// RatingNotIn applies the NotIn predicate on the "rating" field.
func RatingNotIn(vs ...int) predicate.Partner {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRating), v...))
	})
}

// RatingGT applies the GT predicate on the "rating" field.
func RatingGT(v int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRating), v))
	})
}

// RatingGTE applies the GTE predicate on the "rating" field.
func RatingGTE(v int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRating), v))
	})
}

// RatingLT applies the LT predicate on the "rating" field.
func RatingLT(v int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRating), v))
	})
}

// RatingLTE applies the LTE predicate on the "rating" field.
func RatingLTE(v int) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRating), v))
	})
}

// RadiusOfOperationEQ applies the EQ predicate on the "radiusOfOperation" field.
func RadiusOfOperationEQ(v float64) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRadiusOfOperation), v))
	})
}

// RadiusOfOperationNEQ applies the NEQ predicate on the "radiusOfOperation" field.
func RadiusOfOperationNEQ(v float64) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRadiusOfOperation), v))
	})
}

// RadiusOfOperationIn applies the In predicate on the "radiusOfOperation" field.
func RadiusOfOperationIn(vs ...float64) predicate.Partner {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRadiusOfOperation), v...))
	})
}

// RadiusOfOperationNotIn applies the NotIn predicate on the "radiusOfOperation" field.
func RadiusOfOperationNotIn(vs ...float64) predicate.Partner {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRadiusOfOperation), v...))
	})
}

// RadiusOfOperationGT applies the GT predicate on the "radiusOfOperation" field.
func RadiusOfOperationGT(v float64) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRadiusOfOperation), v))
	})
}

// RadiusOfOperationGTE applies the GTE predicate on the "radiusOfOperation" field.
func RadiusOfOperationGTE(v float64) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRadiusOfOperation), v))
	})
}

// RadiusOfOperationLT applies the LT predicate on the "radiusOfOperation" field.
func RadiusOfOperationLT(v float64) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRadiusOfOperation), v))
	})
}

// RadiusOfOperationLTE applies the LTE predicate on the "radiusOfOperation" field.
func RadiusOfOperationLTE(v float64) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRadiusOfOperation), v))
	})
}

// HasMaterials applies the HasEdge predicate on the "materials" edge.
func HasMaterials() predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MaterialsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MaterialsTable, MaterialsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMaterialsWith applies the HasEdge predicate on the "materials" edge with a given conditions (other predicates).
func HasMaterialsWith(preds ...predicate.Material) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MaterialsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MaterialsTable, MaterialsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Partner) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Partner) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Partner) predicate.Partner {
	return predicate.Partner(func(s *sql.Selector) {
		p(s.Not())
	})
}
