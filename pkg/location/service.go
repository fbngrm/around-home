package location

type LocationService struct {
	calculator Calculator
}

func NewService(c Calculator) *LocationService {
	return &LocationService{calculator: c}
}

func (s *LocationService) CalculateDistance(l1, l2 Location) float64 {
	return s.calculator.DistanceInMeters(l1.Lat, l1.Long, l2.Lat, l2.Long)
}

func (s *LocationService) IsInOperationRadius(distance, radius float64) bool {
	return radius >= distance
}
