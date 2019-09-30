package config

import (
	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
)

var (
	App   AppConfig
	Redis RedisConfig
	Db    DbConfig
)

func InitConfig() {
	initYaml()
}

func initYaml() {
	var yml YmlConfig
	switch os.Getenv(property.AppEnv) {
	case property.Test:
		yml = readYml("../../app-test.yaml")
	default:
		yml = readYml("app.yaml")
	}

	App = yml.GetAppConfig()
	Db = yml.GetDbConfig()
	Redis = yml.GetRedisConfig()
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
