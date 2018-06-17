package controllers

import (
	"authentication-server/tests"
	"github.com/lestrrat/go-test-mysqld"
	"log"
	"authentication-server/app/controllers/v1"
	"github.com/jinzhu/gorm"
	"authentication-server/app/domains/entity"
	"authentication-server/app"
	"fmt"
)

type UserControllerTest struct {
	tests.AppTest
}

var testMySQLd *mysqltest.TestMysqld
var userController = v1.UserController{}

// Start MySQL
func (t UserControllerTest) Before() {

	testMySQLd, err := mysqltest.NewMysqld(nil)
	if err != nil {
		log.Fatal("Test MySQL server", err)
	}
	testMySQLd.Start()

	var dataSource = testMySQLd.Datasource("auth_server", "test", "test", 3306)
	var driver = "mysql"

	db, err := gorm.Open(driver, dataSource)
	if err != nil {
		log.Fatal("Dtabase connection error:", err)
	}

	db.DB()
	app.Db = db

	app.Db.Create(&entity.Users{})

	var result string
	fmt.Print(db.Raw("SELECT * FROM users").Scan(&result))
}

// Stop MySQL
func (t UserControllerTest) After() {
	app.Db.Close()
	testMySQLd.Stop()
}

//func (t UserControllerTest) TestPostUserOk() {
//	users := entity.Users{
//		Username: "test",
//		Email: "test@gmail.com",
//		Password: "testtest",
//	}
//
//	var response = userController.PostUser(users)
//
//	success := map[string]string {
//		"message": "user creation succeeded.",
//	}
//
//	t.AssertEqual(success, response)
//}