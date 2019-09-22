package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type PermissionService interface {
	GetPermissions() ([]*entity.Permission, *model.ErrorResBody)

	GetPermissionByRoleId(id int) (*entity.Permission, *model.ErrorResBody)

	InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResBody)
}
