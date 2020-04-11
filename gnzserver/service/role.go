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

type RoleServiceImpl struct {
	EtcdClient     cache.EtcdClient
	RoleRepository driver.RoleRepository
}

func GetRoleServiceInstance() RoleService {
	if rsInstance == nil {
		rsInstance = NewRoleService()
	}
	return rsInstance
}

func NewRoleService() RoleService {
	log.Logger.Info("New `RoleService` instance")
	return RoleServiceImpl{
		EtcdClient:     cache.GetEtcdClientInstance(),
		RoleRepository: driver.GetRoleRepositoryInstance(),
	}
}

func (rs RoleServiceImpl) GetRoles() ([]*entity.Role, *model.ErrorResBody) {
	roles, err := rs.RoleRepository.FindAll()
	if roles == nil {
		return []*entity.Role{}, err
	}
	return []*entity.Role{}, nil
}

func (rs RoleServiceImpl) GetRoleById(id int) (*entity.Role, *model.ErrorResBody) {
	return rs.RoleRepository.FindById(id)
}

func (rs RoleServiceImpl) GetRoleByName(name string) (*entity.Role, *model.ErrorResBody) {
	return rs.RoleRepository.FindByName(name)
}

func (rs RoleServiceImpl) GetRoleByNames(names []string) ([]entity.Role, *model.ErrorResBody) {
	roles := rs.EtcdClient.GetRoleByNames(names)
	if roles != nil && len(roles) > 0 {
		return roles, nil
	}

	return rs.RoleRepository.FindByNames(names)
}

func (rs RoleServiceImpl) GetRolesByGroupId(groupId int) ([]*entity.Role, *model.ErrorResBody) {
	return rs.RoleRepository.FindByGroupId(groupId)
}

func (rs RoleServiceImpl) InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResBody) {
	role.Uuid = uuid.New()
	return rs.RoleRepository.Save(*role)
}

func (rs RoleServiceImpl) InsertWithRelationalData(groupId int, role entity.Role) (*entity.Role, *model.ErrorResBody) {
	role.Uuid = uuid.New()
	return rs.RoleRepository.SaveWithRelationalData(groupId, role)
}
