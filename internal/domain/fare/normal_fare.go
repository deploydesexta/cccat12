package fare

import "github.com/deploydesexta/cccat12/internal/domain/ride/segment"

const normalFare = 2.1

func NormalFareCalculator(s segment.Segment) float64 {
	return s.Distance * normalFare
}
