package test

import (
	"testing"
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"encoding/json"
	"net/http/httptest"
	"src/github.com/stretchr/testify/assert"
	"github.com/tomoyane/grant-n-z/controller"
	"net/http"
	"strings"
	"github.com/tomoyane/grant-n-z/domain"
	"gopkg.in/go-playground/validator.v9"
)

func TestGenerateUser(t *testing.T) {
	e := echo.New()
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	user := entity.User{
		Username: "test123456789",
		Email: "test@gmail.com",
		Password: "21312abcdefg",
	}

	userData, _ := json.Marshal(user)

	request := httptest.NewRequest(echo.POST, "/v1/users", strings.NewReader(string(userData)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	if assert.NoError(t,  controller.GenerateUser(c)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}
}
