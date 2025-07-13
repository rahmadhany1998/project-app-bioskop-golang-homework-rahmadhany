package entity

type TokenStore struct {
	ID     int    `json:"-"`
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}
