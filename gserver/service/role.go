package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type RoleService interface {
	GetRoles() ([]*entity.Role, *model.ErrorResBody)

	GetRoleById(id int) (*entity.Role, *model.ErrorResBody)

	InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResBody)
}
