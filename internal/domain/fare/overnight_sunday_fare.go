package fare

import "github.com/deploydesexta/cccat12/internal/domain/ride/segment"

const overnightSundayFare = 5

func OvernightSundayFareCalculator(s segment.Segment) float64 {
	return s.Distance * overnightSundayFare
}
