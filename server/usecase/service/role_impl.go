package service

import (
	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/usecase/repository"
)

type roleServiceImpl struct {
	roleRepository repository.RoleRepository
}

func NewRoleService() RoleService {
	log.Logger.Info("Inject `roleRepository` to `RoleService`")
	return roleServiceImpl{roleRepository: repository.RoleRepositoryImpl{Db: config.Db}}
}

func (rs roleServiceImpl) GetRoles() ([]*entity.Role, *entity.ErrorResponse) {
	roles, err := rs.roleRepository.FindAll()
	if roles == nil {
		return []*entity.Role{}, err
	}
	return []*entity.Role{}, err
}

func (rs roleServiceImpl) GetRoleById(id int) (*entity.Role, *entity.ErrorResponse) {
	return rs.roleRepository.FindById(id)
}

func (rs roleServiceImpl) InsertRole(role *entity.Role) (*entity.Role, *entity.ErrorResponse) {
	role.Uuid, _ = uuid.NewV4()
	return rs.roleRepository.Save(*role)
}
