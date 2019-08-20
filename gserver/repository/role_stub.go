package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
)

type RoleRepositoryStub struct {
}

func (r RoleRepositoryStub) FindByUserUuid(userUuidStr string) *entity.Role {
	role := entity.Role{
		Id: 1,
	}
	return &role
}

func (r RoleRepositoryStub) Save(role entity.Role) *entity.Role {
	return &role
}
