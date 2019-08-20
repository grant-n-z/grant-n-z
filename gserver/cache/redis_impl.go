package cache

import (
	"github.com/go-redis/redis"
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
)

type RedisClientImpl struct {
	client *redis.Client
}

func NewRedisClient() RedisClient {
	return RedisClientImpl{
		client: driver.Redis,
	}
}

func (rc RedisClientImpl) Get() {
}

func (rc RedisClientImpl) Set() {
}
