package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/handler"
)

func init() {
	config.InitConfig()
}

func main() {
	fmt.Println("start grant-n-z server :8080")

	userHandler := handler.NewUserHandler()
	serviceHandler := handler.NewServiceHandler()
	roleHandler := handler.NewRoleHandler()

	http.HandleFunc("/api/v1/users",  userHandler.Post)
	http.HandleFunc("/api/v1/services",  serviceHandler.Post)
	http.HandleFunc("/api/v1/roles",  roleHandler.Post)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
