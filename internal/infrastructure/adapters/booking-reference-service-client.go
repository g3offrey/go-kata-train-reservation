package adapters

import (
	"io"
	"net/http"
	"train-reservation/internal/domain"
)

type BookingReferenceDataServiceClient struct {
}

func (client BookingReferenceDataServiceClient) Get() (domain.BookingReference, error) {
	get, err := http.Get("http://localhost:8082/booking_reference")
	if err != nil {
		return "", err
	}
	defer get.Body.Close()

	body, err := io.ReadAll(get.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
