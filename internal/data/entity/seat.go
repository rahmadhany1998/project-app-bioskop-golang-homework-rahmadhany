package entity

type Seat struct {
	ID       int    `json:"seatId"`
	CinemaID int    `json:"cinemaId"`
	SeatCode string `json:"seatCode"`
	Status   string `json:"status,omitempty"` // available / booked (optional for availability API)
}
