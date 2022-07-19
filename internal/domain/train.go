package domain

type TrainId = string

type Train struct {
	Id     TrainId
	Coachs []Coach
}

func (t Train) GetNumberOfSeats() int {
	total := 0

	for _, coach := range t.Coachs {
		total += coach.GetNumberOfSeats()
	}

	return total
}

func (t Train) GetNumberOfAvailableSeats() int {
	total := 0

	for _, coach := range t.Coachs {
		total += coach.GetNumberOfAvailableSeats()
	}

	return total
}

func (t Train) IsTrainFilledMoreThan70Percent(numberOfSeats int) bool {
	totalNumberOfSeats := t.GetNumberOfSeats()
	numberOfAvailableSeats := t.GetNumberOfAvailableSeats()

	numberOfSeatsOccupied := (totalNumberOfSeats - numberOfAvailableSeats) + numberOfSeats
	occupationPercentage := 100 * numberOfSeatsOccupied / totalNumberOfSeats

	return occupationPercentage >= 70
}

func (t Train) FindSeats(numberOfSeats int) (coach Coach, seats []Seat, found bool) {
	if t.IsTrainFilledMoreThan70Percent(numberOfSeats) {
		return Coach{}, nil, false
	}

	for _, coach := range t.Coachs {
		if coach.IsPerfect(numberOfSeats) {
			return coach, append(seats, coach.GetAvailableSeats()[0:numberOfSeats]...), true
		}
	}

	for _, coach := range t.Coachs {
		if coach.HasSeatsAvailable(numberOfSeats) {
			return coach, append(seats, coach.GetAvailableSeats()[0:numberOfSeats]...), true
		}
	}

	return Coach{}, nil, false
}
