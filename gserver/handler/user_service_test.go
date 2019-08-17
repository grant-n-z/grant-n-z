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
	endpointUserServices = "/api/v1/user-services"
)

func TestUserServiceHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointUserServices, nil)
	recorder := httptest.NewRecorder()

	NewUserServiceHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUserServiceHandlerPost(t *testing.T) {
	userService := entity.UserService{
		UserId: 1,
		ServiceId: 1,
	}

	body, _:= json.Marshal(userService)

	request := httptest.NewRequest(http.MethodPost, endpointUserServices, strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	NewUserServiceHandler().Post(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}
