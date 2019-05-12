package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
)

type PermissionService interface {
	GetPermissions() ([]*entity.Permission, *entity.ErrorResponse)

	GetPermissionByRoleId(id int) (*entity.Permission, *entity.ErrorResponse)

	InsertPermission(permission *entity.Permission) (*entity.Permission, *entity.ErrorResponse)
}
