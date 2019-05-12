package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
)

type RoleService interface {
	GetRoles() ([]*entity.Role, *entity.ErrorResponse)

	GetRoleById(id int) (*entity.Role, *entity.ErrorResponse)

	InsertRole(role *entity.Role) (*entity.Role, *entity.ErrorResponse)
}
