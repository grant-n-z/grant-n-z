package repository

import "github.com/tomoyane/grant-n-z/server/domain/entity"

type RoleRepository interface {
	Save(role entity.Role) (*entity.Role, *entity.ErrorResponse)
}
