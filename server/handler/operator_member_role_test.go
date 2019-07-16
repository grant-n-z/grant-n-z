package handler

import (
	"strings"
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/tomoyane/grant-n-z/server/entity"

)

const (
	endpointRoleMembers = "/api/v1/operator-member-roles"
)

func TestOperateMemberRoleHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointRoleMembers, nil)
	recorder := httptest.NewRecorder()

	NewOperatorMemberRoleHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestOperateMemberRoleHandlerPost(t *testing.T) {
	roleMember := entity.OperatorMemberRole{
		RoleId: 1,
		UserId: 1,
	}

	body, _:= json.Marshal(roleMember)

	request := httptest.NewRequest(http.MethodPost, endpointRoleMembers, strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	NewOperatorMemberRoleHandler().Post(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}
