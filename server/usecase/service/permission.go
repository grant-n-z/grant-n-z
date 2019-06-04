package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type PermissionService interface {
	GetPermissions() ([]*entity.Permission, *model.ErrorResponse)

	GetPermissionByRoleId(id int) (*entity.Permission, *model.ErrorResponse)

	InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResponse)
}
