package position

import (
	"github.com/deploydesexta/cccat12/internal/domain/distance/coord"
	"time"
)

type Position struct {
	coord coord.Coord
	date  time.Time
}

func New(lat, long float64, date time.Time) Position {
	return Position{coord: coord.New(lat, long), date: date}
}

func (p Position) Coord() coord.Coord {
	return p.coord
}

func (p Position) Date() time.Time {
	return p.date
}
