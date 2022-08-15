package partner

import materialsDomain "github.com/fbngrm/around-home/pkg/materials"

type Partner struct {
	materials         []materialsDomain.Material
	name              string
	address           string
	rating            int
	radiusOfOperation int
}
