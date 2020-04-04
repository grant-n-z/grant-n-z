package entity

import (
	"strings"
	"testing"
)

func TestUserGroupString(t *testing.T) {
	table := UserGroupTable.String()
	if !strings.EqualFold(table, "user_groups") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := UserGroupId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	userId := UserGroupUserId.String()
	if !strings.EqualFold(userId, "user_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	groupId := UserGroupGroupId.String()
	if !strings.EqualFold(groupId, "group_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := UserGroupCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := UserGroupUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
