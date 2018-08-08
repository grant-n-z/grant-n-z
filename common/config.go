package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	Db *gorm.DB
)

func InitDB() {
	db, err := gorm.Open("mysql", os.Getenv("DB_SOURCE"))
	if err != nil {
		panic(err)
	}

	db.DB()
	Db = db
}