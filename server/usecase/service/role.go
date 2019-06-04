package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type RoleService interface {
	GetRoles() ([]*entity.Role, *model.ErrorResponse)

	GetRoleById(id int) (*entity.Role, *model.ErrorResponse)

	InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResponse)
}
