package cache

import (
	"github.com/go-redis/redis"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/log"
)

var rcInstance RedisClient

type RedisClient interface {
	Get()

	Set()
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
	return RedisClientImpl{
		client: driver.Redis,
	}
}

func (rc RedisClientImpl) Get() {
}

func (rc RedisClientImpl) Set() {
}
