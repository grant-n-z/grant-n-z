package cache

import (
	"os"
	"testing"

	"github.com/tomoyane/grant-n-z/gnz/config"
)

// Connection not use etcd
func TestConnection_NotUseEtcd(t *testing.T) {
	connection = nil
	os.Setenv("GRANT_N_Z_ETCD_HOST", "")
	os.Setenv("GRANT_N_Z_ETCD_PORT", "")

	config.InitGrantNZCacherConfig("../../gnzcacher/grant_n_z_cacher.yaml")
	InitEtcd()
}

// Connection etcd
func TestConnection(t *testing.T) {
	connection = nil
	os.Setenv("GRANT_N_Z_ETCD_HOST", "localhost")
	os.Setenv("GRANT_N_Z_ETCD_PORT", "2222")

	config.InitGrantNZCacherConfig("../../gnzcacher/grant_n_z_cacher.yaml")
	InitEtcd()
}

// Connection etcd
func TestClose_ConnectionIsNil(t *testing.T) {
	connection = nil
	os.Setenv("GRANT_N_Z_ETCD_HOST", "")
	os.Setenv("GRANT_N_Z_ETCD_PORT", "")

	config.InitGrantNZCacherConfig("../../gnzcacher/grant_n_z_cacher.yaml")
	InitEtcd()
	Close()
}

// Connection etcd
func TestClose_ConnectionIsNotNil(t *testing.T) {
	connection = nil
	os.Setenv("GRANT_N_Z_ETCD_HOST", "localhost")
	os.Setenv("GRANT_N_Z_ETCD_PORT", "2222")

	config.InitGrantNZCacherConfig("../../gnzcacher/grant_n_z_cacher.yaml")
	InitEtcd()
	Close()
}

