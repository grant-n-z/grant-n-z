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
	"github.com/satori/go.uuid"
)

var (
	Db *gorm.DB
	dbSource string
	Yaml domain.Yml
	userUuid uuid.UUID
	roleUuid uuid.UUID
	serviceUuid uuid.UUID
	memberUuid uuid.UUID
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
		userUuid, _ = uuid.NewV4()
		hash, _ := bcrypt.GenerateFromPassword([] byte("admin"), bcrypt.DefaultCost)
		user := entity.User{
			Username: "admin",
			Email: "admin@gmail.com",
			Password: string(hash),
			Uuid: userUuid,
		}
		Db.Create(&user)
	}

	if !Db.HasTable(entity.Role{}.TableName()) {
		Db.CreateTable(&entity.Role{})
		roleUuid, _ = uuid.NewV4()
		role := entity.Role{
			Uuid: roleUuid,
			Permission: "admin",
		}
		Db.Create(&role)
	}

	if !Db.HasTable(entity.Service{}.TableName()) {
		Db.CreateTable(&entity.Service{})
		serviceUuid, _ = uuid.NewV4()
		service := entity.Service{
			Uuid: serviceUuid,
			Name: "admin-service",
		}
		Db.Create(&service)
	}

	if !Db.HasTable(entity.Member{}.TableName()) {
		Db.CreateTable(&entity.Member{})
		memberUuid, _ = uuid.NewV4()
		member := entity.Member{
			Uuid: memberUuid,
			ServiceUuid: serviceUuid,
			UserUuid: userUuid,
		}
		Db.Create(&member)
	}

	if !Db.HasTable(entity.Principal{}.TableName()) {
		Db.CreateTable(&entity.Principal{})
		principal := entity.Principal{
			MemberUuid: memberUuid,
			RoleUuid: roleUuid,
		}
		Db.Create(&principal)
	}

	if !Db.HasTable(entity.Token{}.TableName()) {
		Db.CreateTable(&entity.Token{})
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
