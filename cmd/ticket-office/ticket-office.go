package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"train-reservation/internal/domain"
	"train-reservation/internal/infrastructure/adapters"
)

type Result struct {
	TrainId          string   `json:"train_id"`
	BookingReference string   `json:"booking_reference"`
	Seats            []string `json:"seats"`
}

func main() {
	trainId := os.Args[1]
	numberOfSeats, _ := strconv.Atoi(os.Args[2])

	trainDataServiceClient := adapters.TrainDataServiceClient{}
	bookingReferenceDataServiceClient := adapters.BookingReferenceDataServiceClient{}
	ticketOffice := domain.NewTicketOffice(trainDataServiceClient, bookingReferenceDataServiceClient, trainDataServiceClient)
	reserve, err := ticketOffice.Reserve(numberOfSeats, trainId)
	if err != nil {
		panic(err)
	}

	var result Result
	result.TrainId = trainId
	result.BookingReference = reserve.BookingReference
	for _, seat := range reserve.Seats {
		result.Seats = append(result.Seats, fmt.Sprint(seat.Id, reserve.Coach.Id))
	}

	resultBytes, _ := json.Marshal(result)
	fmt.Println(string(resultBytes))
}
