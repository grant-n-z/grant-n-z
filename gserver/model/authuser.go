package model

import "github.com/satori/go.uuid"

// User data in jwt
type AuthUser struct {
	Username string    `json:"username"`
	UserUuid uuid.UUID `json:"user_uuid"`
	UserId   int       `json:"user_id"`
	Expires  string    `json:"expires"`
	RoleId   int       `json:"role_id"`
}
