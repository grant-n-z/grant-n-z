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

	fmt.Println("start grant-n-z server :8080")

	http.HandleFunc("/api/v1/users",  userHandler.Api)
	http.HandleFunc("/api/v1/services",  serviceHandler.Api)
	http.HandleFunc("/api/v1/roles",  roleHandler.Api)

	log.Logger.Fatal(http.ListenAndServe(":8080", nil))
}
