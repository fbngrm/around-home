package match

import (
	"context"

	"github.com/fbngrm/around-home/pkg/location"
	"github.com/fbngrm/around-home/pkg/materials"
	"github.com/fbngrm/around-home/pkg/partner"
)

type PartnerRepo interface {
	GetPartnersWithMaterial(ctx context.Context, m materials.Material) ([]partner.Partner, error)
}

type LocationService interface {
	CalculateDistance(l1, l2 location.Location) float64
	IsInOperationRadius(distance, radius float64) bool
}
