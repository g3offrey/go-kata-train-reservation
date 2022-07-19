package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"train-reservation/internal/domain"
)

type TrainDataServiceClient struct {
}

type SeatData struct {
	BookingReference string `json:"booking_reference"`
	SeatNumber       string `json:"seat_number"`
	Coach            string `json:"coach"`
}

type TrainData struct {
	Seats map[string]SeatData `json:"seats"`
}

func (client TrainDataServiceClient) Get(trainId domain.TrainId) (domain.Train, error) {
	get, err := http.Get("http://localhost:8081/data_for_train/" + trainId)
	if err != nil {
		return domain.Train{}, err
	}
	defer get.Body.Close()

	var trainData TrainData
	err = json.NewDecoder(get.Body).Decode(&trainData)
	if err != nil {
		return domain.Train{}, err
	}

	if trainData.Seats == nil {
		return domain.Train{}, errors.New("train data not found")
	}

	var coachs = map[string]domain.Coach{}
	for _, seat := range trainData.Seats {
		coachId := seat.Coach
		coach, found := coachs[coachId]
		if !found {
			coach = domain.Coach{Id: coachId}
			coachs[coachId] = coach
		}

		coach.Seats = append(coach.Seats, domain.Seat{
			Id:               seat.SeatNumber,
			BookingReference: seat.BookingReference,
		})

		coachs[coachId] = coach
	}

	result := domain.Train{
		Id: trainId,
	}

	for _, coach := range coachs {
		sort.Slice(coach.Seats, func(i, j int) bool {
			return coach.Seats[i].Id < coach.Seats[j].Id
		})
		result.Coachs = append(result.Coachs, coach)
	}
	sort.Slice(result.Coachs, func(i, j int) bool {
		return result.Coachs[i].Id < result.Coachs[j].Id
	})

	return result, nil
}

func (client TrainDataServiceClient) Reserve(train domain.Train, confirmation domain.ReservationConfirmation) (bool, error) {
	values := url.Values{}
	values.Add("train_id", train.Id)
	values.Add("booking_reference", confirmation.BookingReference)

	var seatIds []domain.SeatId
	for _, seat := range confirmation.Seats {
		seatIds = append(seatIds, fmt.Sprint(seat.Id, confirmation.Coach.Id))
	}
	seatIdsBytes, _ := json.Marshal(seatIds)
	values.Add("seats", string(seatIdsBytes))

	form, err := http.PostForm("http://localhost:8081/reserve", values)
	if err != nil {
		return false, err
	}
	defer form.Body.Close()

	if form.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}
