package domain

import "fmt"

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
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		yml.DbModel.User,
		yml.DbModel.Password,
		yml.DbModel.Host,
		yml.DbModel.Port,
		yml.DbModel.Db)
}
