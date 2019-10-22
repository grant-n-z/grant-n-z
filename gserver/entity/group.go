package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	GroupTable GroupTableConfig = iota
	GroupId
	GroupUuid
	GroupName
	GroupCreatedAt
	GroupUpdatedAt
)

// The table `groups` struct
type Group struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `gorm:"type:varchar(128)"validate:"required"json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// The table `groups` and `user_groups` struct
type GroupWithUserGroup struct {
	Group
	UserGroup
}

type GroupTableConfig int

func (gc GroupTableConfig) String() string {
	switch gc {
	case GroupTable:
		return "groups"
	case GroupId:
		return "id"
	case GroupUuid:
		return "uuid"
	case GroupName:
		return "name"
	case GroupCreatedAt:
		return "created_at"
	case GroupUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
