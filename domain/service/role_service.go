package service

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/satori/go.uuid"
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/handler"
)

type RoleService struct {
	RoleRepository repository.RoleRepository
}

func (r RoleService) GetRoleByUserUuid(userUuid string) *entity.Role {
	return r.RoleRepository.FindByUserUuid(userUuid)
}

func (r RoleService) GetRoleByPermission(permission string) *entity.Role {
	return r.RoleRepository.FindByPermission(permission)
}

func (r RoleService) InsertRole(role entity.Role) *entity.Role {
	role.Uuid, _ = uuid.NewV4()
	return r.RoleRepository.Save(role)
}

func (r RoleService) PostRoleData(c echo.Context, role *entity.Role, token string) (insertedRole *entity.Role, errRes *handler.ErrorResponse) {

	if err := c.Bind(role); err != nil {
		return nil, handler.BadRequest("")
	}

	if err := c.Validate(role); err != nil {
		return nil, handler.BadRequest("")
	}

	roleData := r.GetRoleByPermission(role.Permission)
	if roleData == nil {
		return nil, handler.InternalServerError("")
	}

	if len(roleData.Permission) > 0 {
		return nil, handler.Conflict("")
	}

	roleData = r.InsertRole(*role)
	if roleData == nil {
		return nil, handler.InternalServerError("")
	}

	return roleData, nil
}