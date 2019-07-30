package main

import (
	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/log"

	gserver "github.com/tomoyane/grant-n-z/server"
)

func init() {
	config.InitConfig()
	log.InitLogger(config.App.LogLevel)
}

func main() {
	gserver.NewGrantNZServer().Run()
}
