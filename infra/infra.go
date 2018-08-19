package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/tomoyane/grant-n-z/domain"
	"fmt"
	"net/http"
)

var (
	Db *gorm.DB
	dbSource string
	yml domain.YmlModel
)

func InitDB() {
	switch os.Getenv("ENV") {
	case "test":
		yml = readYml("../app-test.yaml")
		dbSource = yml.GetDataSourceUrl()
	default:
		yml = readYml("app.yaml")
		dbSource = yml.GetDataSourceUrl()
	}

	fmt.Print(dbSource)

	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	db.DB()
	Db = db
}

func GetHostName() string {
	host, err := os.Hostname()
	if err != nil {
		domain.ErrorResponse{}.Print(http.StatusInternalServerError, "failed hostname", "")
	}
	return host
}

func readYml(ymlName string) domain.YmlModel {
	yml, err := ioutil.ReadFile(ymlName)
	if err != nil {
		domain.ErrorResponse{}.Print(http.StatusInternalServerError, "failed read yml", "")
	}

	var db domain.YmlModel
	err = yaml.Unmarshal(yml, &db)

	return db
}
