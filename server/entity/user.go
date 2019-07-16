package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	UserId UserColumn = iota
	UserUuid
	UserUsername
	UserEmail
	UserPassword
	UserCreatedAt
	UserUpdatedAt
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
	case UserId:
		return "id"
	case UserUuid:
		return "uuid"
	case UserUsername:
		return "username"
	case UserEmail:
		return "email"
	case UserPassword:
		return "password"
	case UserCreatedAt:
		return "created_at"
	case UserUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}

func (u User) TableName() string {
	return "users"
}

