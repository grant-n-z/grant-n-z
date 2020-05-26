package model

import (
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"

	"github.com/tomoyane/grant-n-z/gnz/entity"
)

// Payload in jwt
type JwtPayload struct {
	UserUuid     string                 `json:"user_uuid"`
	Username     string                 `json:"user_name"`
	ServerId     string                 `json:"server_id"`
	Expires      string                 `json:"expires"`
	IssueDate    string                 `json:"issue_date"`
	UserPolicies []structure.UserPolicy `json:"user_policies"`
	IsRefresh    bool                   `json:"is_refresh"`
}

// The table `users` and `operator_policies` and `roles` struct
type UserWithOperatorPolicy struct {
	entity.User
	entity.OperatorPolicy
	entity.Role
}

// The table `users` and `user_services` and `services` struct
type UserWithUserServiceWithService struct {
	entity.User
	entity.UserService
	entity.Service
}

// Add user id
type AddUser struct {
	UserEmail string `validate:"required"json:"user_email"`
}

// user struct
type UserResponse struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
