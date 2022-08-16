package partner

import (
	"errors"
	"fmt"

	"github.com/fbngrm/around-home/ent"
	"github.com/fbngrm/around-home/pkg/location"
	materialsDomain "github.com/fbngrm/around-home/pkg/materials"
)

func toDomain(partner *ent.Partner) (Partner, error) {
	if partner == nil {
		return Partner{}, errors.New("partner is nil")
	}
	var partnerMaterials []materialsDomain.Material
	for _, material := range partner.Edges.Materials {
		partnerMaterials = append(partnerMaterials, material.Material)
	}
	loc, err := location.New(partner.Address)
	if err != nil {
		return Partner{}, fmt.Errorf("could not create location [%v]: %w", partner.Address, err)
	}
	return Partner{
		ID:                partner.ID,
		Materials:         partnerMaterials,
		Address:           loc,
		Rating:            partner.Rating,
		RadiusOfOperation: partner.RadiusOfOperation,
	}, nil
}
