package materials

import (
	"fmt"
	"strings"
)

type Material string

const (
	Wood            = "WOOD"
	Carpet Material = "CARPET"
	Tiles  Material = "TILES"
)

func NewMaterial(s string) (Material, error) {
	switch strings.ToUpper(s) {
	case string(Wood):
		return Wood, nil
	case string(Carpet):
		return Carpet, nil
	case string(Tiles):
		return Tiles, nil
	default:
	}
	// note, we want to always escape the input here since it
	// is potentially harmful user input
	return "", fmt.Errorf("material not supported [%q]", s)
}

// Values provides a list of valid values for Material enum.
func (Material) Values() []string {
	var kinds []string
	for _, s := range []Material{Wood, Carpet, Tiles} {
		kinds = append(kinds, string(s))
	}
	return kinds
}
