package service

import (
	"strings"

	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gnz/config"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/ctx"
	"github.com/tomoyane/grant-n-z/gnz/data"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var sInstance Service

type serviceImpl struct {
	serviceRepository    data.ServiceRepository
	roleRepository       data.RoleRepository
	permissionRepository data.PermissionRepository
}

type Service interface {
	// Get service
	GetServices() ([]*entity.Service, *model.ErrorResBody)

	// Get service by service id
	GetServiceById(id int) (*entity.Service, *model.ErrorResBody)

	// Get service by service name
	GetServiceByName(name string) (*entity.Service, *model.ErrorResBody)

	// Get service by service api key
	GetServiceOfApiKey() (*entity.Service, *model.ErrorResBody)

	// Get service of user
	GetServiceOfUser() ([]*entity.Service, *model.ErrorResBody)

	// Insert service
	InsertService(service entity.Service) (*entity.Service, *model.ErrorResBody)

	// Insert service
	InsertServiceWithRelationalData(service *entity.Service) (*entity.Service, *model.ErrorResBody)
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
	return serviceImpl{
		serviceRepository:    data.GetServiceRepositoryInstance(driver.Rdbms),
		roleRepository:       data.GetRoleRepositoryInstance(driver.Rdbms),
		permissionRepository: data.GetPermissionRepositoryInstance(driver.Rdbms),
	}
}

func (ss serviceImpl) GetServices() ([]*entity.Service, *model.ErrorResBody) {
	return ss.serviceRepository.FindAll()
}

func (ss serviceImpl) GetServiceById(id int) (*entity.Service, *model.ErrorResBody) {
	return ss.serviceRepository.FindById(id)
}

func (ss serviceImpl) GetServiceByName(name string) (*entity.Service, *model.ErrorResBody) {
	return ss.serviceRepository.FindByName(name)
}

func (ss serviceImpl) GetServiceOfApiKey() (*entity.Service, *model.ErrorResBody) {
	service, err := ss.serviceRepository.FindByApiKey(ctx.GetApiKey().(string))
	if service == nil || err != nil {
		err := model.BadRequest("Api-Key is invalid")
		return nil, err
	}

	return service, nil
}

func (ss serviceImpl) GetServiceOfUser() ([]*entity.Service, *model.ErrorResBody) {
	return ss.serviceRepository.FindServicesByUserId(ctx.GetUserId().(int))
}

func (ss serviceImpl) InsertService(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	service.Uuid = uuid.New()
	key := uuid.New()
	service.ApiKey = strings.Replace(key.String(), "-", "", -1)
	return ss.serviceRepository.Save(service)
}

func (ss serviceImpl) InsertServiceWithRelationalData(service *entity.Service) (*entity.Service, *model.ErrorResBody) {
	service.Uuid = uuid.New()
	key := uuid.New()
	service.ApiKey = strings.Replace(key.String(), "-", "", -1)

	// TODO: Cache roles
	roles, err := ss.roleRepository.FindByNames([]string{config.AdminRole, config.UserRole})
	if err != nil {
		log.Logger.Info("Failed to get role for insert groups process")
		return nil, model.InternalServerError()
	}

	// TODO: Cache permissions
	permissions, err := ss.permissionRepository.FindByNames([]string{config.AdminPermission, config.ReadPermission, config.WritePermission})
	if err != nil {
		log.Logger.Info("Failed to get permission for insert groups process")
		return nil, model.InternalServerError()
	}

	return ss.serviceRepository.SaveWithRelationalData(*service, roles, permissions)
}
