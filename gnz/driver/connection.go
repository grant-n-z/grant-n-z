package driver

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/config"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

// Global DataBase Client
var (
	connection *gorm.DB
)

// Initialize database driver for GrantNZ server
func InitRdbms() {
	if !strings.EqualFold(config.Db.Engine, "mysql") {
		panic("Current status, only support mysql.")
	}

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
		Close()
		panic("Cannot connect MySQL")
	}

	if strings.EqualFold(config.App.LogLevel, "DEBUG") || strings.EqualFold(config.App.LogLevel, "debug") {
		db.LogMode(true)
	}

	log.Logger.Info("Connected MySQL", config.Db.Host)
	db.DB()
	connection = db
}

func Close() {
	if connection != nil {
		connection.Close()
		log.Logger.Info("Closed MySQL connection")
	} else {
		log.Logger.Info("Already closed MySQL connection")
	}
}
