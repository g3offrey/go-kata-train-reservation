package domain

type SeatId = string

type Seat struct {
	Id               SeatId
	BookingReference BookingReference
}

func (s Seat) IsAvailable() bool {
	return s.BookingReference == ""
}
