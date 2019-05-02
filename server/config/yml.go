package config

import (
	"fmt"
	"os"
	"strings"
)

type YmlConfig struct {
	App     AppConfig `yaml:"app"`
	DbModel DbConfig  `yaml:"db"`
}

type AppConfig struct {
	Version     string `yaml:"version"`
	PrivateKey  string `yaml:"private-key"`
	PublicKey   string `yaml:"public-key"`
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log-level"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

func (yml YmlConfig) GetAppVersion() string {
	return yml.App.Version
}

func (yml YmlConfig) GetAppEnvironment() string {
	return os.Getenv(yml.App.Environment[1:])
}

func (yml YmlConfig) GetAppLogLevel() string {
	return os.Getenv(yml.App.LogLevel[1:])
}

func (yml YmlConfig) GetDataSourceUrl() string {
	user := yml.DbModel.User
	password := yml.DbModel.Password
	host := yml.DbModel.Host
	port := yml.DbModel.Port
	db := yml.DbModel.Db

	if strings.Contains(user, "$") {
		user = os.Getenv(yml.DbModel.User[1:])
	}

	if strings.Contains(password, "$") {
		password = os.Getenv(yml.DbModel.Password[1:])
	}

	if strings.Contains(host, "$") {
		host = os.Getenv(yml.DbModel.Host[1:])
	}

	if strings.Contains(port, "$") {
		port = os.Getenv(yml.DbModel.Port[1:])
	}

	if strings.Contains(db, "$") {
		db = os.Getenv(yml.DbModel.Db[1:])
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", user, password, host, port, db)
}
