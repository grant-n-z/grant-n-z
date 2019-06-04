package repository

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type RoleRepository interface {
	FindAll() ([]*entity.Role, *model.ErrorResponse)

	FindById(id int) (*entity.Role, *model.ErrorResponse)

	Save(role entity.Role) (*entity.Role, *model.ErrorResponse)
}
