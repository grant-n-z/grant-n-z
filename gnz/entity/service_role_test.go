package entity

import (
	"strings"
	"testing"
)

func TestServiceRoleString(t *testing.T) {
	table := ServiceRoleTable.String()
	if !strings.EqualFold(table, "service_roles") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := ServiceRoleId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	roleId := ServiceRoleRoleId.String()
	if !strings.EqualFold(roleId, "role_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	serviceId := ServiceRoleServiceId.String()
	if !strings.EqualFold(serviceId, "service_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := ServiceRoleCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := ServiceRoleUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}