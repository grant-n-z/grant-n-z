package handler

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

const (
	endpointRoleMembers = "/api/v1/role-members"
)

func TestRoleMemberHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointRoleMembers, nil)
	recorder := httptest.NewRecorder()

	NewRoleMemberHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestRoleMemberHandlerPost(t *testing.T) {
}
