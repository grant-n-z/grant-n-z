package repository

import "github.com/tomoyane/grant-n-z/server/domain/entity"

type RoleRepository interface {
	FindByUserUuid(userUuid string) *entity.Role

	Save(role entity.Role) *entity.Role
}
