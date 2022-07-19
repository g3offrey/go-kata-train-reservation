package domain

type TicketOffice struct {
	trainProvider            TrainProvider
	reservationClient        ReservationClient
	bookingReferenceProvider BookingReferenceProvider
}

func (ticketOffice TicketOffice) Reserve(numberOfSeats int, trainId TrainId) (ReservationConfirmation, error) {
	train, err := ticketOffice.trainProvider.Get(trainId)
	if err != nil {
		return ReservationConfirmation{}, err
	}

	coach, seats, found := train.FindSeats(numberOfSeats)
	if !found {
		return ReservationConfirmation{}, nil
	}

	bookingReference, err := ticketOffice.bookingReferenceProvider.Get()
	if err != nil {
		return ReservationConfirmation{}, err
	}

	reservationConfirmation := ReservationConfirmation{
		Coach:            coach,
		Seats:            seats,
		BookingReference: bookingReference,
	}

	reserved, err := ticketOffice.reservationClient.Reserve(train, reservationConfirmation)
	if err != nil {
		return ReservationConfirmation{}, err
	}
	if !reserved {
		return ReservationConfirmation{}, nil
	}

	return reservationConfirmation, nil
}

func NewTicketOffice(trainProvider TrainProvider, bookingReferenceProvider BookingReferenceProvider, reservationClient ReservationClient) *TicketOffice {
	ticketOffice := TicketOffice{
		trainProvider:            trainProvider,
		bookingReferenceProvider: bookingReferenceProvider,
		reservationClient:        reservationClient,
	}

	return &ticketOffice
}
