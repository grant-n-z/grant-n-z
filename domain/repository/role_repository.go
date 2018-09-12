package repository

import "github.com/tomoyane/grant-n-z/domain/entity"

type RoleRepository interface {
	FindByUserUuid(userUuid string) *entity.Role

	Save(role entity.Role) *entity.Role
}
