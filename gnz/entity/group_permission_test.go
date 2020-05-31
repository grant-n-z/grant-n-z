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

	internalId := GroupPermissionInternalId.String()
	if !strings.EqualFold(internalId, "internal_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	permissionUuid := GroupPermissionPermissionUuid.String()
	if !strings.EqualFold(permissionUuid, "permission_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	groupUuid := GroupPermissionGroupUuid.String()
	if !strings.EqualFold(groupUuid, "group_uuid") {
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
