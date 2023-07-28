package distance

import (
	"github.com/deploydesexta/cccat12/internal/domain/distance/coord"
	"math"
)

// Calculate calculates the distance between two points
// Distance Calculate é domain Entity, pois é utilizado internamento no domínio e não é utilizado por um ator externo.
// No DDD isso seria um domain service. Como não existe esse conceito em clean arch, é um domain entity.
func Calculate(from, to coord.Coord) float64 {
	earthRadius := 6371.0
	degreesToRadians := math.Pi / 180.0
	deltaLat := (to.Lat() - from.Lat()) * degreesToRadians
	deltaLong := (to.Long() - from.Long()) * degreesToRadians
	a :=
		math.Pow(math.Sin(deltaLat/2), 2) +
			math.Cos(from.Lat()*degreesToRadians)*
				math.Cos(to.Lat()*degreesToRadians)*
				math.Pow(math.Sin(deltaLong/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return math.Round(earthRadius * c)
}
