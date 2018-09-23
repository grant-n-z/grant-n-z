package service

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/satori/go.uuid"
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
