package entity

import (
	"strings"
	"testing"
)

func TestPolicyString(t *testing.T) {
	table := PolicyTable.String()
	if !strings.EqualFold(table, "policies") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := PolicyId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	name := PolicyName.String()
	if !strings.EqualFold(name, "name") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	roleId := PolicyRoleId.String()
	if !strings.EqualFold(roleId, "role_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	permissionId := PolicyPermissionId.String()
	if !strings.EqualFold(permissionId, "permission_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	serviceId := PolicyServiceId.String()
	if !strings.EqualFold(serviceId, "service_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	userGroupId := PolicyUserGroupId.String()
	if !strings.EqualFold(userGroupId, "user_group_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := PolicyCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := PermissionUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
