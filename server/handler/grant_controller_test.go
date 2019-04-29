package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGrantUnauthorized01(t *testing.T) {
	request := httptest.NewRequest(echo.POST, "/v1/grants", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	assert.Error(t, PostToken(c))
}
