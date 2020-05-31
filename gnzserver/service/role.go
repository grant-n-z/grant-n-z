package service

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

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

	// Get role by role uuid
	GetRoleByUuid(uuid string) (*entity.Role, *model.ErrorResBody)

	// Get role by name
	GetRoleByName(name string) (*entity.Role, *model.ErrorResBody)

	// Get role by name array
	GetRoleByNames(name []string) ([]entity.Role, *model.ErrorResBody)

	// Get role by group uuid
	// Join group_roles and roles
	GetRolesByGroupUuid(groupUuid string) ([]*entity.Role, *model.ErrorResBody)

	// Insert role
	InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResBody)

	// Insert role with relational data
	InsertWithRelationalData(groupUuid string, role entity.Role) (*entity.Role, *model.ErrorResBody)
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
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return []*entity.Role{}, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return roles, nil
}

func (rs RoleServiceImpl) GetRoleByUuid(uuid string) (*entity.Role, *model.ErrorResBody) {
	role, err := rs.RoleRepository.FindByUuid(uuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found role")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return role, nil
}

func (rs RoleServiceImpl) GetRoleByName(name string) (*entity.Role, *model.ErrorResBody) {
	role, err := rs.RoleRepository.FindByName(name)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found role")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return role, nil
}

func (rs RoleServiceImpl) GetRoleByNames(names []string) ([]entity.Role, *model.ErrorResBody) {

	roles, err := rs.RoleRepository.FindByNames(names)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found roles")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return roles, nil
}

func (rs RoleServiceImpl) GetRolesByGroupUuid(groupUuid string) ([]*entity.Role, *model.ErrorResBody) {
	roles, err := rs.RoleRepository.FindByGroupUuid(groupUuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found roles")
		}
		return nil, model.InternalServerError()
	}

	return roles, nil
}

func (rs RoleServiceImpl) InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResBody) {
	roleId := uuid.New()
	roleIdMd5 := md5.Sum(roleId.NodeID())
	role.Uuid = roleId
	role.InternalId = hex.EncodeToString(roleIdMd5[:])

	savedRole, err := rs.RoleRepository.Save(*role)
	if err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return savedRole, nil
}

func (rs RoleServiceImpl) InsertWithRelationalData(groupUuid string, role entity.Role) (*entity.Role, *model.ErrorResBody) {
	roleId := uuid.New()
	roleIdMd5 := md5.Sum(roleId.NodeID())
	role.Uuid = roleId
	role.InternalId = hex.EncodeToString(roleIdMd5[:])

	savedRole, err := rs.RoleRepository.SaveWithRelationalData(groupUuid, role)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit roles data.")
		}
		return nil, model.InternalServerError()
	}

	return savedRole, nil
}
