package entity

import (
	"strings"
	"testing"
)

func TestGroupPermissionString(t *testing.T) {
	table := GroupPermissionTable.String()
	if !strings.EqualFold(table, "group_permissions") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := GroupPermissionId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	permissionId := GroupPermissionPermissionId.String()
	if !strings.EqualFold(permissionId, "permission_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	groupId := GroupPermissionGroupId.String()
	if !strings.EqualFold(groupId, "group_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := GroupPermissionCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := GroupPermissionUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
