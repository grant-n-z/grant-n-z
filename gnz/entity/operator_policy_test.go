package entity

import (
	"strings"
	"testing"
)

func TestOperatorPolicyString(t *testing.T) {
	table := OperatorPolicyTable.String()
	if !strings.EqualFold(table, "operator_policies") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := OperatorPolicyId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	internalId := OperatorPolicyInternalId.String()
	if !strings.EqualFold(internalId, "internal_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	roleUuid := OperatorPolicyRoleUuid.String()
	if !strings.EqualFold(roleUuid, "role_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	userUuid := OperatorPolicyUserUuid.String()
	if !strings.EqualFold(userUuid, "user_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := OperatorPolicyCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := OperatorPolicyUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
