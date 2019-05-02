package main

import (
	"fmt"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/handler"
	"github.com/tomoyane/grant-n-z/server/log"
)

func init() {
	config.InitConfig()
	log.InitLogger(config.LogLevel)
}

func main() {
	userHandler := handler.NewUserHandler()
	serviceHandler := handler.NewServiceHandler()
	roleHandler := handler.NewRoleHandler()

	banner := `start grant-n-z server :8080
___________________________________________________
    ____                      _      
   / __/ _    ____   _____ __//_      _____   ____ 
  / /__ //__ /__ /  /___ //_ __/     /___ /  /__ /
 / /_ //___///_//_ //  //  //_  === //  // === //__
/____///   /_____///  //  /__/     //  //     /___/
___________________________________________________
High performance authentication and authorization. version is %s
`
	fmt.Printf(banner, config.AppVersion)

	http.HandleFunc("/api/v1/users",  userHandler.Api)
	http.HandleFunc("/api/v1/services",  serviceHandler.Api)
	http.HandleFunc("/api/v1/roles",  roleHandler.Api)

	log.Logger.Fatal(http.ListenAndServe(":8080", nil))
}
