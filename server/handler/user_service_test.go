package handler

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

const (
	endpointUserServices = "/api/v1/user-services"
)

func TestUserServiceHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointUserServices, nil)
	recorder := httptest.NewRecorder()

	NewUserServiceHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUserServiceHandlerPost(t *testing.T) {
}
