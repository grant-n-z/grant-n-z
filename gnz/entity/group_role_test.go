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

	roleId := GroupRoleRoleId.String()
	if !strings.EqualFold(roleId, "role_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	groupId := GroupRoleGroupId.String()
	if !strings.EqualFold(groupId, "group_id") {
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
