package service

import (
	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var rsInstance RoleService

type RoleService interface {
	GetRoles() ([]*entity.Role, *model.ErrorResBody)

	GetRoleById(id int) (*entity.Role, *model.ErrorResBody)

	GetRoleByName(name string) (*entity.Role, *model.ErrorResBody)

	InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResBody)
}

type roleServiceImpl struct {
	roleRepository data.RoleRepository
}

func GetRoleServiceInstance() RoleService {
	if rsInstance == nil {
		rsInstance = NewRoleService()
	}
	return rsInstance
}

func NewRoleService() RoleService {
	log.Logger.Info("New `RoleService` instance")
	return roleServiceImpl{roleRepository: data.RoleRepositoryImpl{Db: driver.Db}}
}

func (rs roleServiceImpl) GetRoles() ([]*entity.Role, *model.ErrorResBody) {
	roles, err := rs.roleRepository.FindAll()
	if roles == nil {
		return []*entity.Role{}, err
	}
	return []*entity.Role{}, err
}

func (rs roleServiceImpl) GetRoleById(id int) (*entity.Role, *model.ErrorResBody) {
	return rs.roleRepository.FindById(id)
}

func (rs roleServiceImpl) GetRoleByName(name string) (*entity.Role, *model.ErrorResBody) {
	return rs.roleRepository.FindByName(name)
}

func (rs roleServiceImpl) InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResBody) {
	role.Uuid = uuid.New()
	return rs.roleRepository.Save(*role)
}
