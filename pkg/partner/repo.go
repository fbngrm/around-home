package partner

import (
	"context"
	"fmt"
	"log"

	ent "github.com/fbngrm/around-home/ent"
	"github.com/fbngrm/around-home/ent/material"
	"github.com/fbngrm/around-home/pkg/materials"
	_ "github.com/lib/pq"
)

type Repo struct {
	DB *ent.Client
}

func (r *Repo) GetPartnersWithMaterial(ctx context.Context, m materials.Material) ([]Partner, error) {
	partners, err := r.DB.Material.Query().Where(material.MaterialEQ(m)).QueryPartners().WithMaterials().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get partners with material [%v]: %w", m, err)
	}

	// convert to domain objects - this step seems unneccessary but we want to stick
	// with clean architecture, for the purpose of consistency and the because this
	// is a coding challenge
	results := make([]Partner, len(partners))
	for i, partner := range partners {
		p, err := toDomain(partner)
		// if conversion fails, we ignore this partner and return the valid ones
		// in a real-world scenario, we would inject a logger and have a more advanced error handling
		if err != nil {
			log.Printf("warning: could not convert to domain Partner object, skipping [%d]: %s\n", partner.ID, err.Error())
			continue
		}
		results[i] = p
	}
	return results, nil
}
