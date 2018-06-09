package entity

import "time"

type Users struct {
	Id int `json:"id"`
	Uuid string `json:"uuid"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (Users) TableName() string {
	return "users"
}
