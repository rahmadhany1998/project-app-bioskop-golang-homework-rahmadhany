package dto

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}
