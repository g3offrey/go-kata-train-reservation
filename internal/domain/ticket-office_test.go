package domain_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"train-reservation/internal/domain"
	"train-reservation/internal/domain/mocks"
)

func TestReserveWhenPlaceAvailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTrainProvider := mock_domain.NewMockTrainProvider(ctrl)
	mockBookingReferenceProvider := mock_domain.NewMockBookingReferenceProvider(ctrl)
	mockReservationClient := mock_domain.NewMockReservationClient(ctrl)
	train := domain.Train{
		Id: "PoudlardExpress",
		Coachs: []domain.Coach{
			{
				Id: "A",
				Seats: []domain.Seat{
					{Id: "1"},
					{Id: "2"},
				},
			},
		},
	}

	mockTrainProvider.EXPECT().Get(gomock.Eq("PoudlardExpress")).Return(train, nil)
	mockBookingReferenceProvider.EXPECT().Get().Return("123abc", nil)
	mockReservationClient.EXPECT().Reserve(
		gomock.Eq(train),
		gomock.Eq(
			domain.ReservationConfirmation{
				Coach: train.Coachs[0],
				Seats: []domain.Seat{
					{Id: "1"},
				},
				BookingReference: "123abc",
			},
		),
	).Return(true, nil)

	ticketOffice := domain.NewTicketOffice(mockTrainProvider, mockBookingReferenceProvider, mockReservationClient)
	reserve, err := ticketOffice.Reserve(1, "PoudlardExpress")

	assert.NoError(t, err)
	assert.Equal(t, reserve, domain.ReservationConfirmation{
		Coach: train.Coachs[0],
		Seats: []domain.Seat{
			{Id: "1"},
		},
		BookingReference: "123abc",
	})
}
