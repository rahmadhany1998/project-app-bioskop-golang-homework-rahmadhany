package entity

type Payment struct {
	ID        string      `json:"transactionId"`
	BookingID string      `json:"bookingId"`
	Method    string      `json:"paymentMethod"`
	Details   interface{} `json:"paymentDetails"`
	Status    string      `json:"status"`
	CreatedAt string      `json:"created_at"`
}
