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

	roleId := OperatorPolicyRoleId.String()
	if !strings.EqualFold(roleId, "role_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	userId := OperatorPolicyUserId.String()
	if !strings.EqualFold(userId, "user_id") {
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
