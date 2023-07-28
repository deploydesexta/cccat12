package fare

import (
	"errors"
	"github.com/deploydesexta/cccat12/internal/domain/ride/segment"
)

var InvalidSegment = errors.New("invalid segment")

type Calculator func(s segment.Segment) float64

func NewFareCalculator(s segment.Segment) (Calculator, error) {
	if s.IsOvernight() && !s.IsSunday() {
		return OvernightFareCalculator, nil
	}

	if s.IsOvernight() && s.IsSunday() {
		return OvernightSundayFareCalculator, nil
	}

	if !s.IsOvernight() && s.IsSunday() {
		return SundayFareCalculator, nil
	}

	if !s.IsOvernight() && !s.IsSunday() {
		return NormalFareCalculator, nil
	}

	return nil, InvalidSegment
}
