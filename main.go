package main

import (
	"fmt"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/handler"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/router"
)

func init() {
	config.InitConfig()
	log.InitLogger(config.LogLevel)
	log.Logger.Debug("Completed init process")
}

func main() {
	route := router.Router{
		UserHandler:        handler.NewUserHandler(),
		ServiceHandler:     handler.NewServiceHandler(),
		RoleHandler:        handler.NewRoleHandler(),
		RoleMemberHandler:  handler.NewRoleMemberHandler(),
		UserServiceHandler: handler.NewUserServiceHandler(),
		PermissionHandler:  handler.NewPermissionHandler(),
		PolicyHandler:      handler.NewPolicyHandlerHandler(),
	}

	route.V1()

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

	route.Run(":8080")
}
