package service

import (
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/repository"
)

type RoleService struct {
	RoleRepository repository.RoleRepository
}

func (r RoleService) GetRoleByUserUuid(userUuid string) *entity.Role {
	return r.RoleRepository.FindByUserUuid(userUuid)
}

func (r RoleService) InsertRole(userUuid uuid.UUID) *entity.Role {
	role := entity.Role{}
	return r.RoleRepository.Save(role)
}
