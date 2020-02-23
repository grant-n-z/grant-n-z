package cache

import (
	"github.com/go-redis/redis"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var rcInstance RedisClient

type RedisClient interface {
	SetPolicy(policy entity.Policy)
}

type RedisClientImpl struct {
	client *redis.Client
}

func GetRedisClientInstance() RedisClient {
	if rcInstance == nil {
		rcInstance = NewRedisClient()
	}
	return rcInstance
}

func NewRedisClient() RedisClient {
	log.Logger.Info("New `RedisClient` instance")
	return RedisClientImpl{client: connection}
}

func (rc RedisClientImpl) SetPolicy(policy entity.Policy) {
}

func (rc RedisClientImpl) SetPolicyWithExpires(policy entity.Policy, expiresMinutes int) {

}
