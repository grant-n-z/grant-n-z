package entity

import (
	"time"
)

const (
	GroupServiceId GroupServiceColumn = iota
	GroupServiceUserId
	GroupServiceGroupId
	GroupServiceCreatedAt
	GroupServiceUpdatedAt
)

type GroupService struct {
	Id        int       `gorm:"primary_key"json:"id"`
	UserId    int       `validate:"required"json:"user_id"`
	GroupId   int       `validate:"required"json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GroupServiceColumn int

func (gsc GroupServiceColumn) String() string {
	switch gsc {
	case GroupServiceId:
		return "id"
	case GroupServiceUserId:
		return "user_id"
	case GroupServiceGroupId:
		return "group_id"
	case GroupServiceCreatedAt:
		return "created_at"
	case GroupServiceUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}

func (gsc GroupServiceColumn) TableName() string {
	return "groups"
}
