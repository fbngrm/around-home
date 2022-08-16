package location

import (
	"fmt"
	"strconv"
	"strings"
)

type Location struct {
	Lat  float64
	Long float64
}

// Todo: add expected format to errors
// Todo: add a proper format checker/handler
func New(location string) (Location, error) {
	parts := strings.Split(location, ":")
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

func (l Location) String() string {
	return fmt.Sprintf("%f:%f", l.Lat, l.Long)
}
