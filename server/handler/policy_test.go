package handler

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
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
}
