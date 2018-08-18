package controller

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"strings"
	"src/github.com/stretchr/testify/assert"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain"
	"gopkg.in/go-playground/validator.v9"
	"encoding/json"
	"github.com/tomoyane/grant-n-z/controller"
)

var (
	user = entity.User {
		Username: "test123456789",
		Email: "test@gmail.com",
		Password: "21312abcdefg",
	}
)

func TestCreateUser(t *testing.T) {
	e := echo.New()
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	userData, _ := json.Marshal(user)

	request := httptest.NewRequest(echo.POST, "/v1/users", strings.NewReader(string(userData)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	if assert.NoError(t,  controller.GenerateUser(c)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}
}