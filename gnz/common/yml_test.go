package common

import (
	"os"
	"strings"
	"testing"
)

// GetAppConfig test
func TestGetAppConfig(t *testing.T) {
	appConfig := AppConfig{LogLevel: "$LOG_LEVEL"}
	ymlConfig := YmlConfig{App: appConfig}

	// Test data
	os.Setenv("LOG_LEVEL", "test")

	if !strings.EqualFold(ymlConfig.GetAppConfig().LogLevel, "test") {
		t.Errorf("Incorrect GetAppConfig test. log_level = %s", ymlConfig.GetAppConfig().LogLevel)
		t.FailNow()
	}
}

// GetCacherConfig test
func TestGetCacherConfig(t *testing.T) {
	cacherConfig := CacherConfig{TimeMillisStr: "$CACHER_TIME_MILLIS"}
	ymlConfig := YmlConfig{Cacher: cacherConfig}

	// Test data
	os.Setenv("CACHER_TIME_MILLIS", "100")

	if !strings.EqualFold(ymlConfig.GetCacherConfig().TimeMillisStr, "100") {
		t.Errorf("Incorrect CacherConfig test. time-millis = %s", ymlConfig.GetCacherConfig().TimeMillisStr)
		t.FailNow()
	}
}

// GetServerConfig test
func TestGetServerConfig(t *testing.T) {
	serverConfig := ServerConfig{
		Port:                   "$SERVER_PORT",
		SignedInPrivateKeyPath: "$SERVER_PRIVATE_KEY_PATH",
		ValidatePublicKeyPath:  "$SERVER_PUBLIC_KEY_PATH",
		TokenExpireHourStr:     "$SERVER_TOKEN_EXPIRE_HOUR",
		SignAlgorithm:          "$SERVER_SIGN_ALGORITHM",
	}
	ymlConfig := YmlConfig{Server: serverConfig}

	// Test data
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("SERVER_PRIVATE_KEY_PATH", "./test-private.key")
	os.Setenv("SERVER_PUBLIC_KEY_PATH", "./test-public.key")
	os.Setenv("SERVER_TOKEN_EXPIRE_HOUR", "100")
	os.Setenv("SERVER_SIGN_ALGORITHM", "rsa256")

	if !strings.EqualFold(ymlConfig.GetServerConfig().Port, "8080") {
		t.Errorf("Incorrect GetServerConfig test. port = %s", ymlConfig.GetServerConfig().Port)
		t.FailNow()
	}

	if strings.EqualFold(ymlConfig.GetServerConfig().SignedInPrivateKeyPath, "") {
		t.Errorf("Incorrect GetServerConfig test. privaet_key_base64 = %s", ymlConfig.GetServerConfig().SignedInPrivateKeyPath)
		t.FailNow()
	}

	if strings.EqualFold(ymlConfig.GetServerConfig().ValidatePublicKeyPath, "") {
		t.Errorf("Incorrect GetServerConfig test. public_key_base64 = %s", ymlConfig.GetServerConfig().SignedInPrivateKeyPath)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetServerConfig().TokenExpireHourStr, "100") {
		t.Errorf("Incorrect GetServerConfig test. token-expire-hour = %s", ymlConfig.GetServerConfig().ValidatePublicKeyPath)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetServerConfig().SignAlgorithm, "rsa256") {
		t.Errorf("Incorrect GetServerConfig test. rsa-algorithm = %s", ymlConfig.GetServerConfig().SignAlgorithm)
		t.FailNow()
	}
}

// GetEtcdConfig test
func TestGetEtcdConfig(t *testing.T) {
	etcdConfig := EtcdConfig{Host: "$ETCD_HOST", Port: "$ETCD_PORT"}
	ymlConfig := YmlConfig{Etcd: etcdConfig}

	// Test data
	os.Setenv("ETCD_HOST", "localhost")
	os.Setenv("ETCD_PORT", "2380")

	if !strings.EqualFold(ymlConfig.GetEtcdConfig().Host, "localhost") {
		t.Errorf("Incorrect GetEtcdConfig test. host = %s", ymlConfig.GetEtcdConfig().Host)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetEtcdConfig().Port, "2380") {
		t.Errorf("Incorrect GetEtcdConfig test. port = %s", ymlConfig.GetEtcdConfig().Port)
		t.FailNow()
	}
}

// GetDbConfig test
func TestGetDbConfig(t *testing.T) {
	dbConfig := DbConfig{
		Engine:   "$DB_ENGINE",
		Host:     "$DB_HOST",
		User:     "$DB_USER",
		Password: "$DB_PASSWORD",
		Port:     "$DB_PORT",
		Name:     "$DB_NAME",
	}
	ymlConfig := YmlConfig{Db: dbConfig}

	// Test data
	os.Setenv("DB_ENGINE", "mysql")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "grant_n_z")

	if !strings.EqualFold(ymlConfig.GetDbConfig().Engine, "mysql") {
		t.Errorf("Incorrect GetEtcdConfig test. engine = %s", ymlConfig.GetDbConfig().Engine)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetDbConfig().Host, "localhost") {
		t.Errorf("Incorrect GetEtcdConfig test. host = %s", ymlConfig.GetDbConfig().Host)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetDbConfig().User, "root") {
		t.Errorf("Incorrect GetEtcdConfig test. user = %s", ymlConfig.GetDbConfig().User)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetDbConfig().Password, "root") {
		t.Errorf("Incorrect GetEtcdConfig test. password = %s", ymlConfig.GetDbConfig().Password)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetDbConfig().Port, "3306") {
		t.Errorf("Incorrect GetEtcdConfig test. port = %s", ymlConfig.GetDbConfig().Port)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetDbConfig().Name, "grant_n_z") {
		t.Errorf("Incorrect GetEtcdConfig test. db = %s", ymlConfig.GetDbConfig().Name)
		t.FailNow()
	}
}
