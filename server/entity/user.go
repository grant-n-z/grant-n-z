package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	USER_ID UserColumn = iota
	USER_UUID
	USER_USERNAME
	USER_EMAIL
	USER_PASSWORD
	USER_CREATED_AT
	USER_UPDATED_AT
)

type User struct {
	Id        int       `json:"id"`
	Uuid      uuid.UUID `json:"uuid"`
	Username  string    `validate:"required"json:"username"`
	Email     string    `validate:"required,email"json:"email"`
	Password  string    `validate:"min=8,required"json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserColumn int

func (uc UserColumn) String() string {
	switch uc {
	case USER_ID:
		return "id"
	case USER_UUID:
		return "uuid"
	case USER_USERNAME:
		return "username"
	case USER_EMAIL:
		return "email"
	case USER_PASSWORD:
		return "password"
	case USER_CREATED_AT:
		return "created_at"
	case USER_UPDATED_AT:
		return "updated_at"
	}
	panic("Unknown value")
}

func (u User) TableName() string {
	return "users"
}

