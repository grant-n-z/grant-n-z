package handler

import (
	"fmt"
	"strings"
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/tomoyane/grant-n-z/server/entity"
)

const (
	endpointPolicies = "/api/v1/policies"
)

func TestPolicyHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointPolicies, nil)
	recorder := httptest.NewRecorder()

	NewPolicyHandlerHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestPolicyHandlerPost(t *testing.T) {
	id, _ := uuid.NewV4()
	name := fmt.Sprintf("unit_test_%s", id.String())

	policy := entity.Policy{
		Name:         name,
		PermissionId: 1,
		RoleId:       1,
	}

	body, _ := json.Marshal(policy)

	request := httptest.NewRequest(http.MethodPost, endpointPolicies, strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	NewPolicyHandlerHandler().Post(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}
