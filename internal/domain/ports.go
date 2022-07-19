package domain

type BookingReferenceProvider interface {
	Get() (BookingReference, error)
}

type TrainProvider interface {
	Get(trainId TrainId) (Train, error)
}

type ReservationClient interface {
	Reserve(train Train, confirmation ReservationConfirmation) (bool, error)
}
