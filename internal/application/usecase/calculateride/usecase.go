package calculateride

import (
	"github.com/deploydesexta/cccat12/internal/domain/ride"
	"time"
)

type (
	UseCase struct {
	}

	Input struct {
		Positions *[]struct {
			Lat  float64
			Long float64
			Date string
		}
	}

	Output struct {
		Price float64
	}
)

func New() *UseCase {
	return &UseCase{}
}

func (c UseCase) Execute(input Input) (Output, error) {
	r := ride.New()
	for _, s := range *input.Positions {
		t, err := time.Parse("2006-01-02T15:04:05", s.Date)
		if err != nil {
			return Output{}, err
		}

		err = r.AddPosition(s.Lat, s.Long, t)
		if err != nil {
			return Output{}, err
		}
	}
	return Output{r.Calculate()}, nil
}
