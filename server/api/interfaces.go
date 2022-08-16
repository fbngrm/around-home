package api

import (
	"context"

	"github.com/fbngrm/around-home/pkg/match"
	"github.com/fbngrm/around-home/pkg/materials"
)

type Matcher interface {
	GetMatches(ctx context.Context, m materials.Material, location string) (match.Matches, error)
}
