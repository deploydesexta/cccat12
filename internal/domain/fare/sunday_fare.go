package fare

import "github.com/deploydesexta/cccat12/internal/domain/ride/segment"

const sundayFare = 2.9

func SundayFareCalculator(s segment.Segment) float64 {
	return s.Distance * sundayFare
}
