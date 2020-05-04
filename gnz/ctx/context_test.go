package ctx

import (
	"github.com/google/uuid"
	"strings"
	"testing"
)

// GetCtx test
func TestGetCtx(t *testing.T) {
	InitContext()

	if GetCtx() == nil {
		t.Errorf("Incorrect GetCtx test. ctx is null")
		t.FailNow()
	}
}

// SetClientSecret, GetClientSecret test
func TestApiKey(t *testing.T) {
	InitContext()
	SetClientSecret("test_api_key")
	apiKey := GetClientSecret()
	if !strings.EqualFold(apiKey.(string), "test_api_key") {
		t.Errorf("Incorrect SetClientSecret, GetClientSecret test. Secret is not `test_api_key`")
		t.FailNow()
	}
}

// SetServiceId, GetServiceId test
func TestServiceId(t *testing.T) {
	InitContext()
	SetServiceId(1)
	serviceId := GetServiceId()
	if serviceId != 1 {
		t.Errorf("Incorrect SetServiceId, GetServiceId test. ServiceId is not `1`")
		t.FailNow()
	}
}

// SetUserId, GetUserId test
func TestUserId(t *testing.T) {
	InitContext()
	SetUserId(1)
	userId := GetUserId()
	if userId != 1 {
		t.Errorf("Incorrect SetUserId, GetUserId test. UserId is not `1`")
		t.FailNow()
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
		t.FailNow()
	}
}
