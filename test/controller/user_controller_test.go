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
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/test/stub"
	"os"
)

var(
	e = echo.New()
)

func TestMain(m *testing.M) {
	di.InitUserService(stub.UserRepositoryStub{})
	e := echo.New()
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	code := m.Run()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	user := entity.User {
		Username: "test",
		Email: "test1@gmail.com",
		Password: "21312abcdefg",
	}
	userData, _ := json.Marshal(user)

	request := httptest.NewRequest(echo.POST, "/v1/users", strings.NewReader(string(userData)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	if assert.NoError(t, controller.PostUser(c)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}
}

func TestCreateUserBadRequest01(t *testing.T) {
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	inCorrectData := `{"key":"value"}`

	request := httptest.NewRequest(echo.POST, "/v1/users", strings.NewReader(inCorrectData))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	assert.Error(t, controller.PostUser(c))
}

func TestCreateUserBadRequest02(t *testing.T) {
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	// Incorrect validation
	user := entity.User {
		Username: "test123456789",
		Email: "testgmail.com",
		Password: "2131",
	}
	userData, _ := json.Marshal(user)

	request := httptest.NewRequest(echo.POST, "/v1/users", strings.NewReader(string(userData)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	assert.Error(t, controller.PostUser(c))
}

func TestCreateUserUnprocessableEntity(t *testing.T) {
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	// Already exit user
	user := entity.User {
		Username: "test123456789",
		Email: "test2@gmail.com",
		Password: "21312abcdefg",
	}
	userData, _ := json.Marshal(user)

	request := httptest.NewRequest(echo.POST, "/v1/users", strings.NewReader(string(userData)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	assert.Error(t, controller.PostUser(c))
}