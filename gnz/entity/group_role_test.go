package entity

import (
	"strings"
	"testing"
)

func TestGroupRoleString(t *testing.T) {
	table := GroupRoleTable.String()
	if !strings.EqualFold(table, "group_roles") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := GroupRoleId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	roleUuid := GroupRoleRoleUuid.String()
	if !strings.EqualFold(roleUuid, "role_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	groupUuid := GroupRoleGroupUuid.String()
	if !strings.EqualFold(groupUuid, "group_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := GroupRoleCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := GroupRoleUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
