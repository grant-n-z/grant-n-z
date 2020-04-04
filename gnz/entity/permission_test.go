package entity

import (
	"strings"
	"testing"
)

func TestPermissionString(t *testing.T) {
	table := PermissionTable.String()
	if !strings.EqualFold(table, "permissions") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := PermissionId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	uuid := PermissionUuid.String()
	if !strings.EqualFold(uuid, "uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	name := PermissionName.String()
	if !strings.EqualFold(name, "name") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := PermissionCreatedAt.String()
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
