package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type PermissionRepository interface {
	FindAll() ([]*entity.Permission, *model.ErrorResBody)

	FindById(id int) (*entity.Permission, *model.ErrorResBody)

	Save(permission entity.Permission) (*entity.Permission, *model.ErrorResBody)
}
