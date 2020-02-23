package cache

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/tomoyane/grant-n-z/gnz/config"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var connection *redis.Client

// Initialize cache database driver
func InitRedis() {
	db, _ := strconv.Atoi(config.Redis.Db)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Logger.Warn(err.Error())
		log.Logger.Warn("Cannot connect Redis. If needs to high performance, run GrantNZ cache server with Redis")
		Close()
		return
	}

	log.Logger.Info("Connected Redis", config.Redis.Host)
	connection = client
}

// Close redis
func Close() {
	if connection != nil {
		connection.Close()
		log.Logger.Info("Closed Redis connection")
	} else {
		log.Logger.Info("Already closed Redis connection")
	}
}
