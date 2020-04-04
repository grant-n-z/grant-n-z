package config

import (
	"os"
	"strings"
	"testing"
)

// GetAppConfig test
func TestGetAppConfig(t *testing.T) {
	appConfig := AppConfig{LogLevel: "$GRANT_N_Z_LOG_LEVEL"}
	ymlConfig := YmlConfig{App: appConfig}

	// Test data
	os.Setenv("GRANT_N_Z_LOG_LEVEL", "test")

	if !strings.EqualFold(ymlConfig.GetAppConfig().LogLevel, "test") {
		t.Errorf("Incorrect GetAppConfig test. log_level = %s", ymlConfig.GetAppConfig().LogLevel)
		t.FailNow()
	}
}

// GetServerConfig test
func TestGetServerConfig(t *testing.T) {
	serverConfig := ServerConfig{SignedInPrivateKeyBase64: "$GRANT_N_Z_PRIVATE_KEY"}
	ymlConfig := YmlConfig{Server: serverConfig}

	// Test data
	os.Setenv("GRANT_N_Z_PRIVATE_KEY", "dGVzdF9rZXkK")

	if !strings.EqualFold(ymlConfig.GetServerConfig().SignedInPrivateKeyBase64, "dGVzdF9rZXkK") {
		t.Errorf("Incorrect GetServerConfig test. privaet_key_base64 = %s", ymlConfig.GetServerConfig().SignedInPrivateKeyBase64)
		t.FailNow()
	}
}

// GetEtcdConfig test
func TestGetEtcdConfig(t *testing.T) {
	etcdConfig := EtcdConfig{Host: "$GRANT_N_Z_ETCD_HOST", Port: "$GRANT_N_Z_ETCD_PORT"}
	ymlConfig := YmlConfig{Etcd: etcdConfig}

	// Test data
	os.Setenv("GRANT_N_Z_ETCD_HOST", "localhost")
	os.Setenv("GRANT_N_Z_ETCD_PORT", "2380")

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
		Engine: "$GRANT_N_Z_ENGINE",
		Host: "$GRANT_N_Z_HOST",
		User: "$GRANT_N_Z_USER",
		Password: "$GRANT_N_Z_PASSWORD",
		Port: "$GRANT_N_Z_PORT",
		Db: "$GRANT_N_Z_DB",
	}
	ymlConfig := YmlConfig{Db: dbConfig}

	// Test data
	os.Setenv("GRANT_N_Z_ENGINE", "mysql")
	os.Setenv("GRANT_N_Z_HOST", "localhost")
	os.Setenv("GRANT_N_Z_USER", "root")
	os.Setenv("GRANT_N_Z_PASSWORD", "root")
	os.Setenv("GRANT_N_Z_PORT", "3306")
	os.Setenv("GRANT_N_Z_DB", "grant_n_z")

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

	if !strings.EqualFold(ymlConfig.GetDbConfig().Db, "grant_n_z") {
		t.Errorf("Incorrect GetEtcdConfig test. db = %s", ymlConfig.GetDbConfig().Db)
		t.FailNow()
	}
}
