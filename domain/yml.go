package domain

import (
	"fmt"
	"os"
	"strings"
)

type Yml struct {
	DbModel DbModel `yaml:"db"`
}

type DbModel struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

func (yml Yml) GetDataSourceUrl() string {
	user := yml.DbModel.User
	pass := yml.DbModel.Password
	host := yml.DbModel.Host
	port := yml.DbModel.Port
	db := yml.DbModel.Db

	if strings.Contains(user, "$") {
		user = os.Getenv(yml.DbModel.User[1:])
	}

	if strings.Contains(pass, "$") {
		pass = os.Getenv(yml.DbModel.Password[1:])
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

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", user, pass, host, port, db)
}
