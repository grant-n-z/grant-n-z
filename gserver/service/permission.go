package service

import (
	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var psInstance PermissionService

type PermissionService interface {
	GetPermissions() ([]*entity.Permission, *model.ErrorResBody)

	GetPermissionById(id int) (*entity.Permission, *model.ErrorResBody)

	GetPermissionByName(name string) (*entity.Permission, *model.ErrorResBody)

	InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResBody)
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
	log.Logger.Info("Inject `PermissionRepository` to `PermissionService`")
	return permissionServiceImpl{
		permissionRepository: data.GetPermissionRepositoryInstance(driver.Db),
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

func (ps permissionServiceImpl) InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	permission.Uuid = uuid.New()
	return ps.permissionRepository.Save(*permission)
}
