package match

import (
	"context"
	"fmt"

	"github.com/fbngrm/around-home/pkg/location"
	latLong "github.com/fbngrm/around-home/pkg/location"
	"github.com/fbngrm/around-home/pkg/materials"
	"github.com/fbngrm/around-home/pkg/partner"
)

type MatchService struct {
	partnerRepo     PartnerRepo
	locationService LocationService
}

func NewMatchService(pr PartnerRepo, ls LocationService) *MatchService {
	return &MatchService{
		partnerRepo:     pr,
		locationService: ls,
	}
}

func (p *MatchService) GetMatches(ctx context.Context, m materials.Material, location string) (Matches, error) {
	requestLocation, err := latLong.New(location)
	if err != nil {
		return nil, fmt.Errorf("could not parse customer location from request: %w", err)
	}

	partnersWithMaterial, err := p.partnerRepo.GetPartnersWithMaterial(ctx, m)
	if err != nil {
		return nil, fmt.Errorf("could not get partners with material [%v]: %w", m, err)
	}

	partnersInRadius := p.getMatchesInOperatingRadius(ctx, partnersWithMaterial, requestLocation)
	partnersInRadius.sortByRatingAndDistance()

	return partnersInRadius, nil
}

func (p *MatchService) getMatchesInOperatingRadius(ctx context.Context, partners []partner.Partner, requestLocation location.Location) Matches {
	var operatingInRadius Matches
	for _, partner := range partners {
		distance := p.locationService.CalculateDistance(partner.Address, requestLocation)
		if !p.locationService.IsInOperationRadius(distance, partner.RadiusOfOperation) {
			continue
		}
		operatingInRadius = append(operatingInRadius, Match{
			Partner:  partner,
			Distance: distance,
		})
	}
	return operatingInRadius
}
