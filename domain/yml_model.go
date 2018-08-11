package domain

type YmlModel struct {
	DbModel DbModel `yaml:"db"`
}

type DbModel struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

func (yml YmlModel) GetDataSourceUrl() string {
	return yml.DbModel.User + ":" + yml.DbModel.Password +
		"@tcp(" + yml.DbModel.Host + ":" + yml.DbModel.Port + ")/" + yml.DbModel.Db +
		"?charset=utf8&parseTime=True"
}
