package main

import (
	"github.com/tomoyane/grant-n-z/server/common/config"
	"github.com/tomoyane/grant-n-z/server/common/driver"
	"github.com/tomoyane/grant-n-z/server/log"

	gserver "github.com/tomoyane/grant-n-z/server"
)

func init() {
	config.InitConfig()
	log.InitLogger(config.App.LogLevel)
	driver.InitDriver()
}

func main() {
	gserver.NewGrantNZServer().Run()
}
