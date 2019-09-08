package entity

import (
	"time"
)

const (
	UserGroupTable UserGroupTableConfig = iota
	UserGroupId
	UserGroupUserId
	UserGroupGroupId
	UserGroupCreatedAt
	UserGroupUpdatedAt
)

type UserGroup struct {
	Id        int       `gorm:"primary_key"json:"id"`
	UserId    int       `validate:"required"json:"user_id"`
	GroupId   int       `validate:"required"json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserGroupTableConfig int

func (ugc UserGroupTableConfig) String() string {
	switch ugc {
	case UserGroupTable:
		return "user_groups"
	case UserGroupId:
		return "id"
	case UserGroupUserId:
		return "user_id"
	case UserGroupGroupId:
		return "group_id"
	case UserGroupCreatedAt:
		return "created_at"
	case UserGroupUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
