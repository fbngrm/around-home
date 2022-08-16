package partner

import (
	"github.com/fbngrm/around-home/pkg/location"
	materialsDomain "github.com/fbngrm/around-home/pkg/materials"
)

type Partner struct {
	ID                int
	Materials         []materialsDomain.Material
	Address           location.Location
	RadiusOfOperation float64 // in meters
	Rating            int     // todo: add enum type
}
