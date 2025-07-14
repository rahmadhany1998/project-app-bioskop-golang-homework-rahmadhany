package entity

type Seat struct {
	ID       int    `json:"seatId,omitempty"`
	CinemaID int    `json:"cinemaId,omitempty"`
	SeatCode string `json:"seatCode"`
	Status   string `json:"status,omitempty"` // available / booked (optional for availability API)
}
