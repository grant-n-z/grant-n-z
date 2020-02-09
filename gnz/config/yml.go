package config

import (
	"os"
	"strings"
)

// grant_n_z_cache.yaml
type YmlConfig struct {
	App    AppConfig    `yaml:"app"`
	Server ServerConfig `yaml:"server"`
	Db     DbConfig     `yaml:"db"`
	Redis  RedisConfig  `yaml:"redis"`
}

// About app data in grant_n_z_cache.yaml
type AppConfig struct {
	Version  string `yaml:"version"`
	LogLevel string `yaml:"log-level"`
}

// About server data in grant_n_z_cache.yaml
type ServerConfig struct {
	SignedInPrivateKeyBase64 string `yaml:"signed-in-token-private-key-base64"`
}

// About db data in grant_n_z_cache.yaml
type DbConfig struct {
	Engine   string `yaml:"engine"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

// About redis data in grant_n_z_cache.yaml
type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
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

// Getter ServerConfig
func (yml YmlConfig) GetServerConfig() ServerConfig {
	privateKeyBase64 := yml.Server.SignedInPrivateKeyBase64
	if strings.Contains(privateKeyBase64, "$") {
		privateKeyBase64 = os.Getenv(yml.Server.SignedInPrivateKeyBase64[1:])
	}

	return yml.Server
}

// Getter RedisConfig
func (yml YmlConfig) GetRedisConfig() RedisConfig {
	host := yml.Redis.Host
	password := yml.Redis.Password
	port := yml.Redis.Port
	db := yml.Redis.Db

	if strings.Contains(host, "$") {
		host = os.Getenv(yml.Redis.Host[1:])
	}

	if strings.Contains(password, "$") {
		password = os.Getenv(yml.Redis.Password[1:])
	}

	if strings.Contains(port, "$") {
		port = os.Getenv(yml.Redis.Port[1:])
	}

	if strings.Contains(db, "$") {
		db = os.Getenv(yml.Redis.Db[1:])
	}

	yml.Redis.Host = host
	yml.Redis.Password = password
	yml.Redis.Port = port
	yml.Redis.Db = db
	return yml.Redis
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
