package app

import (
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	"os"
)

var (
	AppVersion string
	BuildTime string
)

var (
	Db *gorm.DB
)

func InitDB() {
	var err error
	var dsn = os.Getenv("DB_USER") +
		":" + os.Getenv("DB_PASS") +
		"@" + os.Getenv("DB_HOST") +
		"/" + os.Getenv("DB_NAME") +
		"?parseTime=true&loc=Asia%2FTokyo"

	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func init() {
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
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