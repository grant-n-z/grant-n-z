package stub

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/satori/go.uuid"
)

type RoleRepositoryStub struct {
}

func (r RoleRepositoryStub) FindByUserUuid(userUuidStr string) *entity.Role {
	userUuid, _ := uuid.FromString(userUuidStr)
	role := entity.Role{
		Id: 1,
		Type: "Bearer",
		UserUuid: userUuid,
	}
	return &role
}

func (r RoleRepositoryStub) Save(role entity.Role) *entity.Role {
	return &role
}