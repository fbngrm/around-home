package match

import (
	"sort"

	"github.com/fbngrm/around-home/pkg/partner"
)

type Match struct {
	Partner  partner.Partner
	Distance float64
}

type Matches []Match

func (m Matches) sortByRatingAndDistance() {
	sort.Slice(m, func(i, j int) bool {
		var sortedByRating, sortedByDistance bool

		// sort by partner rating
		sortedByRating = m[i].Partner.Rating > m[j].Partner.Rating

		// sort by distance
		if m[i].Partner.Rating == m[j].Partner.Rating {
			sortedByDistance = m[i].Distance < m[j].Distance
			return sortedByDistance
		}

		return sortedByRating
	})
}
