package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type GroupService interface {
	Get(queryParam string) (interface{}, *model.ErrorResBody)

	GetGroups() ([]*entity.Group, *model.ErrorResBody)

	GetGroupByName(name string) (*entity.Group, *model.ErrorResBody)

	InsertGroup(group *entity.Group) (*entity.Group, *model.ErrorResBody)
}
