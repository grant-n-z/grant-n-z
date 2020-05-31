package entity

import (
	"strings"
	"testing"
)

func TestUserString(t *testing.T) {
	table := UserTable.String()
	if !strings.EqualFold(table, "users") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := UserId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	internalId := UserInternalId.String()
	if !strings.EqualFold(internalId, "internal_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	uuid := UserUuid.String()
	if !strings.EqualFold(uuid, "uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	username := UserUsername.String()
	if !strings.EqualFold(username, "username") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	email := UserEmail.String()
	if !strings.EqualFold(email, "email") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	password := UserPassword.String()
	if !strings.EqualFold(password, "password") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := UserCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := UserUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
