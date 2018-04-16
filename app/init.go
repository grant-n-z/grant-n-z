package app

import (
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strings"
	"os"
)

var (
	AppVersion string
	BuildTime string
)

var (
	Db *gorm.DB
)

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

func getConnectionString() string {
	host := getParamString("db.host", os.Getenv("DB_HOST"))
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", os.Getenv("DB_USER"))
	pass := getParamString("db.password", os.Getenv("DB_PASS"))
	dbname := getParamString("db.name", os.Getenv("DB_NAME"))
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", " ")
	timezone := getParamString("db.timezone", "parseTime=true&loc=Asia%2FTokyo")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s?%s", user, pass, protocol, host, port, dbname, dbargs, timezone)
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
		HeaderFilter,
		revel.InterceptorFilter,
		revel.CompressFilter,
		revel.ActionInvoker,
	}

	revel.OnAppStart(InitDB)
}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:])
}