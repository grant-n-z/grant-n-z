package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type ServiceMemberRoleService interface {
	GetAll() ([]*entity.ServiceMemberRole, *model.ErrorResponse)

	GetByRoleId(roleId int) ([]*entity.ServiceMemberRole, *model.ErrorResponse)

	GetByUserServiceId(userServiceId int) ([]*entity.ServiceMemberRole, *model.ErrorResponse)

	Insert(serviceMemberRole *entity.ServiceMemberRole) (*entity.ServiceMemberRole, *model.ErrorResponse)
}
