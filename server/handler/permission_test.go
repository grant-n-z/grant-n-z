package handler

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

const (
	endpointPermissions = "/api/v1/permissions"
)

func TestPermissionHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointPermissions, nil)
	recorder := httptest.NewRecorder()

	NewPermissionHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)

}

func TestPermissionHandlerPost(t *testing.T) {
}
