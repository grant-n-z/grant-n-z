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
	// Get all roles
	GetRoles() ([]*entity.Role, *model.ErrorResBody)

	// Get role by role id
	GetRoleById(id int) (*entity.Role, *model.ErrorResBody)

	// Get role by name
	GetRoleByName(name string) (*entity.Role, *model.ErrorResBody)

	// Get role by name array
	GetRoleByNames(name []string) ([]*entity.Role, *model.ErrorResBody)

	// Get role by group id
	// Join group_roles and roles
	GetRolesByGroupId(groupId int) ([]*entity.Role, *model.ErrorResBody)

	// Insert role
	InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResBody)

	// Insert role with relational data
	InsertWithRelationalData(groupId int, role entity.Role) (*entity.Role, *model.ErrorResBody)
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
	return roleServiceImpl{roleRepository: data.GetRoleRepositoryInstance(driver.Db)}
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

func (rs roleServiceImpl) GetRoleByNames(name []string) ([]*entity.Role, *model.ErrorResBody) {
	return rs.roleRepository.FindByNames(name)
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
