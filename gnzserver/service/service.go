package service

import (
	"strings"

	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
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

	// Get service by service id
	GetServiceById(id int) (*entity.Service, *model.ErrorResBody)

	// Get service by service name
	GetServiceByName(name string) (*entity.Service, *model.ErrorResBody)

	// Get service by service secret
	GetServiceOfSecret() (*entity.Service, *model.ErrorResBody)

	// Get service of user
	GetServiceOfUser() ([]*entity.Service, *model.ErrorResBody)

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
	return ss.ServiceRepository.FindAll()
}

func (ss ServiceImpl) GetServiceById(id int) (*entity.Service, *model.ErrorResBody) {
	return ss.ServiceRepository.FindById(id)
}

func (ss ServiceImpl) GetServiceByName(name string) (*entity.Service, *model.ErrorResBody) {
	return ss.ServiceRepository.FindByName(name)
}

func (ss ServiceImpl) GetServiceOfSecret() (*entity.Service, *model.ErrorResBody) {
	service := ss.EtcdClient.GetService(ctx.GetClientSecret().(string))
	if service != nil {
		return service, nil
	}

	service, err := ss.ServiceRepository.FindBySecret(ctx.GetClientSecret().(string))
	if service == nil || err != nil {
		err := model.BadRequest("Client-Secret is invalid")
		return nil, err
	}

	return service, nil
}

func (ss ServiceImpl) GetServiceOfUser() ([]*entity.Service, *model.ErrorResBody) {
	return ss.ServiceRepository.FindServicesByUserId(ctx.GetUserId().(int))
}

func (ss ServiceImpl) InsertService(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	service.Uuid = uuid.New()
	service.Secret = ss.GenerateSecret()
	return ss.ServiceRepository.Save(service)
}

func (ss ServiceImpl) InsertServiceWithRelationalData(service *entity.Service) (*entity.Service, *model.ErrorResBody) {
	service.Uuid = uuid.New()
	service.Secret = ss.GenerateSecret()

	defaultRoles := []string{common.AdminRole, common.UserRole}
	roles := ss.EtcdClient.GetRoleByNames(defaultRoles)
	if roles == nil || len(roles) == 0 {
		masterRoles, err := ss.RoleRepository.FindByNames([]string{common.AdminRole, common.UserRole})
		if err != nil {
			log.Logger.Info("Failed to get role for insert groups process")
			return nil, model.InternalServerError()
		}
		roles = masterRoles
	}

	defaultPermissions := []string{common.AdminPermission, common.ReadPermission, common.WritePermission}
	permissions := ss.EtcdClient.GetPermissionByNames(defaultPermissions)
	if permissions == nil || len(permissions) == 0 {
		masterPermissions, err := ss.PermissionRepository.FindByNames(defaultPermissions)
		if err != nil {
			log.Logger.Info("Failed to get permission for insert groups process")
			return nil, model.InternalServerError()
		}
		permissions = masterPermissions
	}

	return ss.ServiceRepository.SaveWithRelationalData(*service, roles, permissions)
}

func (ss ServiceImpl) GenerateSecret() string {
	key := uuid.New()
	return strings.Replace(key.String(), "-", "", -1)
}
