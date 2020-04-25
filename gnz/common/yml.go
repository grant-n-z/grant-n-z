package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"io/ioutil"

	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
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
	TimeMillisStr string `yaml:"time-millis"`
	TimeMillis    int
}

// About server data in grant_n_z_server.yaml
type ServerConfig struct {
	Port                   string `yaml:"port"`
	SignedInPrivateKeyPath string `yaml:"signed-in-token-private-key-path"`
	ValidatePublicKeyPath  string `yaml:"validate-token-public-key-path"`
	TokenExpireHourStr     string `yaml:"token-expire-hour"`
	SignAlgorithm          string `yaml:"sign-algorithm"`
	SignedInPrivateKey     *rsa.PrivateKey
	ValidatePublicKey      *rsa.PublicKey
	SigningMethod          jwt.SigningMethod
	TokenExpireHour        int
}

// About db data in grant_n_z_{component}.yaml
type DbConfig struct {
	Engine   string `yaml:"engine"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
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
	timMillisStr := yml.Cacher.TimeMillisStr

	if strings.Contains(timMillisStr, "$") {
		timMillisStr = os.Getenv(yml.Cacher.TimeMillisStr[1:])
	}

	yml.Cacher.TimeMillisStr = timMillisStr
	yml.Cacher.TimeMillis, _ = strconv.Atoi(timMillisStr)
	return yml.Cacher
}

// Getter ServerConfig
func (yml YmlConfig) GetServerConfig() ServerConfig {
	port := yml.Server.Port
	privateKeyStr := yml.Server.SignedInPrivateKeyPath
	publicKeyStr := yml.Server.SignedInPrivateKeyPath
	tokenExpireHourStr := yml.Server.TokenExpireHourStr
	signAlgorithm := yml.Server.SignAlgorithm

	if strings.Contains(port, "$") {
		port = os.Getenv(yml.Server.Port[1:])
	}

	if strings.Contains(privateKeyStr, "$") {
		privateKeyStr = os.Getenv(yml.Server.SignedInPrivateKeyPath[1:])
	}

	if strings.Contains(publicKeyStr, "$") {
		publicKeyStr = os.Getenv(yml.Server.ValidatePublicKeyPath[1:])
	}

	if strings.Contains(tokenExpireHourStr, "$") {
		tokenExpireHourStr = os.Getenv(yml.Server.TokenExpireHourStr[1:])
	}

	if strings.Contains(signAlgorithm, "$") {
		signAlgorithm = os.Getenv(yml.Server.SignAlgorithm[1:])
	}

	yml.Server.Port = port
	yml.Server.SignedInPrivateKeyPath = privateKeyStr
	yml.Server.ValidatePublicKeyPath = publicKeyStr
	yml.Server.TokenExpireHourStr = tokenExpireHourStr

	// Generate server config data
	yml.Server.TokenExpireHour, _ = strconv.Atoi(tokenExpireHourStr)

	yml.Server.SignAlgorithm = signAlgorithm
	switch yml.Server.SignAlgorithm {
	case "rsa256":
		yml.Server.SigningMethod = jwt.SigningMethodRS256
	case "rsa512":
		yml.Server.SigningMethod = jwt.SigningMethodRS512
	case "ecdsa256":
		yml.Server.SigningMethod = jwt.SigningMethodES256
	case "ecdsa512":
		yml.Server.SigningMethod = jwt.SigningMethodES512
	default:
		panic("Invalid sign-algorithm data. Only support `rsa256` or `rsa512` or `ecdsa256` or `ecdsa512`")
	}

	privateKeyBytes, err := ioutil.ReadFile(yml.Server.SignedInPrivateKeyPath)
	fmt.Println(string(privateKeyBytes))
	if err != nil {
		panic("Not found private key file. " + err.Error())
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic("Invalid private key data. " + err.Error())
	}
	yml.Server.SignedInPrivateKey = signKey

	publicKeyBytes, err := ioutil.ReadFile(yml.Server.ValidatePublicKeyPath)
	if err != nil {
		panic("Not found public key file. " + err.Error())
	}

	validateKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		panic("Invalid public key data. " + err.Error())
	}
	yml.Server.ValidatePublicKey = validateKey

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
	name := yml.Db.Name

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

	if strings.Contains(name, "$") {
		name = os.Getenv(yml.Db.Name[1:])
	}

	yml.Db.Engine = engine
	yml.Db.User = user
	yml.Db.Password = password
	yml.Db.Host = host
	yml.Db.Port = port
	yml.Db.Name = name
	return yml.Db
}
