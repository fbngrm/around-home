package location

type Calculator interface {
	DistanceInMeters(lat1, lon1, lat2, lon2 float64) float64
}
