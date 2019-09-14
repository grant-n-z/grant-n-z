package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type GroupService interface {
	Get(queryParam string) (interface{}, *model.ErrorResponse)

	GetGroups() ([]*entity.Group, *model.ErrorResponse)

	GetGroupByName(name string) (*entity.Group, *model.ErrorResponse)

	InsertGroup(group *entity.Group) (*entity.Group, *model.ErrorResponse)
}
