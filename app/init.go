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
	db, err := gorm.Open("mysql", getConnectionString())

	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}

	db.DB()
	Db = db
}

func getConnectionString() string {
	host := getParamString("db.host", os.Getenv("DB_HOST"))
	port := getParamString("db.port", os.Getenv("DB_PORT"))
	user := getParamString("db.user", os.Getenv("DB_USER"))
	pass := getParamString("db.password", os.Getenv("DB_PASS"))
	dbname := getParamString("db.name", os.Getenv("DB_NAME"))
	protocol := getParamString("db.protocol", "tcp")
	timezone := getParamString("db.timezone", "parseTime=true&loc=Asia%2FTokyo")

	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?%s", user, pass, protocol, host, port, dbname, timezone)
}

func getParamString(param string, defaultValue string) string {
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