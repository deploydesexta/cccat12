package ride_test

import (
	"github.com/deploydesexta/cccat12/internal/domain/ride"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCalculateRide(t *testing.T) {

	assertions := assert.New(t)

	t.Run("Deve fazer o cálculo do preço de uma corrida durante o dia", func(t *testing.T) {
		r := ride.New()

		err := r.AddPosition(-27.584905257808835, -48.545022195325124, time.Date(2021, 3, 1, 10, 0, 0, 0, time.UTC))
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		err = r.AddPosition(-27.496887588317275, -48.522234807851476, time.Date(2021, 3, 1, 10, 0, 0, 0, time.UTC))
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		price := r.Calculate()

		expectedPrice := 21.0
		assertions.Equal(expectedPrice, price)
	})

	t.Run("Deve fazer o cálculo do preço de uma corrida durante a noite", func(t *testing.T) {
		r := ride.New()

		err := r.AddPosition(-27.584905257808835, -48.545022195325124, time.Date(2021, 3, 1, 23, 0, 0, 0, time.UTC))
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		err = r.AddPosition(-27.496887588317275, -48.522234807851476, time.Date(2021, 3, 1, 23, 0, 0, 0, time.UTC))
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		price := r.Calculate()

		expectedPrice := 39.0
		assertions.Equal(expectedPrice, price)
	})

	t.Run("Deve fazer o cálculo do preço de uma corrida no domingo de dia", func(t *testing.T) {
		r := ride.New()

		err := r.AddPosition(-27.584905257808835, -48.545022195325124, time.Date(2021, 3, 7, 10, 0, 0, 0, time.UTC))
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		err = r.AddPosition(-27.496887588317275, -48.522234807851476, time.Date(2021, 3, 7, 10, 0, 0, 0, time.UTC))
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		price := r.Calculate()

		expectedPrice := 29.0
		assertions.Equal(expectedPrice, price)
	})

	t.Run("Deve fazer o cálculo do preço de uma corrida no domingo de noite", func(t *testing.T) {
		r := ride.New()

		err := r.AddPosition(-27.584905257808835, -48.545022195325124, time.Date(2021, 3, 7, 23, 0, 0, 0, time.UTC))
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		err = r.AddPosition(-27.496887588317275, -48.522234807851476, time.Date(2021, 3, 7, 23, 0, 0, 0, time.UTC))
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		price := r.Calculate()

		expectedPrice := 50.0
		assertions.Equal(expectedPrice, price)
	})
}

func TestCalculateRide_DuringDayWithMinimumPrice(t *testing.T) {

	r := ride.New()

	err := r.AddPosition(-27.584905257808835, -48.545022195325124, time.Date(2021, 3, 4, 10, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	err = r.AddPosition(-27.579020277800876, -48.50838017206791, time.Date(2021, 3, 4, 10, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	price := r.Calculate()

	expectedPrice := 10.0
	assert.Equal(t, expectedPrice, price)
}
