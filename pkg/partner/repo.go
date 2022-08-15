package partner

import (
	"context"
	"fmt"

	ent "github.com/fbngrm/around-home/ent"
	"github.com/fbngrm/around-home/ent/material"
	"github.com/fbngrm/around-home/pkg/materials"
	_ "github.com/lib/pq"
)

type PartnerRepo struct {
	db *ent.Client
}

func (r *PartnerRepo) GetPartnersWithMaterial(ctx context.Context, m materials.Material) ([]Partner, error) {
	partners, err := r.db.Material.Query().Where(material.MaterialEQ(m)).QueryPartners().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get partners with material [%v]: %w", m, err)
	}

	// convert to domain objects - this step seems unneccessary but we want to stick
	// with clean architecture, for the purpose of consistency and the because this
	// is a coding challenge
	results := make([]Partner, len(partners))
	for i, partner := range partners {
		results[i] = toDomain(partner)
	}
	return results, nil
}
