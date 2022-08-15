package materials

type Material string

const (
	Wood            = "WOOD"
	Carpet Material = "CARPET"
	Tiles  Material = "TILES"
)

// Values provides list valid values for Enum.
func (Material) Values() []string {
	var kinds []string
	for _, s := range []Material{Wood, Carpet, Tiles} {
		kinds = append(kinds, string(s))
	}
	return kinds
}
