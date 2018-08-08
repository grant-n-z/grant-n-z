package main

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/common"
	"github.com/tomoyane/grant-n-z/controller"
)

func main() {
	common.InitDB()
	e := echo.New()

	e.POST("/v1/users", controller.UserController)

	e.Logger.Fatal(e.Start(":8080"))
}