package repository

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type RoleMemberRepository interface {
	FindAll() ([]*entity.RoleMember, *model.ErrorResponse)

	FindByUserId(userId int) ([]*entity.RoleMember, *model.ErrorResponse)

	Save(role entity.RoleMember) (*entity.RoleMember, *model.ErrorResponse)
}
