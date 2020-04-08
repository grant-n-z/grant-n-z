package service

import (
	"github.com/google/uuid"
	"github.com/tomoyane/grant-n-z/gnz/cache"

	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var rsInstance RoleService

type RoleService interface {
	// Get all roles
	GetRoles() ([]*entity.Role, *model.ErrorResBody)

	// Get role by role id
	GetRoleById(id int) (*entity.Role, *model.ErrorResBody)

	// Get role by name
	GetRoleByName(name string) (*entity.Role, *model.ErrorResBody)

	// Get role by name array
	GetRoleByNames(name []string) ([]entity.Role, *model.ErrorResBody)

	// Get role by group id
	// Join group_roles and roles
	GetRolesByGroupId(groupId int) ([]*entity.Role, *model.ErrorResBody)

	// Insert role
	InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResBody)

	// Insert role with relational data
	InsertWithRelationalData(groupId int, role entity.Role) (*entity.Role, *model.ErrorResBody)
}

type roleServiceImpl struct {
	etcdClient     cache.EtcdClient
	roleRepository driver.RoleRepository
}

func GetRoleServiceInstance() RoleService {
	if rsInstance == nil {
		rsInstance = NewRoleService()
	}
	return rsInstance
}

func NewRoleService() RoleService {
	log.Logger.Info("New `RoleService` instance")
	return roleServiceImpl{
		etcdClient:     cache.GetEtcdClientInstance(),
		roleRepository: driver.GetRoleRepositoryInstance(),
	}
}

func (rs roleServiceImpl) GetRoles() ([]*entity.Role, *model.ErrorResBody) {
	roles, err := rs.roleRepository.FindAll()
	if roles == nil {
		return []*entity.Role{}, err
	}
	return []*entity.Role{}, nil
}

func (rs roleServiceImpl) GetRoleById(id int) (*entity.Role, *model.ErrorResBody) {
	return rs.roleRepository.FindById(id)
}

func (rs roleServiceImpl) GetRoleByName(name string) (*entity.Role, *model.ErrorResBody) {
	return rs.roleRepository.FindByName(name)
}

func (rs roleServiceImpl) GetRoleByNames(names []string) ([]entity.Role, *model.ErrorResBody) {
	roles := rs.etcdClient.GetRoleByNames(names)
	if len(roles) > 0 {
		return roles, nil
	}

	return rs.roleRepository.FindByNames(names)
}

func (rs roleServiceImpl) GetRolesByGroupId(groupId int) ([]*entity.Role, *model.ErrorResBody) {
	return rs.roleRepository.FindByGroupId(groupId)
}

func (rs roleServiceImpl) InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResBody) {
	role.Uuid = uuid.New()
	return rs.roleRepository.Save(*role)
}

func (rs roleServiceImpl) InsertWithRelationalData(groupId int, role entity.Role) (*entity.Role, *model.ErrorResBody) {
	role.Uuid = uuid.New()
	return rs.roleRepository.SaveWithRelationalData(groupId, role)
}
