package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type RoleRepository interface {
	FindAll() ([]*entity.Role, *model.ErrorResBody)

	FindById(id int) (*entity.Role, *model.ErrorResBody)

	Save(role entity.Role) (*entity.Role, *model.ErrorResBody)
}
