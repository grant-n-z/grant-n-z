package config

import (
	"os"

	"gopkg.in/yaml.v2"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	test   = "test"
	appEnv = "APP_ENV"
)

var (
	Db       *gorm.DB
	App      AppConfig
	yml      YmlConfig
	dbSource string
)

func InitConfig() {
	initYaml()
	initDb()
}

func initDb() {
	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	db.DB()
	db.LogMode(false)
	Db = db
}

func initYaml() {
	switch os.Getenv(appEnv) {
	case test:
		yml = readYml("../../app-test.yaml")
		dbSource = yml.GetDataSourceUrl()
		App = yml.GetAppConfig()
	default:
		yml = readYml("app.yaml")
		dbSource = yml.GetDataSourceUrl()
		App = yml.GetAppConfig()
	}
}

func readYml(ymlName string) YmlConfig {
	var yml YmlConfig
	data, err := ioutil.ReadFile(ymlName)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &yml); err != nil {
		panic(err)
	}

	return yml
}
