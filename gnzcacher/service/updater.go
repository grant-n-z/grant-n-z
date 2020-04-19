package service

import (
	"time"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/entity"
)

// 10 minute cache expires
const expiresMinutes = 600 * time.Second

type UpdaterService interface {
	// Update policy cache
	UpdatePolicy(policies []*entity.Policy)

	// Update permission cache
	UpdatePermission(permissions []*entity.Permission)

	// Update role cache
	UpdateRole(roles []*entity.Role)

	// Update service cache
	UpdateService(services []*entity.Service)

	// Update user_service cache
	UpdateUserService(userServices []*entity.UserService)
}

type UpdaterServiceImpl struct {
	EtcdClient cache.EtcdClient
}

func NewUpdaterService() UpdaterService {
	return UpdaterServiceImpl{EtcdClient: cache.NewEtcdClient()}
}

func (us UpdaterServiceImpl) UpdatePolicy(policies []*entity.Policy) {
	for _, policy := range policies {
		us.EtcdClient.SetPolicy(*policy, expiresMinutes)
	}
}

func (us UpdaterServiceImpl) UpdatePermission(permissions []*entity.Permission) {
	for _, permission := range permissions {
		us.EtcdClient.SetPermission(*permission, expiresMinutes)
	}
}

func (us UpdaterServiceImpl) UpdateRole(roles []*entity.Role) {
	for _, role := range roles {
		us.EtcdClient.SetRole(*role, expiresMinutes)
	}
}

func (us UpdaterServiceImpl) UpdateService(services []*entity.Service) {
	for _, service := range services {
		us.EtcdClient.SetService(*service, expiresMinutes)
	}
}

func (us UpdaterServiceImpl) UpdateUserService(userServices []*entity.UserService) {
	userServiceMap := map[int][]entity.UserService{}
	for _, userService := range userServices {
		savedUserServices := userServiceMap[userService.UserId]
		if savedUserServices == nil {
			var userServiceArray []entity.UserService
			userServiceArray = append(userServiceArray, *userService)
			userServiceMap[userService.UserId] = userServiceArray
		} else {
			savedUserServices = append(savedUserServices, *userService)
			userServiceMap[userService.UserId] = savedUserServices
		}
	}

	for key, value := range userServiceMap {
		us.EtcdClient.SetUserService(key, value, expiresMinutes)
	}
}
