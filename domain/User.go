package domain

import "time"

type User struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Username  string    `validate:"required"json:"username"`
	Email     string    `validate:"required,email"json:"email"`
	Password  string    `validate:"required"json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
