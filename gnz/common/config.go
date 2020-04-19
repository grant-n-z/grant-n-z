package common

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
)

var (
	App     AppConfig
	Etcd    EtcdConfig
	Db      DbConfig
	GServer ServerConfig
	GCacher CacherConfig
)

// Initialize GrantNZ server config
// The config is grant_n_z_server.yaml data structure
func InitGrantNZServerConfig(yamlPath string) {
	yml := readLocalYml(yamlPath)
	App = yml.GetAppConfig()
	Db = yml.GetDbConfig()
	Etcd = yml.GetEtcdConfig()
	GServer = yml.GetServerConfig()
}

// Initialize GrantNZ server config
// The config is grant_n_z_cacher.yaml data structure
func InitGrantNZCacherConfig(yamlPath string) {
	yml := readLocalYml(yamlPath)
	App = yml.GetAppConfig()
	Db = yml.GetDbConfig()
	Etcd = yml.GetEtcdConfig()
	GCacher = yml.GetCacherConfig()
}

// Read yaml file
func readLocalYml(ymlPath string) YmlConfig {
	var yml YmlConfig
	data, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &yml); err != nil {
		panic(err)
	}

	return yml
}
