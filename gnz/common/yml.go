package common

import (
	"os"
	"strings"
)

// grant_n_z_{component}.yaml
type YmlConfig struct {
	App    AppConfig    `yaml:"app"`
	Cacher CacherConfig `yaml:"cacher"`
	Server ServerConfig `yaml:"server"`
	Db     DbConfig     `yaml:"db"`
	Etcd   EtcdConfig   `yaml:"etcd"`
}

// About app data in grant_n_z_{component}.yaml
type AppConfig struct {
	Version  string `yaml:"version"`
	LogLevel string `yaml:"log-level"`
}

// About app data in grant_n_z_cacher.yaml
type CacherConfig struct {
	TimeMillis string `yaml:"time-millis"`
}

// About server data in grant_n_z_server.yaml
type ServerConfig struct {
	SignedInPrivateKeyBase64 string `yaml:"signed-in-token-private-key-base64"`
}

// About db data in grant_n_z_{component}.yaml
type DbConfig struct {
	Engine   string `yaml:"engine"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

// About etcd data in grant_n_z_{component}.yaml
type EtcdConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// Getter AppConfig
func (yml YmlConfig) GetAppConfig() AppConfig {
	logLevel := yml.App.LogLevel

	if strings.Contains(logLevel, "$") {
		logLevel = os.Getenv(yml.App.LogLevel[1:])
	}

	yml.App.LogLevel = logLevel
	return yml.App
}

// Getter CacherConfig
func (yml YmlConfig) GetCacherConfig() CacherConfig {
	timMillis := yml.Cacher.TimeMillis

	if strings.Contains(timMillis, "$") {
		timMillis = os.Getenv(yml.Cacher.TimeMillis[1:])
	}

	yml.Cacher.TimeMillis = timMillis
	return yml.Cacher
}

// Getter ServerConfig
func (yml YmlConfig) GetServerConfig() ServerConfig {
	privateKeyBase64 := yml.Server.SignedInPrivateKeyBase64
	if strings.Contains(privateKeyBase64, "$") {
		privateKeyBase64 = os.Getenv(yml.Server.SignedInPrivateKeyBase64[1:])
	}

	yml.Server.SignedInPrivateKeyBase64 = privateKeyBase64
	return yml.Server
}

// Getter EtcdConfig
func (yml YmlConfig) GetEtcdConfig() EtcdConfig {
	if &yml.Etcd == nil {
		return EtcdConfig{}
	}
	host := yml.Etcd.Host
	port := yml.Etcd.Port

	if strings.Contains(host, "$") {
		host = os.Getenv(yml.Etcd.Host[1:])
	}

	if strings.Contains(port, "$") {
		port = os.Getenv(yml.Etcd.Port[1:])
	}

	yml.Etcd.Host = host
	yml.Etcd.Port = port
	return yml.Etcd
}

// Getter DbConfig
func (yml YmlConfig) GetDbConfig() DbConfig {
	engine := yml.Db.Engine
	user := yml.Db.User
	password := yml.Db.Password
	host := yml.Db.Host
	port := yml.Db.Port
	db := yml.Db.Db

	if strings.Contains(engine, "$") {
		engine = os.Getenv(yml.Db.Engine[1:])
	}

	if strings.Contains(user, "$") {
		user = os.Getenv(yml.Db.User[1:])
	}

	if strings.Contains(password, "$") {
		password = os.Getenv(yml.Db.Password[1:])
	}

	if strings.Contains(host, "$") {
		host = os.Getenv(yml.Db.Host[1:])
	}

	if strings.Contains(port, "$") {
		port = os.Getenv(yml.Db.Port[1:])
	}

	if strings.Contains(db, "$") {
		db = os.Getenv(yml.Db.Db[1:])
	}

	yml.Db.Engine = engine
	yml.Db.User = user
	yml.Db.Password = password
	yml.Db.Host = host
	yml.Db.Port = port
	yml.Db.Db = db
	return yml.Db
}
