package dto

type PaymentRequest struct {
	BookingID      string            `json:"bookingId"`
	PaymentMethod  string            `json:"paymentMethod"`
	PaymentDetails map[string]string `json:"paymentDetails"`
}
