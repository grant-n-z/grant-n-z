package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type ServiceMemberRoleService interface {
	GetAll() ([]*entity.ServiceMemberRole, *model.ErrorResponse)

	GetByRoleId(roleId int) ([]*entity.ServiceMemberRole, *model.ErrorResponse)

	GetByUserServiceId(userServiceId int) ([]*entity.ServiceMemberRole, *model.ErrorResponse)

	Insert(serviceMemberRole *entity.ServiceMemberRole) (*entity.ServiceMemberRole, *model.ErrorResponse)
}
