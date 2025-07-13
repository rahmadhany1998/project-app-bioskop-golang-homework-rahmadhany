package dto

type BookingRequest struct {
	CinemaID      int    `json:"cinemaId"`
	SeatID        string `json:"seatId"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	PaymentMethod string `json:"paymentMethod"`
}
