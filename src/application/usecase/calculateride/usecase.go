package calculateride

import (
	"github.com/deploydesexta/cccat12/src/domain/ride"
	"time"
)

type (
	UseCase struct {
	}

	Input struct {
		Segments *[]struct {
			Distance float64
			Date     string
		}
	}

	Output struct {
		Price float64
	}
)

func NewUseCase() UseCase {
	return UseCase{}
}

func (c UseCase) Execute(input Input) (Output, error) {
	r := ride.New()
	for _, s := range *input.Segments {
		t, err := time.Parse("2006-01-02T15:04:05", s.Date)
		if err != nil {
			return Output{}, err
		}

		err = r.AddSegment(s.Distance, t)
		if err != nil {
			return Output{}, err
		}
	}
	return Output{r.Calculate()}, nil
}
