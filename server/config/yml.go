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
	Version    string `yaml:"version"`
	PrivateKey string `yaml:"private-key"`
	PublicKey  string `yaml:"public-key"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

func (yml YmlConfig) GetDataSourceUrl() string {
	user := yml.DbModel.User
	password := yml.DbModel.Password
	host := yml.DbModel.Host
	port := yml.DbModel.Port
	db := yml.DbModel.Db

	fmt.Printf("host = %s\n", host)
	fmt.Printf("port = %s\n", port)
	fmt.Printf("db = %s\n", db)
	fmt.Printf("user = %s\n", user)

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
