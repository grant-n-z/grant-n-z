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

func (r RoleService) InsertRole(userUuid uuid.UUID) *entity.Role {
	role := entity.Role{
		Type: "user",
		UserUuid: userUuid,
	}
	return r.RoleRepository.Save(role)
}
