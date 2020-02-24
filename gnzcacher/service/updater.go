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
}

type UpdaterServiceImpl struct {
	RedisClient cache.RedisClient
}

func NewUpdaterService() UpdaterService {
	return UpdaterServiceImpl{RedisClient: cache.NewRedisClient()}
}

func (us UpdaterServiceImpl) UpdatePolicy(policies []*entity.Policy) {
	for _, policy := range policies {
		us.RedisClient.SetPolicy(*policy, expiresMinutes)
	}
}

func (us UpdaterServiceImpl) UpdatePermission(permissions []*entity.Permission) {
	for _, permission := range permissions {
		us.RedisClient.SetPermission(*permission, expiresMinutes)
	}
}

func (us UpdaterServiceImpl) UpdateRole(roles []*entity.Role) {
	for _, role := range roles {
		us.RedisClient.SetRole(*role, expiresMinutes)
	}
}

func (us UpdaterServiceImpl) UpdateService(services []*entity.Service) {
	for _, service := range services {
		us.RedisClient.SetService(*service, expiresMinutes)
	}
}
