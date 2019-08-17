package main

import (
	"github.com/tomoyane/grant-n-z/gserver/common/config"
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/log"

	gserver "github.com/tomoyane/grant-n-z/gserver"
)

func init() {
	config.InitConfig()
	log.InitLogger(config.App.LogLevel)
	driver.InitDriver()
}

func main() {
	gserver.NewGrantNZServer().Run()
}
