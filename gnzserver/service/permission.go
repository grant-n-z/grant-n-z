package service

import (
	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnz/data"
	"github.com/tomoyane/grant-n-z/gnz/entity"
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

type permissionServiceImpl struct {
	permissionRepository data.PermissionRepository
}

func GetPermissionServiceInstance() PermissionService {
	if psInstance == nil {
		psInstance = NewPermissionService()
	}
	return psInstance
}

func NewPermissionService() PermissionService {
	log.Logger.Info("New `PermissionService` instance")
	return permissionServiceImpl{
		permissionRepository: data.GetPermissionRepositoryInstance(driver.Rdbms),
	}
}

func (ps permissionServiceImpl) GetPermissions() ([]*entity.Permission, *model.ErrorResBody) {
	permissions, err := ps.permissionRepository.FindAll()
	if permissions == nil {
		return []*entity.Permission{}, err
	}

	return permissions, err
}

func (ps permissionServiceImpl) GetPermissionById(id int) (*entity.Permission, *model.ErrorResBody) {
	return ps.permissionRepository.FindById(id)
}

func (ps permissionServiceImpl) GetPermissionByName(name string) (*entity.Permission, *model.ErrorResBody) {
	return ps.permissionRepository.FindByName(name)
}

func (ps permissionServiceImpl) GetPermissionsByGroupId(groupId int) ([]*entity.Permission, *model.ErrorResBody) {
	return ps.permissionRepository.FindByGroupId(groupId)
}

func (ps permissionServiceImpl) InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	permission.Uuid = uuid.New()
	return ps.permissionRepository.Save(*permission)
}

func (ps permissionServiceImpl) InsertWithRelationalData(groupId int, permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	permission.Uuid = uuid.New()
	return ps.permissionRepository.SaveWithRelationalData(groupId, permission)
}
