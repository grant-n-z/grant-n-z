package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/tomoyane/grant-n-z/domain"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"fmt"
)

var (
	Db *gorm.DB
	dbSource string
	Yaml domain.Yml
)

func InitYaml()  {
	switch os.Getenv("ENV") {
	case "test":
		Yaml = readYml("../app-test.yaml")
		dbSource = Yaml.GetDataSourceUrl()
	default:
		Yaml = readYml("app.yaml")
		dbSource = Yaml.GetDataSourceUrl()
	}

	fmt.Print(dbSource)
}

func InitDB() {
	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	db.DB()
	Db = db
}

func MigrateDB() {
	if (!Db.HasTable(&entity.User{})) {
		Db.CreateTable(&entity.User{})
		Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.User{})
	}

	if (!Db.HasTable(&entity.Token{})) {
		Db.CreateTable(&entity.Token{})
		Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Token{})
	}

	if (!Db.HasTable(&entity.Role{})) {
		Db.CreateTable(&entity.Role{})
		Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Role{})
	}
}

func GetHostName() string {
	host, err := os.Hostname()
	if err != nil {
		domain.ErrorResponse{}.Print(http.StatusInternalServerError, "failed hostname", "")
	}
	return host
}

func readYml(ymlName string) domain.Yml {
	yml, err := ioutil.ReadFile(ymlName)
	if err != nil {
		domain.ErrorResponse{}.Print(http.StatusInternalServerError, "failed read yml", "")
	}

	var db domain.Yml
	err = yaml.Unmarshal(yml, &db)

	return db
}
