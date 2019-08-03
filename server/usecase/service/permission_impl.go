package service

import (
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/server/model"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/usecase/repository"
)

type permissionServiceImpl struct {
	permissionRepository repository.PermissionRepository
}

func NewPermissionService() PermissionService {
	log.Logger.Info("Inject `permissionRepository` to `PermissionService`")
	return permissionServiceImpl{
		permissionRepository: repository.NewPermissionRepository(config.Db),
	}
}

func (ps permissionServiceImpl) GetPermissions() ([]*entity.Permission, *model.ErrorResponse) {
	permissions, err := ps.permissionRepository.FindAll()
	if permissions == nil {
		return []*entity.Permission{}, err
	}

	return permissions, err
}

func (ps permissionServiceImpl) GetPermissionByRoleId(id int) (*entity.Permission, *model.ErrorResponse) {
	return ps.permissionRepository.FindById(id)
}

func (ps permissionServiceImpl) InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResponse) {
	permission.Uuid, _ = uuid.NewV4()
	return ps.permissionRepository.Save(*permission)
}
