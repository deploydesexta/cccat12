package fare

import "github.com/deploydesexta/cccat12/internal/domain/ride/segment"

const overnightFare = 3.9

func OvernightFareCalculator(s segment.Segment) float64 {
	return s.Distance * overnightFare
}
