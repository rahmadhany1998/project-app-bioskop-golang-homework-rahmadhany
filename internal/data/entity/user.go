package entity

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	CreatedAt string `json:"created_at"`
}
