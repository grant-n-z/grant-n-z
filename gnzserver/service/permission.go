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

var psInstance PermissionService

type PermissionService interface {
	// Get all permissions
	GetPermissions() ([]*entity.Permission, *model.ErrorResBody)

	// Get permission by uuid
	GetPermissionByUuid(uuid string) (*entity.Permission, *model.ErrorResBody)

	// Get permission by name
	GetPermissionByName(name string) (*entity.Permission, *model.ErrorResBody)

	// Get permissions by group uuid
	// Join group_permission and permission
	GetPermissionsByGroupUuid(groupUuid string) ([]*entity.Permission, *model.ErrorResBody)

	// Inert permission
	InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResBody)

	// Insert permission with relational data
	InsertWithRelationalData(groupUuid string, permission entity.Permission) (*entity.Permission, *model.ErrorResBody)
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
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return []*entity.Permission{}, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return permissions, nil
}

func (ps PermissionServiceImpl) GetPermissionByUuid(uuid string) (*entity.Permission, *model.ErrorResBody) {
	permission, err := ps.PermissionRepository.FindByUuid(uuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return &entity.Permission{}, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return permission, nil
}

func (ps PermissionServiceImpl) GetPermissionByName(name string) (*entity.Permission, *model.ErrorResBody) {
	permission, err := ps.PermissionRepository.FindByName(name)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found permission")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return permission, nil
}

func (ps PermissionServiceImpl) GetPermissionsByGroupUuid(groupUuid string) ([]*entity.Permission, *model.ErrorResBody) {
	permissions, err := ps.PermissionRepository.FindByGroupUuid(groupUuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found permissions")
		}
		return nil, model.InternalServerError()
	}

	return permissions, nil
}

func (ps PermissionServiceImpl) InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	permissionId := uuid.New()
	permissionMd5 := md5.Sum(permissionId.NodeID())
	permission.Uuid = permissionId
	permission.InternalId = hex.EncodeToString(permissionMd5[:])

	savedPermission, err := ps.PermissionRepository.Save(*permission)
	if err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return savedPermission, nil
}

func (ps PermissionServiceImpl) InsertWithRelationalData(groupUuid string, permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	permissionId := uuid.New()
	permissionMd5 := md5.Sum(permissionId.NodeID())
	permission.Uuid = permissionId
	permission.InternalId = hex.EncodeToString(permissionMd5[:])

	savedData, err := ps.PermissionRepository.SaveWithRelationalData(groupUuid, permission)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit group permission data.")
		}
		return nil, model.InternalServerError()
	}

	return savedData, nil
}
