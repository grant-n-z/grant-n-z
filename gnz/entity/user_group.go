package entity

import (
	"github.com/google/uuid"
	"time"
)

const (
	UserGroupTable UserGroupTableConfig = iota
	UserGroupId
	UserGroupUuid
	UserGroupUserUuid
	UserGroupGroupUuid
	UserGroupCreatedAt
	UserGroupUpdatedAt
)

// The table `user_groups` struct
type UserGroup struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Uuid      uuid.UUID `validate:"required"json:"uuid"`
	UserUuid  uuid.UUID `validate:"required"json:"user_uuid"`
	GroupUuid uuid.UUID `validate:"required"json:"group_uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserGroup table config struct
type UserGroupTableConfig int

func (ugc UserGroupTableConfig) String() string {
	switch ugc {
	case UserGroupTable:
		return "user_groups"
	case UserGroupId:
		return "id"
	case UserGroupUuid:
		return "uuid"
	case UserGroupUserUuid:
		return "user_uuid"
	case UserGroupGroupUuid:
		return "group_uuid"
	case UserGroupCreatedAt:
		return "created_at"
	case UserGroupUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
