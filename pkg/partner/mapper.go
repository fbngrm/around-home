package partner

import (
	"github.com/fbngrm/around-home/ent"
	"github.com/fbngrm/around-home/pkg/materials"
)

func toDomain(partner *ent.Partner) Partner {
	if partner == nil {
		return Partner{}
	}
	var partnerMaterials []materials.Material
	for _, material := range partner.Edges.Materials {
		partnerMaterials = append(partnerMaterials, material.Material)
	}
	return Partner{
		materials:         partnerMaterials,
		name:              partner.Name,
		address:           partner.Address,
		rating:            partner.Rating,
		radiusOfOperation: partner.RadiusOfOperation,
	}
}
