package groups

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tomoyane/grant-n-z/gnzserver/entity"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

const (
	endpointRoles = "/api/v1/roles"
)

func TestRoleHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointRoles, nil)
	recorder := httptest.NewRecorder()

	NewRoleHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestRoleHandlerPost(t *testing.T) {
	id, _ := uuid.NewV4()
	name := fmt.Sprintf("unit_test_%s", id.String())

	role := entity.Role{
		Name: name,
	}

	body, _ := json.Marshal(role)

	request := httptest.NewRequest(http.MethodPost, endpointRoles, strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	NewRoleHandler().Post(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}
