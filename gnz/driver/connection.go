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

type Database struct {
	DbConfig common.DbConfig
}

// Constructor
func NewDatabase() Database {
	return Database{
		DbConfig: common.Db,
	}
}

// Initialize database driver for GrantNZ server
func (r Database) Connect() {
	if !strings.EqualFold(r.DbConfig.Engine, "mysql") {
		panic("Current status, only support mysql.")
	}

	//hosts := strings.Split(r.DbConfig.Hosts, ",")
	//databaseCnt := len(hosts)
	//for _, host := range hosts {
	//
	//}
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		common.Db.User,
		common.Db.Password,
		common.Db.Hosts,
		common.Db.Port,
		common.Db.Name,
	)

	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		log.Logger.Warn(err.Error())
		r.Close()
		panic("Cannot connect MySQL")
	}

	if strings.EqualFold(common.App.LogLevel, "DEBUG") || strings.EqualFold(common.App.LogLevel, "debug") {
		db.LogMode(true)
	} else {
		db.LogMode(false)
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
func (r Database) PingRdbms() {
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
func (r Database) Close() {
	if connection != nil {
		connection.Close()
		log.Logger.Info("Closed MySQL connection")
	} else {
		log.Logger.Info("Already closed MySQL connection")
	}
}
