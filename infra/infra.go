package infra

import (
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/domain"
	"os"
	"fmt"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/handler"
	"net/http"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"golang.org/x/crypto/bcrypt"
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
	if !Db.HasTable(entity.User{}.TableName()) {
		Db.CreateTable(&entity.User{})
		hash, _ := bcrypt.GenerateFromPassword([] byte("admin"), bcrypt.DefaultCost)
		user := entity.User{
			Username: "admin",
			Email: "admin@gmail.com",
			Password: string(hash),
		}
		Db.Create(&user)
	}

	if !Db.HasTable(entity.Principal{}.TableName()) {
		Db.CreateTable(&entity.Principal{})
	}

	if !Db.HasTable(entity.Role{}.TableName()) {
		Db.CreateTable(&entity.Role{})
	}

	if !Db.HasTable(entity.Token{}.TableName()) {
		Db.CreateTable(&entity.Token{})
	}

	if !Db.HasTable(entity.Member{}.TableName()) {
		Db.CreateTable(&entity.Member{})
	}

	if !Db.HasTable(entity.MemberRole{}.TableName()) {
		Db.CreateTable(&entity.MemberRole{})
	}
}

func GetHostName() string {
	host, err := os.Hostname()
	if err != nil {
		handler.ErrorResponse{}.Print(http.StatusInternalServerError, "failed hostname", "")
	}
	return host
}

func readYml(ymlName string) domain.Yml {
	yml, err := ioutil.ReadFile(ymlName)
	if err != nil {
		handler.ErrorResponse{}.Print(http.StatusInternalServerError, "failed read yml", "")
	}

	var db domain.Yml
	err = yaml.Unmarshal(yml, &db)

	return db
}
