package api

import (
	"context"
	"fmt"

	apiv1 "github.com/fbngrm/around-home/gen/proto/go/match/v1"
	"github.com/fbngrm/around-home/pkg/materials"
)

type Api struct {
	matchService Matcher
}

func NewApi(matchService Matcher) *Api {
	return &Api{
		matchService: matchService,
	}
}

func (a *Api) MatchPartnersWithRequest(ctx context.Context, req *apiv1.MatchPartnersWithRequestInput) (*apiv1.MatchPartnersWithRequestOutput, error) {
	// we need to take special care here since all request input is potentially harmful
	// thus, we cannot just assert a type to the req.Material
	material, err := materials.NewMaterial(req.Material)
	if err != nil {
		return nil, fmt.Errorf("could not parse material: %w", err)
	}

	matches, err := a.matchService.GetMatches(ctx, material, req.Location)
	if err != nil {
		return nil, fmt.Errorf("could not get matches: %w", err)
	}

	results := make([]*apiv1.Partner, len(matches))
	for i, m := range matches {
		results[i] = &apiv1.Partner{
			Id:       uint32(m.Partner.ID),
			Rating:   uint32(m.Partner.Rating),
			Location: m.Partner.Address.String(),
			Radius:   m.Partner.RadiusOfOperation,
			Distance: m.Distance,
		}
	}
	return &apiv1.MatchPartnersWithRequestOutput{Partner: results}, nil
}
