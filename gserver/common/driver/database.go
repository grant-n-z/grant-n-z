package driver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gserver/common/config"
	"github.com/tomoyane/grant-n-z/gserver/log"
)

// Global DataBase Client
var (
	Db    *gorm.DB
	Redis *redis.Client
)

// Initialize database driver
func InitDriver() {
	if !strings.EqualFold(config.Db.Engine, "mysql") {
		panic("Current status, only support mysql.")
	}

	initDataBase()
	initRedis()
}

// Initialize master database driver
func initDataBase() {
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
		CloseConnection()
		panic("Cannot connect MySQL")
	}

	if strings.EqualFold(config.App.LogLevel, "DEBUG") || strings.EqualFold(config.App.LogLevel, "debug") {
		db.LogMode(true)
	}

	log.Logger.Info("Connected MySQL", config.Db.Host)
	db.DB()
	Db = db
}

// Initialize cache database driver
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
		CloseConnection()
		panic("Cannot connect Redis")
	}

	log.Logger.Info("Connected Redis", config.Redis.Host)
	Redis = client
}

// Close database connection
func CloseConnection() {
	if Db != nil {
		Db.Close()
		log.Logger.Info("Closed MySQL connection")
	} else {
		log.Logger.Info("Already closed MySQL connection")
	}

	if Redis != nil {
		Redis.Close()
		log.Logger.Info("Closed Redis connection")
	} else {
		log.Logger.Info("Already closed Redis connection")
	}
}
