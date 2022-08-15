package match

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	apiv1 "github.com/fbngrm/around-home/gen/proto/go/match/v1"
	"github.com/fbngrm/around-home/pkg/materials"
	"github.com/fbngrm/around-home/pkg/partner"
)

type partnerMatchService struct {
	partnerRepo PartnerRepo
}

func NewPartnerService(partnerRepo PartnerRepo) *partnerMatchService {
	return &partnerMatchService{
		partnerRepo: partnerRepo,
	}
}

func (p *partnerMatchService) MatchPartnersWithRequest(ctx context.Context, req *apiv1.ListPartnersForTaskRequest) (*apiv1.ListPartnersForTaskResponse, error) {
	// we need to take special care here since all request input is potentially harmful
	// thus, we cannot just assert a type to the req.Material
	material, err := materials.NewMaterial(req.Material)
	if err != nil {
		return nil, fmt.Errorf("could not parse material: %w", err)
	}

	l := ToLocation(req.Location)

	partners, err := p.partnerRepo.GetPartnersWithMaterial(ctx, material)
	if err != nil {
		return nil, fmt.Errorf("could not get partners with material [%v]: %w", material, err)
	}

	return &apiv1.ListPartnersForTaskResponse{}, nil
}

type Match struct {
	partner  partner.Partner
	distance int
}

type Location struct {
	Lat  float64
	Lang float64
}

func ToLocation(location string) (Location, error) {
	parts := strings.Split(location, "/")
	if len(parts) != 2 {
		return Location{}, fmt.Errorf("could not parse location [%q]", location)
	}
	value, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return Location{}, fmt.Errorf("could not parse lat [%q]", location)
	}
	lat := float64(value)
	value, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return Location{}, fmt.Errorf("could not parse long [%q]", location)
	}
	long := float64(value)
	return Location{
		Lat:  lat,
		Long: long,
	}, nil
}

func (p *partnerMatchService) GetPartnersInOperatingRadius(ctx context.Context, partners []partner.Partner, location Location) ([]partner.Partner, error) {

	return nil, nil
}
