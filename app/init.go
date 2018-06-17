package app

import (
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
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
		db, err := gorm.Open("mysql", getConnection())

		if err != nil {
			revel.ERROR.Println("MySQL open error", err)
			panic(err)
		}

		db.DB()
		Db = db
	}
}

func getConnection() string {
	host := revel.Config.StringDefault("db.host", "0.0.0.0")
	port := revel.Config.StringDefault("db.port", "3306")
	user := revel.Config.StringDefault("db.user", "")
	pass := revel.Config.StringDefault("db.pass", "")
	name := revel.Config.StringDefault("db.name", "")
	protocol := getValueOfParam("db.protocol", "tcp")
	timezone := getValueOfParam("db.timezone", "parseTime=true&loc=Asia%2FTokyo")

	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?%s", user, pass, protocol, host, port, name, timezone)
}

func getValueOfParam(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}