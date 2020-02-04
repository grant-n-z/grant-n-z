package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
)

var (
	App     AppConfig
	Redis   RedisConfig
	Db      DbConfig
	GServer ServerConfig
)

// Initialize GrantNZ server config
// The config is app.yaml data structure
func InitGrantNZServerConfig() {
	yml := readLocalYml("app.yaml")
	App = yml.GetAppConfig()
	GServer = yml.GetServerConfig()
	Db = yml.GetDbConfig()
	Redis = yml.GetRedisConfig()
}

// Initialize GrantNZ cache scheduler config
// The config is app.yaml data structure
func InitGrantNZCacheSchedulerConfig() {
	yml := readLocalYml("app.yaml")
	Db = yml.GetDbConfig()
	Redis = yml.GetRedisConfig()
}

// Read yaml file
func readLocalYml(ymlName string) YmlConfig {
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
