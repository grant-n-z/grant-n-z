package driver

import (
	"fmt"
	"strconv"
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
	AppConfig common.AppConfig
}

// Constructor
func NewDatabase() Database {
	return Database{
		DbConfig: common.Db,
		AppConfig: common.App,
	}
}

// Initialize database driver for GrantNZ server
func (r Database) Connect() {
	if !strings.EqualFold(r.DbConfig.Engine, "mysql") {
		panic("Current status, only support mysql.")
	}

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		r.DbConfig.User,
		r.DbConfig.Password,
		r.DbConfig.Hosts,
		r.DbConfig.Port,
		r.DbConfig.Name,
	)

	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		log.Logger.Warn(err.Error())
		r.Close()
		panic("Cannot connect MySQL")
	}

	if strings.EqualFold(r.AppConfig.LogLevel, "DEBUG") || strings.EqualFold(r.AppConfig.LogLevel, "debug") {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

	openConnection, _ := strconv.Atoi(r.DbConfig.MaxOpenConnection)
	idleConnection, _ := strconv.Atoi(r.DbConfig.MaxIdleConnection)
	db.DB().SetMaxOpenConns(openConnection)
	db.DB().SetMaxIdleConns(idleConnection)

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
