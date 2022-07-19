package domain

type CoachId = string

type Coach struct {
	Id    CoachId
	Seats []Seat
}

func (c Coach) GetNumberOfSeats() int {
	return len(c.Seats)
}

func (c Coach) GetNumberOfAvailableSeats() int {
	return len(c.GetAvailableSeats())
}

func (c Coach) GetAvailableSeats() []Seat {
	availableSeats := []Seat{}

	for _, seat := range c.Seats {
		if seat.IsAvailable() {
			availableSeats = append(availableSeats, seat)
		}
	}

	return availableSeats
}

func (c Coach) HasSeatsAvailable(numberOfSeats int) bool {
	return len(c.GetAvailableSeats()) >= numberOfSeats
}

func (c Coach) IsPerfect(numberOfSeats int) bool {
	numberOfSeatsInCoach := c.GetNumberOfSeats()
	numberOfAvailableSeatsInCoach := c.GetNumberOfAvailableSeats()

	numberOfSeatsOccupied := (numberOfSeatsInCoach - numberOfAvailableSeatsInCoach) + numberOfSeats
	occupationPercentage := 100 * numberOfSeatsOccupied / numberOfSeatsInCoach

	return occupationPercentage <= 70
}
