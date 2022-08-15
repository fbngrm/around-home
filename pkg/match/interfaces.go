package match

import (
	"context"

	"github.com/fbngrm/around-home/pkg/materials"
	"github.com/fbngrm/around-home/pkg/partner"
)

type PartnerRepo interface {
	GetPartnersWithMaterial(ctx context.Context, m materials.Material) ([]partner.Partner, error)
}
