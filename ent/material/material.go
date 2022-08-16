// Code generated by ent, DO NOT EDIT.

package material

import (
	"fmt"

	"github.com/fbngrm/around-home/pkg/materials"
)

const (
	// Label holds the string label denoting the material type in the database.
	Label = "material"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMaterial holds the string denoting the material field in the database.
	FieldMaterial = "material"
	// EdgePartners holds the string denoting the partners edge name in mutations.
	EdgePartners = "partners"
	// Table holds the table name of the material in the database.
	Table = "materials"
	// PartnersTable is the table that holds the partners relation/edge. The primary key declared below.
	PartnersTable = "partner_materials"
	// PartnersInverseTable is the table name for the Partner entity.
	// It exists in this package in order to avoid circular dependency with the "partner" package.
	PartnersInverseTable = "partners"
)

// Columns holds all SQL columns for material fields.
var Columns = []string{
	FieldID,
	FieldMaterial,
}

var (
	// PartnersPrimaryKey and PartnersColumn2 are the table columns denoting the
	// primary key for the partners relation (M2M).
	PartnersPrimaryKey = []string{"partner_id", "material_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// MaterialValidator is a validator for the "material" field enum values. It is called by the builders before save.
func MaterialValidator(m materials.Material) error {
	switch m {
	case "WOOD", "CARPET", "TILES":
		return nil
	default:
		return fmt.Errorf("material: invalid enum value for material field: %q", m)
	}
}