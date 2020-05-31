package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	GroupTable GroupTableConfig = iota
	GroupId
	GroupInternalId
	GroupUuid
	GroupName
	GroupCreatedAt
	GroupUpdatedAt
)

// The table `groups` struct
type Group struct {
	Id         int       `json:"id"`
	InternalId string    `json:"internal_id"`
	Uuid       uuid.UUID `json:"uuid"`
	Name       string    `validate:"required"json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Group table config struct
type GroupTableConfig int

func (gc GroupTableConfig) String() string {
	switch gc {
	case GroupTable:
		return "groups"
	case GroupId:
		return "id"
	case GroupInternalId:
		return "internal_id"
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
