package entity

import (
	"strings"
	"testing"
)

func TestRoleString(t *testing.T) {
	table := RoleTable.String()
	if !strings.EqualFold(table, "roles") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := RoleId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	uuid := RoleUuid.String()
	if !strings.EqualFold(uuid, "uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	name := RoleName.String()
	if !strings.EqualFold(name, "name") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := RoleCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := RoleUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
