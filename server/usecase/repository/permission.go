package repository

import "github.com/tomoyane/grant-n-z/server/entity"

type PermissionRepository interface {
	FindAll() ([]*entity.Permission, *entity.ErrorResponse)

	FindById(id int) (*entity.Permission, *entity.ErrorResponse)

	Save(permission entity.Permission) (*entity.Permission, *entity.ErrorResponse)
}
