package driver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/server/common/config"
	"github.com/tomoyane/grant-n-z/server/log"
)

var (
	Db    *gorm.DB
	Redis *redis.Client
)

func InitDriver() {
	if !strings.EqualFold(config.Db.Engine, "mysql") {
		panic("Current status, only support mysql.")
	}

	// TODO: Support MongoDB, Cassandra

	initMysql()
	initRedis()
}

func initMysql() {
	if !strings.EqualFold(config.Db.Engine, "mysql") {
		panic("Current status, only support mysql.")
	}

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.Db.User,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Db,
	)

	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		log.Logger.Warn(err.Error())
		panic("Cannot connect MySQL")
	}

	log.Logger.Info("Connected MySQL", config.Db.Host)
	db.DB()
	db.LogMode(false)
	Db = db
}

func initRedis() {
	db, _ := strconv.Atoi(config.Redis.Db)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Logger.Warn(err.Error())
		panic("Cannot connect Redis")
	}

	log.Logger.Info("Connected Redis", config.Redis.Host)
	Redis = client
}
