package service

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var sInstance Service

type ServiceImpl struct {
	EtcdClient           cache.EtcdClient
	ServiceRepository    driver.ServiceRepository
	RoleRepository       driver.RoleRepository
	PermissionRepository driver.PermissionRepository
}

type Service interface {
	// Get service
	GetServices() ([]*entity.Service, *model.ErrorResBody)

	// Get service by service uuid
	GetServiceByUuid(uuid string) (*entity.Service, *model.ErrorResBody)

	// Get service by service name
	GetServiceByName(name string) (*entity.Service, *model.ErrorResBody)

	// Get service by service secret
	GetServiceBySecret(secret string) (*entity.Service, *model.ErrorResBody)

	// Get service of user
	GetServiceByUser(userUuid string) ([]*entity.Service, *model.ErrorResBody)

	// Insert service
	InsertService(service entity.Service) (*entity.Service, *model.ErrorResBody)

	// Insert service
	InsertServiceWithRelationalData(service *entity.Service) (*entity.Service, *model.ErrorResBody)

	// Generate secret
	GenerateSecret() string
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetServiceInstance() Service {
	if sInstance == nil {
		sInstance = NewServiceService()
	}
	return sInstance
}

// Constructor
func NewServiceService() Service {
	log.Logger.Info("New `Service` instance")
	return ServiceImpl{
		EtcdClient:           cache.GetEtcdClientInstance(),
		ServiceRepository:    driver.GetServiceRepositoryInstance(),
		RoleRepository:       driver.GetRoleRepositoryInstance(),
		PermissionRepository: driver.GetPermissionRepositoryInstance(),
	}
}

func (ss ServiceImpl) GetServices() ([]*entity.Service, *model.ErrorResBody) {
	services, err := ss.ServiceRepository.FindAll()
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return []*entity.Service{}, nil
		}
		return nil, model.InternalServerError(err.Error())
	}

	return services, nil
}

func (ss ServiceImpl) GetServiceByUuid(uuid string) (*entity.Service, *model.ErrorResBody) {
	service, err := ss.ServiceRepository.FindByUuid(uuid)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found service")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return service, nil
}

func (ss ServiceImpl) GetServiceByName(name string) (*entity.Service, *model.ErrorResBody) {
	service, err := ss.ServiceRepository.FindByName(name)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found service")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return service, nil
}

func (ss ServiceImpl) GetServiceBySecret(secret string) (*entity.Service, *model.ErrorResBody) {
	service, err := ss.ServiceRepository.FindBySecret(secret)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.BadRequest("Invalid secret")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return service, nil
}

func (ss ServiceImpl) GetServiceByUser(userUuid string) ([]*entity.Service, *model.ErrorResBody) {
	services, err := ss.ServiceRepository.FindServicesByUserUuid(userUuid)
	if err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found services")
		}
		return nil, model.InternalServerError()
	}

	return services, nil
}

func (ss ServiceImpl) InsertService(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	serviceId := uuid.New()
	serviceIdMd5 := md5.Sum(serviceId.NodeID())
	service.Uuid = serviceId
	service.InternalId = hex.EncodeToString(serviceIdMd5[:])
	service.Secret = ss.GenerateSecret()

	savedService, err := ss.ServiceRepository.Save(service)
	if err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}
		return nil, model.InternalServerError(err.Error())
	}

	return savedService, nil
}

func (ss ServiceImpl) InsertServiceWithRelationalData(service *entity.Service) (*entity.Service, *model.ErrorResBody) {
	serviceId := uuid.New()
	serviceIdMd5 := md5.Sum(serviceId.NodeID())
	service.Uuid = serviceId
	service.InternalId = hex.EncodeToString(serviceIdMd5[:])
	service.Secret = ss.GenerateSecret()

	roles, err := ss.RoleRepository.FindByNames([]string{common.AdminRole, common.UserRole})
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found role")
		}
		return nil, model.InternalServerError()
	}

	permissions, err := ss.PermissionRepository.FindByNames([]string{common.AdminPermission, common.ReadPermission, common.WritePermission})
	if err != nil {
		if !strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found permission")
		}
		return nil, model.InternalServerError(err.Error())
	}

	saveWithRelationalData, err := ss.ServiceRepository.SaveWithRelationalData(*service, roles, permissions)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit services data.")
		}
		return nil, model.InternalServerError()
	}

	return saveWithRelationalData, nil
}

func (ss ServiceImpl) GenerateSecret() string {
	key := uuid.New()
	return strings.Replace(key.String(), "-", "", -1)
}
