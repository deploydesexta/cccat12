package ride

import (
	"github.com/deploydesexta/cccat12/internal/domain/distance"
	"github.com/deploydesexta/cccat12/internal/domain/fare"
	"github.com/deploydesexta/cccat12/internal/domain/ride/position"
	"github.com/deploydesexta/cccat12/internal/domain/ride/segment"
	"log"
	"time"
)

const (
	MinPrice = 10
)

type Ride struct {
	Positions []*position.Position
}

func New() *Ride {
	return &Ride{
		Positions: make([]*position.Position, 0),
	}
}

func (r *Ride) AddPosition(lat, long float64, date time.Time) error {
	p := position.New(lat, long, date)
	r.Positions = append(r.Positions, &p)
	return nil
}

func (r *Ride) Calculate() float64 {
	var price float64
	l := len(r.Positions)

	for i, p := range r.Positions {
		if i+1 >= l {
			break
		}

		nextP := r.Positions[i+1]

		s, err := segment.New(distance.Calculate(p.Coord(), nextP.Coord()), nextP.Date())
		if err != nil {
			log.Printf("error creating segment: %v\n", err)
			break
		}

		fareCalculator, err := fare.NewFareCalculator(*s)
		if err != nil {
			log.Print("invalid fare calculator\n")
			break
		}

		price += fareCalculator(*s)
	}

	if price < MinPrice {
		return MinPrice
	} else {
		return price
	}
}
