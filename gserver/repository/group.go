package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type GroupRepository interface {
	FindAll() ([]*entity.Group, *model.ErrorResponse)

	FindByName(name string) (*entity.Group, *model.ErrorResponse)

	Save(group entity.Group) (*entity.Group, *model.ErrorResponse)
}
