package entity

type Cinema struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	SeatsCount int    `json:"seats"`
}
