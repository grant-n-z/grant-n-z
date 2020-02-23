package service

import (
	"encoding/json"
	"time"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/entity"
)

const expiresMinutes = 300 * time.Second

type UpdaterService interface {
	UpdatePolicy(policies []*entity.Policy)

	UpdatePermission() error

	UpdateRole() error

	UpdateService() error
}

type UpdaterServiceImpl struct {
	RedisClient cache.RedisClient
}

func NewUpdaterService() UpdaterService {
	return UpdaterServiceImpl{}
}

func (us UpdaterServiceImpl) UpdatePolicy(policies []*entity.Policy) {
	for _, policy := range policies {
		policy, _ := json.Marshal(policy)
		//driver.Redis.Set("", string(policy), expiresMinutes)
	}
}

func (us UpdaterServiceImpl) UpdatePermission() error {
	panic("implement me")
}

func (us UpdaterServiceImpl) UpdateRole() error {
	panic("implement me")
}

func (us UpdaterServiceImpl) UpdateService() error {
	panic("implement me")
}
