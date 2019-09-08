package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	GroupId GroupColumn = iota
	GroupUuid
	GroupName
	GroupCreatedAt
	GroupUpdatedAt
)

type Group struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Uuid      uuid.UUID `gorm:"type:varchar(128)"json:"uuid"`
	Name      string    `gorm:"unique;type:varchar(128)"validate:"required"json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GroupColumn int

func (gc GroupColumn) String() string {
	switch gc {
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

func (gc GroupColumn) TableName() string {
	return "groups"
}
