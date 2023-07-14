package ride

import (
	"github.com/deploydesexta/cccat12/src/segment"
	"time"
)

const (
	OvernightFare       = 3.90
	OvernightSundayFare = 5
	SundayFare          = 2.9
	NormalFare          = 2.1
	MinPrice            = 10
)

type Ride struct {
	Segments []*segment.Segment
}

func New() *Ride {
	return &Ride{
		Segments: make([]*segment.Segment, 0),
	}
}

func (r *Ride) AddSegment(distance float64, date time.Time) error {
	s, err := segment.New(distance, date)
	if err != nil {
		return err
	}

	r.Segments = append(r.Segments, s)
	return nil
}

func (r *Ride) Calculate() float64 {
	var price float64

	for _, s := range r.Segments {
		if s.IsOvernight() && !s.IsSunday() {
			price += s.Distance * OvernightFare
		}

		if s.IsOvernight() && s.IsSunday() {
			price += s.Distance * OvernightSundayFare
		}

		if !s.IsOvernight() && s.IsSunday() {
			price += s.Distance * SundayFare
		}

		if !s.IsOvernight() && !s.IsSunday() {
			price += s.Distance * NormalFare
		}

	}

	if price < MinPrice {
		return MinPrice
	} else {
		return price
	}
}
