package users

import (
	"fmt"
	"strings"
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gserver/entity"
)

const (
	endpointUsers = "/api/v1/users"
)

func TestUserHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointUsers, nil)
	recorder := httptest.NewRecorder()

	NewUserHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUserHandlerPost(t *testing.T) {
	id, _ := uuid.NewV4()
	username := "unit_test"
	email := fmt.Sprintf("%s@unittest.com", id.String())
	password := "unittest1234"

	user := entity.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	body, _ := json.Marshal(user)

	request := httptest.NewRequest(http.MethodPost, endpointUsers, strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	NewUserHandler().Post(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}
