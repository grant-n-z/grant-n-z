package controller

import (
	"testing"
	"net/http/httptest"
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/app/controller"
	"github.com/stretchr/testify/assert"
)

func TestGrantUnauthorized01(t *testing.T) {
	request := httptest.NewRequest(echo.POST, "/v1/grants", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	assert.Error(t, controller.PostToken(c))
}
