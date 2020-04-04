package entity

import (
	"strings"
	"testing"
)

func TestServiceString(t *testing.T) {
	table := ServiceTable.String()
	if !strings.EqualFold(table, "services") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := ServiceId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	uuid := ServiceUuid.String()
	if !strings.EqualFold(uuid, "uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	name := ServiceName.String()
	if !strings.EqualFold(name, "name") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	apiKey := ServiceApiKey.String()
	if !strings.EqualFold(apiKey, "api_key") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := ServiceCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := ServiceUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
