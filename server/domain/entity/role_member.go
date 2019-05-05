package entity

import (
	"time"
)

const (
	ROLE_MEMBER_ID RoleMemberColumn = iota
	ROLE_MEMBER_ROLE_ID
	ROLE_MEMBER_USER_ID
	ROLE_MEMBER_CREATED_AT
	ROLE_MEMBER_UPDATED_AT
)

type RoleMember struct {
	Id        int       `json:"id"`
	RoleId    int       `validate:"required"json:"role_id"`
	UserId    int       `validate:"required"json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleMemberColumn int

func (rmc RoleMemberColumn) String() string {
	switch rmc {
	case ROLE_MEMBER_ID:
		return "id"
	case ROLE_MEMBER_ROLE_ID:
		return "role_id"
	case ROLE_MEMBER_USER_ID:
		return "user_id"
	case ROLE_MEMBER_CREATED_AT:
		return "created_at"
	case ROLE_MEMBER_UPDATED_AT:
		return "updated_at"
	}
	panic("Unknown value")
}

func (rm RoleMember) TableName() string {
	return "role_members"
}
