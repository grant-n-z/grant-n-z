package service

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var (
	stubConnection   *gorm.DB
	extractorService ExtractorService
)

func init() {
	log.InitLogger("info")

	stubConnection, _ = gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
}

// Test constructor
func TestNewExtractorService(t *testing.T) {
	NewExtractorService()
}

// Test get policies
func TestGetPolicies(t *testing.T) {
	stubPolicyRepository := driver.RdbmsPolicyRepository{Connection: stubConnection}
	stubUserRepository := driver.RdbmsUserRepository{Connection: stubConnection}
	extractorService = ExtractorServiceImpl{
		PolicyRepository: stubPolicyRepository,
		UserRepository:   stubUserRepository,
	}

	policies := extractorService.GetPolicies(1, 1)
	if len(policies) > 0 {
		t.Errorf("Incorrect TestGetPolicies test")
		t.FailNow()
	}
}

// Test get permissions
func TestGetPermissions(t *testing.T) {
	stubPermissionRepository := driver.RdbmsPermissionRepository{Connection: stubConnection}
	extractorService = ExtractorServiceImpl{
		PermissionRepository: stubPermissionRepository,
	}

	policies := extractorService.GetPermissions(1, 1)
	if len(policies) > 0 {
		t.Errorf("Incorrect TestGetPermissions test")
		t.FailNow()
	}
}

// Test get roles
func TestGetRoles(t *testing.T) {
	stubRoleRepository := driver.RdbmsRoleRepository{Connection: stubConnection}
	extractorService = ExtractorServiceImpl{
		RoleRepository: stubRoleRepository,
	}

	roles := extractorService.GetRoles(1, 1)
	if len(roles) > 0 {
		t.Errorf("Incorrect TestGetRoles test")
		t.FailNow()
	}
}

// Test get services
func TestGetServices(t *testing.T) {
	stubServiceRepository := driver.RdbmsServiceRepository{Connection: stubConnection}
	extractorService = ExtractorServiceImpl{
		ServiceRepository: stubServiceRepository,
	}

	services := extractorService.GetServices(1, 1)
	if len(services) > 0 {
		t.Errorf("Incorrect TestGetServices test")
		t.FailNow()
	}
}

// Test get user services
func TestGetUserServices(t *testing.T) {
	stubUserRepository := driver.RdbmsUserRepository{Connection: stubConnection}
	extractorService = ExtractorServiceImpl{
		UserRepository: stubUserRepository,
	}

	userServices := extractorService.GetUserServices(1, 1)
	if len(userServices) > 0 {
		t.Errorf("Incorrect TestGetUserServices test")
		t.FailNow()
	}
}

// Test get user groups
func TestGetUserGroups(t *testing.T) {
	stubUserRepository := driver.RdbmsUserRepository{Connection: stubConnection}
	stubUGroupRepository := driver.RdbmsGroupRepository{Connection: stubConnection}
	extractorService = ExtractorServiceImpl{
		UserRepository:  stubUserRepository,
		GroupRepository: stubUGroupRepository,
	}

	userGroups := extractorService.GetUserGroups(1, 1)
	if len(userGroups) > 0 {
		t.Errorf("Incorrect TestGetUserGroups test")
		t.FailNow()
	}
}
