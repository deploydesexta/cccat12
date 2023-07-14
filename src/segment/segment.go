package segment

import (
	"errors"
	"time"
)

type Segment struct {
	Distance float64
	Date     time.Time
}

func New(distance float64, date time.Time) (*Segment, error) {
	if !isValidDistance(distance) {
		return nil, errors.New("invalid distance")
	}

	if !isValidDate(date) {
		return nil, errors.New("invalid date")
	}

	return &Segment{distance, date}, nil
}

func (s *Segment) IsOvernight() bool {
	return s.Date.Hour() >= 22 || s.Date.Hour() <= 6
}

func (s *Segment) IsSunday() bool {
	return s.Date.Weekday() == time.Sunday
}

func isValidDistance(distance float64) bool {
	return distance > 0
}

func isValidDate(date time.Time) bool {
	return !date.IsZero()
}
