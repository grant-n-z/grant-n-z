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

func TestIssueTokenOk(t *testing.T) {
	user := entity.User {
		Email: "test2@gmail.com",
		Password: "test12345",
	}
	userData, _ := json.Marshal(user)

	request := httptest.NewRequest(echo.POST, "/v1/tokens", strings.NewReader(string(userData)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	if assert.NoError(t, PostToken(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	}
}

func TestIssueTokenBadRequest01(t *testing.T) {
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	inCorrectData := `{"key":"value"}`

	request := httptest.NewRequest(echo.POST, "/v1/tokens", strings.NewReader(inCorrectData))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	assert.Error(t, PostToken(c))
}

func TestIssueTokenBadRequest02(t *testing.T) {
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	// Incorrect validation
	user := entity.User {
		Email: "testgmail.com",
		Password: "2131",
	}
	userData, _ := json.Marshal(user)

	request := httptest.NewRequest(echo.POST, "/v1/tokens", strings.NewReader(string(userData)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	assert.Error(t, PostToken(c))
}

func TestIssueTokenUnProcessableEntity01(t *testing.T) {
	user := entity.User {
		Email: "test@gmail.com",
		Password: "test12345",
	}
	userData, _ := json.Marshal(user)

	request := httptest.NewRequest(echo.POST, "/v1/tokens", strings.NewReader(string(userData)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	assert.Error(t, PostToken(c))
}

func TestIssueTokenUnProcessableEntity02(t *testing.T) {
	user := entity.User {
		Email: "test2@gmail.com",
		Password: "testtest",
	}
	userData, _ := json.Marshal(user)

	request := httptest.NewRequest(echo.POST, "/v1/tokens", strings.NewReader(string(userData)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	assert.Error(t, PostToken(c))
}
