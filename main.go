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
	http.HandleFunc("/api/v1/users",  userHandler.Post)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
