package entity

type Booking struct {
	ID            string `json:"bookingId"`
	UserID        int    `json:"userId"`
	CinemaID      int    `json:"cinemaId"`
	SeatID        int    `json:"seatId"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	PaymentMethod string `json:"paymentMethod"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
}
