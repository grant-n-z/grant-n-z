package repository

import "github.com/tomoyane/grant-n-z/server/domain/entity"

type RoleRepository interface {
	FindAll() ([]*entity.Role, *entity.ErrorResponse)

	FindById(id int) (*entity.Role, *entity.ErrorResponse)

	Save(role entity.Role) (*entity.Role, *entity.ErrorResponse)
}
