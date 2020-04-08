package driver

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

// Global DataBase Client
var connection *gorm.DB

// Initialize database driver for GrantNZ server
func InitRdbms() {
	if !strings.EqualFold(common.Db.Engine, "mysql") {
		panic("Current status, only support mysql.")
	}

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		common.Db.User,
		common.Db.Password,
		common.Db.Host,
		common.Db.Port,
		common.Db.Db,
	)

	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		log.Logger.Warn(err.Error())
		Close()
		panic("Cannot connect MySQL")
	}

	if strings.EqualFold(common.App.LogLevel, "DEBUG") || strings.EqualFold(common.App.LogLevel, "debug") {
		db.LogMode(true)
	}

	//db.DB().SetMaxOpenConns(10)
	//db.DB().SetMaxIdleConns(10)

	log.Logger.Info(fmt.Sprintf("Connected MySQL. Open connection = %d. Max open connection = %d.",
		db.DB().Stats().OpenConnections,
		db.DB().Stats().MaxOpenConnections),
	)
	connection = db
}

// Ping RDBMS
func PingRdbms() {
	for {
		time.Sleep(1 * time.Minute)
		err := connection.DB().Ping()
		if err != nil {
			log.Logger.Warn("Failed to rdbms ping.", err.Error())
		} else {
			log.Logger.Info("Ping rdbms.")
		}
	}
}

// Close RDBMS
func Close() {
	if connection != nil {
		connection.Close()
		log.Logger.Info("Closed MySQL connection")
	} else {
		log.Logger.Info("Already closed MySQL connection")
	}
}
