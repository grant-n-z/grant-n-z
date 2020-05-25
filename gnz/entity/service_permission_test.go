package entity

import (
	"strings"
	"testing"
)

func TestServicePermissionString(t *testing.T) {
	table := ServicePermissionTable.String()
	if !strings.EqualFold(table, "service_permissions") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := ServicePermissionId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	permissionId := ServicePermissionPermissionUuid.String()
	if !strings.EqualFold(permissionId, "permission_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	serviceId := ServicePermissionServiceUuid.String()
	if !strings.EqualFold(serviceId, "service_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := ServicePermissionCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := ServicePermissionUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
