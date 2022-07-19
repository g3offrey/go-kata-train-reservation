package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"train-reservation/internal/domain"
)

func TestShouldNotWelcomePassengersIfTrainIsMoreThan70PercentFully(t *testing.T) {
	train := domain.Train{
		Id: "PoudlardExpress",
		Coachs: []domain.Coach{
			{
				Id: "A",
				Seats: []domain.Seat{
					{
						Id:               "1",
						BookingReference: "USED",
					},
				},
			},
			{
				Id: "B",
				Seats: []domain.Seat{
					{
						Id:               "1",
						BookingReference: "USED",
					},
					{
						Id: "2",
					},
					{
						Id: "3",
					},
				},
			},
		},
	}

	coach, seats, found := train.FindSeats(2)

	assert.False(t, found)
	assert.Equal(t, domain.Coach{}, coach)
	assert.Equal(t, []domain.Seat(nil), seats)
}

func TestShouldKeepPassengersInSameCoach(t *testing.T) {
	train := domain.Train{
		Id: "PoudlardExpress",
		Coachs: []domain.Coach{
			{
				Id: "A",
				Seats: []domain.Seat{
					{
						Id: "1",
					},
				},
			},
			{
				Id: "B",
				Seats: []domain.Seat{
					{
						Id:               "1",
						BookingReference: "USED",
					},
					{
						Id: "2",
					},
					{
						Id: "3",
					},
					{
						Id: "4",
					},
				},
			},
		},
	}

	coach, seats, found := train.FindSeats(2)

	coachToUse := train.Coachs[1]
	assert.True(t, found)
	assert.Equal(t, coachToUse, coach)
	assert.Equal(t, []domain.Seat{coachToUse.Seats[1], coachToUse.Seats[2]}, seats)
}

func TestShouldTryToKeepCoachNotMoreThan70PercentFully(t *testing.T) {
	train := domain.Train{
		Id: "PoudlardExpress",
		Coachs: []domain.Coach{
			{
				Id: "A",
				Seats: []domain.Seat{
					{
						Id: "1",
					},
					{
						Id: "2",
					},
					{
						Id:               "3",
						BookingReference: "USED",
					},
					{
						Id:               "4",
						BookingReference: "USED",
					},
				},
			},
			{
				Id: "B",
				Seats: []domain.Seat{
					{
						Id: "1",
					},
					{
						Id: "2",
					},
					{
						Id: "3",
					},
				},
			},
		},
	}

	coach, seats, found := train.FindSeats(2)

	coachToUse := train.Coachs[1]
	assert.True(t, found)
	assert.Equal(t, coachToUse, coach)
	assert.Equal(t, []domain.Seat{coachToUse.Seats[0], coachToUse.Seats[1]}, seats)
}
