package handler

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

const (
	endpointServices = "/api/v1/services"
)

func TestServiceHandlerGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, endpointServices, nil)
	recorder := httptest.NewRecorder()

	NewServiceHandler().Get(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestServiceHandlerPost(t *testing.T) {
	id, _ := uuid.NewV4()
	name := fmt.Sprintf("unit_test_%s", id.String())

	service := entity.Service{
		Name: name,
	}

	body, _:= json.Marshal(service)

	request := httptest.NewRequest(http.MethodPost, endpointServices, strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	NewServiceHandler().Post(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}
