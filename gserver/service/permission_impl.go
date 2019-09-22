package service

import (
	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/repository"
)

var psInstance PermissionService

type permissionServiceImpl struct {
	permissionRepository repository.PermissionRepository
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
		permissionRepository: repository.GetPermissionRepositoryInstance(driver.Db),
	}
}

func (ps permissionServiceImpl) GetPermissions() ([]*entity.Permission, *model.ErrorResBody) {
	permissions, err := ps.permissionRepository.FindAll()
	if permissions == nil {
		return []*entity.Permission{}, err
	}

	return permissions, err
}

func (ps permissionServiceImpl) GetPermissionByRoleId(id int) (*entity.Permission, *model.ErrorResBody) {
	return ps.permissionRepository.FindById(id)
}

func (ps permissionServiceImpl) InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	permission.Uuid, _ = uuid.NewV4()
	return ps.permissionRepository.Save(*permission)
}
