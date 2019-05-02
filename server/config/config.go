package config

import (
	"os"

	"gopkg.in/yaml.v2"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	test = "test"
)

var (
	Db          *gorm.DB
	Environment string
	LogLevel    string
	dbSource    string
	ymlConfig   YmlConfig
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
	Db = db
}

func initYaml() {
	switch os.Getenv("ENV") {
	case test:
		ymlConfig = readYml("app-test.yaml")
		dbSource = ymlConfig.GetDataSourceUrl()
		Environment = ymlConfig.GetAppEnvironment()
		LogLevel = ymlConfig.GetAppLogLevel()
	default:
		ymlConfig = readYml("app.yaml")
		dbSource = ymlConfig.GetDataSourceUrl()
		Environment = ymlConfig.GetAppEnvironment()
		LogLevel = ymlConfig.GetAppLogLevel()
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
