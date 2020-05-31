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

	internalId := PolicyInternalId.String()
	if !strings.EqualFold(internalId, "internal_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	name := PolicyName.String()
	if !strings.EqualFold(name, "name") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	roleId := PolicyRoleUuid.String()
	if !strings.EqualFold(roleId, "role_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	permissionId := PolicyPermissionUuid.String()
	if !strings.EqualFold(permissionId, "permission_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	serviceId := PolicyServiceUuid.String()
	if !strings.EqualFold(serviceId, "service_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	userGroupId := PolicyUserGroupUuid.String()
	if !strings.EqualFold(userGroupId, "user_group_uuid") {
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
