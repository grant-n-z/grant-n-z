package app

import (
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"os"
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
		revel.SessionFilter,
		revel.FlashFilter,
		revel.ValidationFilter,
		revel.I18nFilter,
		revel.InterceptorFilter,
		revel.CompressFilter,
		revel.ActionInvoker,
	}

	revel.OnAppStart(InitDB)
}

func InitDB() {
	db, err := gorm.Open("mysql", GetConnection())

	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}

	db.DB()
	Db = db
}

func GetConnection() string {
	host := GetValueOfParam("db.host", os.Getenv("DB_HOST"))
	port := GetValueOfParam("db.port", os.Getenv("DB_PORT"))
	user := GetValueOfParam("db.user", os.Getenv("DB_USER"))
	pass := GetValueOfParam("db.password", os.Getenv("DB_PASS"))
	name := GetValueOfParam("db.name", os.Getenv("DB_NAME"))
	protocol := GetValueOfParam("db.protocol", "tcp")
	timezone := GetValueOfParam("db.timezone", "parseTime=true&loc=Asia%2FTokyo")

	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?%s", user, pass, protocol, host, port, name, timezone)
}

func GetValueOfParam(param string, defaultValue string) string {
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