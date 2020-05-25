package service

import (
	"time"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
)

// 10 minute cache expires
const expiresMinutes = 600 * time.Second

type UpdaterService interface {
	// Update policy cache
	UpdatePolicy(policyMap map[string][]structure.UserPolicy)

	// Update permission cache
	UpdatePermission(permissions []structure.Permission)

	// Update role cache
	UpdateRole(roles []structure.Role)

	// Update service cache
	UpdateService(services []structure.Service)

	// Update user_service cache
	UpdateUserService(serviceMap map[string][]structure.UserService)

	// Update user_group cache
	UpdateUserGroup(groupMap map[string][]structure.UserGroup)
}

type UpdaterServiceImpl struct {
	EtcdClient cache.EtcdClient
}

func NewUpdaterService() UpdaterService {
	return UpdaterServiceImpl{EtcdClient: cache.NewEtcdClient()}
}

func (us UpdaterServiceImpl) UpdatePolicy(policyMap map[string][]structure.UserPolicy) {
	for key, value := range policyMap {
		us.EtcdClient.SetUserPolicy(key, value)
	}
}

func (us UpdaterServiceImpl) UpdatePermission(permissions []structure.Permission) {
	for _, permission := range permissions {
		us.EtcdClient.SetPermission(permission.Uuid, permission)
	}
}

func (us UpdaterServiceImpl) UpdateRole(roles []structure.Role) {
	for _, role := range roles {
		us.EtcdClient.SetRole(role.Uuid, role)
	}
}

func (us UpdaterServiceImpl) UpdateService(services []structure.Service) {
	for _, service := range services {
		us.EtcdClient.SetService(service.Uuid, service)
	}
}

func (us UpdaterServiceImpl) UpdateUserService(serviceMap map[string][]structure.UserService) {
	for key, value := range serviceMap {
		us.EtcdClient.SetUserService(key, value)
	}
}

func (us UpdaterServiceImpl) UpdateUserGroup(policyMap map[string][]structure.UserGroup) {
	for key, value := range policyMap {
		us.EtcdClient.SetUserGroup(key, value)
	}
}
