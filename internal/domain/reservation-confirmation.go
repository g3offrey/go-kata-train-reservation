package domain

type ReservationConfirmation struct {
	Coach            Coach
	Seats            []Seat
	BookingReference BookingReference
}
