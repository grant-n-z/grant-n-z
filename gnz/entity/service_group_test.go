package entity

import (
	"strings"
	"testing"
)

func TestServiceGroupString(t *testing.T) {
	table := ServiceGroupTable.String()
	if !strings.EqualFold(table, "service_groups") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := ServiceGroupId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	groupId := ServiceGroupGroupId.String()
	if !strings.EqualFold(groupId, "group_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	serviceId := ServiceGroupServiceId.String()
	if !strings.EqualFold(serviceId, "service_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := ServiceGroupCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := ServiceGroupUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
