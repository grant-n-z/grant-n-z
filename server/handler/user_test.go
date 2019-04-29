package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/tomoyane/grant-n-z/server/domain"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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

	if assert.NoError(t, PostUser(c)) {
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

	assert.Error(t, PostUser(c))
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

	assert.Error(t, PostUser(c))
}

func TestCreateUserUnProcessableEntity(t *testing.T) {
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

	assert.Error(t, PostUser(c))
}
