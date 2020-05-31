package entity

import (
	"strings"
	"testing"
)

func TestGroupString(t *testing.T) {
	table := GroupTable.String()
	if !strings.EqualFold(table, "groups") {
		t.Errorf("Incorrect Group TestString test")
		t.FailNow()
	}

	id := GroupId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect Group TestString test")
		t.FailNow()
	}

	internalId := GroupInternalId.String()
	if !strings.EqualFold(internalId, "internal_id") {
		t.Errorf("Incorrect Group TestString test")
		t.FailNow()
	}

	uuid := GroupUuid.String()
	if !strings.EqualFold(uuid, "uuid") {
		t.Errorf("Incorrect Group TestString test")
		t.FailNow()
	}

	name := GroupName.String()
	if !strings.EqualFold(name, "name") {
		t.Errorf("Incorrect Group TestString test")
		t.FailNow()
	}

	createdAt := GroupCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect Group TestString test")
		t.FailNow()
	}

	updatedAt := GroupUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect Group TestString test")
		t.FailNow()
	}
}
