package app

import (
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	Db *gorm.DB
)

func init() {
	revel.Filters = []revel.Filter{
		revel.PanicFilter,
		revel.RouterFilter,
		revel.FilterConfiguringFilter,
		revel.ParamsFilter,
		revel.FlashFilter,
		revel.ValidationFilter,
		revel.I18nFilter,
		revel.InterceptorFilter,
		revel.CompressFilter,
		revel.ActionInvoker,
	}

	revel.OnAppStart(initDB)
}

func initDB() {
	var mode = revel.Config.BoolDefault("mode.dev", false)
	if mode == true {
		db, err := gorm.Open("mysql", os.Getenv("DB_SOURCE"))

		if err != nil {
			revel.ERROR.Println("MySQL open error", err)
			panic(err)
		}

		db.DB()
		Db = db
	}
}