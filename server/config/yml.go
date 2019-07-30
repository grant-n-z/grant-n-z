package config

import (
	"fmt"
	"os"
	"strings"
)

type YmlConfig struct {
	App AppConfig `yaml:"app"`
	Db  DbConfig  `yaml:"db"`
}

type AppConfig struct {
	Version          string `yaml:"version"`
	PrivateKeyBase64 string `yaml:"private-key-base64"`
	Environment      string `yaml:"environment"`
	LogLevel         string `yaml:"log-level"`
	PolicyFilePath   string `yaml:"policy-file-path"`
}

type DbConfig struct {
	Engine   string `yaml:"engine"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

func (yml YmlConfig) GetAppConfig() AppConfig {
	privateKeyBase64 := yml.App.PrivateKeyBase64
	environment := yml.App.Environment
	logLevel := yml.App.LogLevel
	policyFilePath := yml.App.PolicyFilePath

	if strings.Contains(privateKeyBase64, "$") {
		privateKeyBase64 = os.Getenv(yml.App.PrivateKeyBase64[1:])
	}

	if strings.Contains(environment, "$") {
		environment = os.Getenv(yml.App.Environment[1:])
	}

	if strings.Contains(logLevel, "$") {
		logLevel = os.Getenv(yml.App.LogLevel[1:])
	}

	if strings.Contains(policyFilePath, "$") {
		policyFilePath = os.Getenv(yml.App.PolicyFilePath[1:])
	}

	yml.App.PrivateKeyBase64 = privateKeyBase64
	yml.App.Environment = environment
	yml.App.LogLevel = logLevel
	return yml.App
}

func (yml YmlConfig) GetDataSourceUrl() string {
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

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", user, password, host, port, db)
}
