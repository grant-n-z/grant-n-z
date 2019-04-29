package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/tomoyane/grant-n-z/server/di"
	"testing"
)

func TestGetRoleByUserUuid(t *testing.T) {
	role := di.ProviderRoleService.GetRoleByUserUuid(userUuid.String())
	assert.NotEmpty(t, role)
	assert.Equal(t, role.Type, "user")
	assert.Equal(t, role.UserUuid, userUuid)
}

func TestInsertRole(t *testing.T) {
	roleData := di.ProviderRoleService.InsertRole(userUuid)

	assert.NotEmpty(t, roleData)
	assert.Equal(t, roleData.Type, "user")
	assert.Equal(t, roleData.UserUuid, userUuid)
}
