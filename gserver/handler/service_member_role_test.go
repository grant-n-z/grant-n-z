package handler

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
	endpointServiceMemberRoles = "/api/v1/service-member-roles"
)

func TestServiceMemberRoleHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointServiceMemberRoles, nil)
	recorder := httptest.NewRecorder()

	NewOperatorPolicyHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestServiceMemberRoleHandlerPost(t *testing.T) {
	roleMember := entity.OperatorPolicy{
		RoleId: 1,
		UserId: 1,
	}

	body, _:= json.Marshal(roleMember)

	request := httptest.NewRequest(http.MethodPost, endpointServiceMemberRoles, strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	NewOperatorPolicyHandler().Post(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}
