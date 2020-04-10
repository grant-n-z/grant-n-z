package model

import (
	"strings"
	"testing"

	"net/http"
)

// Test to json
func TestToJson(t *testing.T) {
	err := ErrorResBody{
		Code:    http.StatusOK,
		Message: "test",
		Detail:  "test",
	}
	jsonStr := err.ToJson()
	if !strings.EqualFold(jsonStr, "{\"code\":200,\"message\":\"test\",\"detail\":\"test\"}") {
		t.Errorf("Incorrect TestToJson test")
		t.FailNow()
	}
}

// Test to json
func TestWriteError(t *testing.T) {
	writer := StubResponseWriter{}
	WriteError(writer, "", http.StatusOK)
}

// Test bad request
func TestBadRequest(t *testing.T) {
	badRequest := BadRequest("test")
	if badRequest == nil {
		t.Errorf("Incorrect TestWritBadRequest test")
		t.FailNow()
	}
}

// Test unauthorized
func TestUnauthorized(t *testing.T) {
	unauthorized := Unauthorized("test")
	if unauthorized == nil {
		t.Errorf("Incorrect TestUnauthorized test")
		t.FailNow()
	}
}

// Test forbidden
func TestForbidden(t *testing.T) {
	forbidden := Forbidden("test")
	if forbidden == nil {
		t.Errorf("Incorrect TestForbidden test")
		t.FailNow()
	}
}

// Test notfound
func TestNotFound(t *testing.T) {
	notfound := NotFound("test")
	if notfound == nil {
		t.Errorf("Incorrect TestNotFound test")
		t.FailNow()
	}
}

// Test conflict
func TestConflict(t *testing.T) {
	conflict := Conflict("test")
	if conflict == nil {
		t.Errorf("Incorrect TestConflict test")
		t.FailNow()
	}
}

// Test method not allowed
func TestMethodNotAllowed(t *testing.T) {
	methodNotAllowed := MethodNotAllowed("test")
	if methodNotAllowed == nil {
		t.Errorf("Incorrect TestMethodNotAllowed test")
		t.FailNow()
	}
}

// Test un processable entity
func TestUnProcessableEntity(t *testing.T) {
	unProcessableEntity := UnProcessableEntity("test")
	if unProcessableEntity == nil {
		t.Errorf("Incorrect TestUnProcessableEntity test")
		t.FailNow()
	}
}

// Test internal server error
func TestInternalServerError(t *testing.T) {
	internalServerError := InternalServerError("test")
	if internalServerError == nil {
		t.Errorf("Incorrect TestInternalServerError test")
		t.FailNow()
	}
}

// Less than stub struct
// ResponseWriter
type StubResponseWriter struct {
}

func (s StubResponseWriter) Header() http.Header {
	return http.Header{}
}

func (s StubResponseWriter) Write([]byte) (int, error) {
	return http.StatusInternalServerError, nil
}

func (s StubResponseWriter) WriteHeader(statusCode int) {
}
