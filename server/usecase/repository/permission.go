package repository

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type PermissionRepository interface {
	FindAll() ([]*entity.Permission, *model.ErrorResponse)

	FindById(id int) (*entity.Permission, *model.ErrorResponse)

	Save(permission entity.Permission) (*entity.Permission, *model.ErrorResponse)
}
