package service

import (
	"github.com/google/uuid"
	"github.com/tomoyane/grant-n-z/gnz/cache"

	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var psInstance PermissionService

type PermissionService interface {
	// Get all permissions
	GetPermissions() ([]*entity.Permission, *model.ErrorResBody)

	// Get permission by id
	GetPermissionById(id int) (*entity.Permission, *model.ErrorResBody)

	// Get permission by name
	GetPermissionByName(name string) (*entity.Permission, *model.ErrorResBody)

	// Get permissions by group id
	// Join group_permission and permission
	GetPermissionsByGroupId(groupId int) ([]*entity.Permission, *model.ErrorResBody)

	// Inert permission
	InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResBody)

	// Insert permission with relational data
	InsertWithRelationalData(groupId int, permission entity.Permission) (*entity.Permission, *model.ErrorResBody)
}

type PermissionServiceImpl struct {
	EtcdClient           cache.EtcdClient
	PermissionRepository driver.PermissionRepository
}

func GetPermissionServiceInstance() PermissionService {
	if psInstance == nil {
		psInstance = NewPermissionService()
	}
	return psInstance
}

func NewPermissionService() PermissionService {
	log.Logger.Info("New `PermissionService` instance")
	return PermissionServiceImpl{
		EtcdClient:           cache.GetEtcdClientInstance(),
		PermissionRepository: driver.GetPermissionRepositoryInstance(),
	}
}

func (ps PermissionServiceImpl) GetPermissions() ([]*entity.Permission, *model.ErrorResBody) {
	permissions, err := ps.PermissionRepository.FindAll()
	if permissions == nil {
		return []*entity.Permission{}, err
	}

	return permissions, nil
}

func (ps PermissionServiceImpl) GetPermissionById(id int) (*entity.Permission, *model.ErrorResBody) {
	return ps.PermissionRepository.FindById(id)
}

func (ps PermissionServiceImpl) GetPermissionByName(name string) (*entity.Permission, *model.ErrorResBody) {
	permission := ps.EtcdClient.GetPermission(name)
	if permission != nil {
		return permission, nil
	}
	return ps.PermissionRepository.FindByName(name)
}

func (ps PermissionServiceImpl) GetPermissionsByGroupId(groupId int) ([]*entity.Permission, *model.ErrorResBody) {
	return ps.PermissionRepository.FindByGroupId(groupId)
}

func (ps PermissionServiceImpl) InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	permission.Uuid = uuid.New()
	return ps.PermissionRepository.Save(*permission)
}

func (ps PermissionServiceImpl) InsertWithRelationalData(groupId int, permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	permission.Uuid = uuid.New()
	return ps.PermissionRepository.SaveWithRelationalData(groupId, permission)
}
