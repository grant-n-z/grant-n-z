package model

import (
	"strings"
	"testing"
)

// Test Builder set name
func TestPolicyResponse_SetName(t *testing.T) {
	response := NewPolicyResponse()
	name := "test"
	builder := response.SetName(&name)
	if !strings.EqualFold(builder.Build().Name, name) {
		t.Errorf("Incorrect TestPolicyResponse_SetName test")
		t.FailNow()
	}

	builder = response.SetName(nil)
	if !strings.EqualFold(builder.Build().Name, "") {
		t.Errorf("Incorrect TestPolicyResponse_SetName test")
		t.FailNow()
	}
}

// Test Builder set role name
func TestPolicyResponse_SetRoleName(t *testing.T) {
	response := NewPolicyResponse()
	roleName := "test"
	builder := response.SetRoleName(&roleName)
	if !strings.EqualFold(builder.Build().RoleName, roleName) {
		t.Errorf("Incorrect TestPolicyResponse_SetRoleName test")
		t.FailNow()
	}

	builder = response.SetRoleName(nil)
	if !strings.EqualFold(builder.Build().RoleName, "") {
		t.Errorf("Incorrect TestPolicyResponse_SetRoleName test")
		t.FailNow()
	}
}

// Test Builder set permission name
func TestPolicyResponse_SetPermissionName(t *testing.T) {
	response := NewPolicyResponse()
	permissionName := "test"
	builder := response.SetPermissionName(&permissionName)
	if !strings.EqualFold(builder.Build().PermissionName, permissionName) {
		t.Errorf("Incorrect TestPolicyResponse_SetPermissionName test")
		t.FailNow()
	}

	builder = response.SetPermissionName(nil)
	if !strings.EqualFold(builder.Build().PermissionName, "") {
		t.Errorf("Incorrect TestPolicyResponse_SetPermissionName test")
		t.FailNow()
	}
}

// Test Builder set service name
func TestPolicyResponse_SetServiceName(t *testing.T) {
	response := NewPolicyResponse()
	serviceName := "test"
	builder := response.SetServiceName(&serviceName)
	if !strings.EqualFold(builder.Build().ServiceName, serviceName) {
		t.Errorf("Incorrect TestPolicyResponse_SetServiceName test")
		t.FailNow()
	}

	builder = response.SetServiceName(nil)
	if !strings.EqualFold(builder.Build().ServiceName, "") {
		t.Errorf("Incorrect TestPolicyResponse_SetServiceName test")
		t.FailNow()
	}
}

// Test Builder set group name
func TestPolicyResponse_SetGroupName(t *testing.T) {
	response := NewPolicyResponse()
	groupName := "test"
	builder := response.SetGroupName(&groupName)
	if !strings.EqualFold(builder.Build().GroupName, groupName) {
		t.Errorf("Incorrect TestPolicyResponse_SetGroupName test")
		t.FailNow()
	}

	builder = response.SetGroupName(nil)
	if !strings.EqualFold(builder.Build().GroupName, "") {
		t.Errorf("Incorrect TestPolicyResponse_SetGroupName test")
		t.FailNow()
	}
}
