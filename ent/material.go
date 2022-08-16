// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/fbngrm/around-home/ent/material"
	"github.com/fbngrm/around-home/pkg/materials"
)

// Material is the model entity for the Material schema.
type Material struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Material holds the value of the "material" field.
	Material materials.Material `json:"material,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MaterialQuery when eager-loading is set.
	Edges MaterialEdges `json:"edges"`
}

// MaterialEdges holds the relations/edges for other nodes in the graph.
type MaterialEdges struct {
	// Partners holds the value of the partners edge.
	Partners []*Partner `json:"partners,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PartnersOrErr returns the Partners value or an error if the edge
// was not loaded in eager-loading.
func (e MaterialEdges) PartnersOrErr() ([]*Partner, error) {
	if e.loadedTypes[0] {
		return e.Partners, nil
	}
	return nil, &NotLoadedError{edge: "partners"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Material) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case material.FieldID:
			values[i] = new(sql.NullInt64)
		case material.FieldMaterial:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Material", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Material fields.
func (m *Material) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case material.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case material.FieldMaterial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field material", values[i])
			} else if value.Valid {
				m.Material = materials.Material(value.String)
			}
		}
	}
	return nil
}

// QueryPartners queries the "partners" edge of the Material entity.
func (m *Material) QueryPartners() *PartnerQuery {
	return (&MaterialClient{config: m.config}).QueryPartners(m)
}

// Update returns a builder for updating this Material.
// Note that you need to call Material.Unwrap() before calling this method if this Material
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Material) Update() *MaterialUpdateOne {
	return (&MaterialClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Material entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Material) Unwrap() *Material {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Material is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Material) String() string {
	var builder strings.Builder
	builder.WriteString("Material(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("material=")
	builder.WriteString(fmt.Sprintf("%v", m.Material))
	builder.WriteByte(')')
	return builder.String()
}

// Materials is a parsable slice of Material.
type Materials []*Material

func (m Materials) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
