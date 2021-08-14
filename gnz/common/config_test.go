package common

import (
	"os"
	"strings"
	"testing"
)

// InitGrantNZServerConfig test
func TestInitGrantNZServerConfig(t *testing.T) {
	os.Setenv("SERVER_PRIVATE_KEY_PATH", "./test-private.key")
	os.Setenv("SERVER_PUBLIC_KEY_PATH", "./test-public.key")
	os.Setenv("SERVER_SIGN_ALGORITHM", "rsa256")
	InitGrantNZServerConfig("../../gnzserver/grant_n_z_server.yaml")
	if strings.EqualFold(App.LogLevel, "") {
		t.Errorf("Incorrect TestInitGrantNZServerConfig test")
		t.FailNow()
	}
}

// InitGrantNZCacherConfig test
func TestInitGrantNZCacherConfig(t *testing.T) {
	InitGrantNZCacherConfig("../../gnzcacher/grant_n_z_cacher.yaml")
	if strings.EqualFold(App.LogLevel, "") {
		t.Errorf("Incorrect TestInitGrantNZCacherConfig test")
		t.FailNow()
	}
}

// readLocalYml test
func TestReadYaml(t *testing.T) {
	// grant_n_z_server
	yml := readLocalYml("../../gnzserver/grant_n_z_server.yaml")
	if strings.EqualFold(yml.Db.Name, "") {
		t.Errorf("Incorrect readLocalYml for grant_n_z_server test")
		t.FailNow()
	}

	// grant_n_z_cacher
	yml = readLocalYml("../../gnzcacher/grant_n_z_cacher.yaml")
	if strings.EqualFold(yml.Db.Name, "") {
		t.Errorf("Incorrect readLocalYml for grant_n_z_cacher test")
		t.FailNow()
	}
}
