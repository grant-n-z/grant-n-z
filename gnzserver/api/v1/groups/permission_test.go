package groups

import (
	"fmt"
	"strings"
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/tomoyane/grant-n-z/gnzserver/entity"
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
	id, _ := uuid.NewV4()
	name := fmt.Sprintf("unit_test_%s", id.String())

	permission := entity.Permission{
		Name: name,
	}

	body, _ := json.Marshal(permission)

	request := httptest.NewRequest(http.MethodPost, endpointPermissions, strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	NewPermissionHandler().Post(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)

}
