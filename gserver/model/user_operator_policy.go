package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// User entity with OperatorPolicy entity
type UserOperatorPolicy struct {
	Id        int       `json:"id"`
	Uuid      uuid.UUID `json:"uuid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RoleId    int       `json:"role_id"`
	UserId    int       `json:"user_id"`
}
