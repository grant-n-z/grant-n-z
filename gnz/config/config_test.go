package config

import (
	"strings"
	"testing"
)

// InitGrantNZServerConfig test
func TestInitGrantNZServerConfig(t *testing.T) {
	InitGrantNZServerConfig("../../gnzcacher/grant_n_z_cacher.yaml")
	if strings.EqualFold(App.Version, "") {
		t.Errorf("Incorrect TestInitGrantNZServerConfig test")
	}
}

// InitGrantNZCacherConfig test
func TestInitGrantNZCacherConfig(t *testing.T) {
	InitGrantNZCacherConfig("../../gnzserver/grant_n_z_server.yaml")
	if strings.EqualFold(App.Version, "") {
		t.Errorf("Incorrect TestInitGrantNZCacherConfig test")
	}
}

// readLocalYml test
func TestReadYaml(t *testing.T) {
	// grant_n_z_server
	yml := readLocalYml("../../gnzserver/grant_n_z_server.yaml")
	if strings.EqualFold(yml.Db.Db, "") {
		t.Errorf("Incorrect readLocalYml for grant_n_z_server test")
	}

	// grant_n_z_cacher
	yml = readLocalYml("../../gnzcacher/grant_n_z_cacher.yaml")
	if strings.EqualFold(yml.Db.Db, "") {
		t.Errorf("Incorrect readLocalYml for grant_n_z_cacher test")
	}
}
