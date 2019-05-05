package service

import (
	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/repository"
	"github.com/tomoyane/grant-n-z/server/log"
)

type RoleService struct {
	RoleRepository repository.RoleRepository
}

func NewRoleService() RoleService {
	log.Logger.Info("inject `RoleRepository` to `RoleService`")
	return RoleService{RoleRepository: repository.RoleRepositoryImpl{}}
}

func (rs RoleService) GetRoles() ([]*entity.Role, *entity.ErrorResponse) {
	return rs.RoleRepository.FindAll()
}

func (rs RoleService) GetRoleById(id int) (*entity.Role, *entity.ErrorResponse) {
	return rs.RoleRepository.FindById(id)
}

func (rs RoleService) InsertRole(role *entity.Role) (*entity.Role, *entity.ErrorResponse) {
	role.Uuid, _ = uuid.NewV4()
	return rs.RoleRepository.Save(*role)
}
