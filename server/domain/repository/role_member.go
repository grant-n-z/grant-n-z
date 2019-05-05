package repository

import "github.com/tomoyane/grant-n-z/server/domain/entity"

type RoleMemberRepository interface {
	FindAll() ([]*entity.RoleMember, *entity.ErrorResponse)

	FindByUserId(userId int) ([]*entity.RoleMember, *entity.ErrorResponse)

	Save(role entity.RoleMember) (*entity.RoleMember, *entity.ErrorResponse)
}
