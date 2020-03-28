package ctx

import (
	"github.com/google/uuid"
	"strings"
	"testing"
)

// GetCtx test
func TestGetCtx(t *testing.T) {
	InitContext()

	if ctx == nil {
		t.Errorf("Incorrect GetCtx test. ctx is null")
	}
}

// SetApiKey, GetApiKey test
func TestApiKey(t *testing.T) {
	InitContext()
	SetApiKey("test_api_key")
	apiKey := GetApiKey()
	if !strings.EqualFold(apiKey.(string), "test_api_key") {
		t.Errorf("Incorrect SetApiKey, GetApiKey test. ApiKey is not `test_api_key`")
	}
}

// SetServiceId, GetServiceId test
func TestServiceId(t *testing.T) {
	InitContext()
	SetServiceId(1)
	serviceId := GetServiceId()
	if serviceId != 1 {
		t.Errorf("Incorrect SetServiceId, GetServiceId test. ServiceId is not `1`")
	}
}

// SetUserId, GetUserId test
func TestUserId(t *testing.T) {
	InitContext()
	SetUserId(1)
	userId := GetUserId()
	if userId != 1 {
		t.Errorf("Incorrect SetUserId, GetUserId test. UserId is not `1`")
	}
}

// SetUserUuid, GetUserUuid test
func TestUserUuid(t *testing.T) {
	InitContext()

	uid := uuid.New()
	SetUserUuid(uid)
	userUuid := GetUserUuid()

	if userUuid != uid {
		t.Errorf("Incorrect SetUserUuid, GetUserUuid test. UserUuid is not %s", uid)
	}
}
