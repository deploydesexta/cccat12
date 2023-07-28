package coord

type Coord struct {
	lat  float64
	long float64
}

func New(lat, long float64) Coord {
	return Coord{lat, long}
}

func (c Coord) Lat() float64 {
	return c.lat
}

func (c Coord) Long() float64 {
	return c.long
}
