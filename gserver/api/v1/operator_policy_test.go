package v1

import (
	"strings"
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/tomoyane/grant-n-z/gserver/entity"

)

const (
	endpointOperatorMemberRoles = "/api/v1/operator-member-roles"
)

func TestOperateMemberRoleHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointOperatorMemberRoles, nil)
	recorder := httptest.NewRecorder()

	NewOperatorPolicyHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestOperateMemberRoleHandlerPost(t *testing.T) {
	roleMember := entity.OperatorPolicy{
		RoleId: 1,
		UserId: 1,
	}

	body, _:= json.Marshal(roleMember)

	request := httptest.NewRequest(http.MethodPost, endpointOperatorMemberRoles, strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	NewOperatorPolicyHandler().Post(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}
