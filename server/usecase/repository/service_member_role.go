package repository

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type ServiceMemberRoleRepository interface {
	FindAll() ([]*entity.ServiceMemberRole, *model.ErrorResponse)

	FindById(id int) ([]*entity.ServiceMemberRole, *model.ErrorResponse)

	FindByRoleId(roleId int) ([]*entity.ServiceMemberRole, *model.ErrorResponse)

	FindByUserServiceId(userServiceId int) ([]*entity.ServiceMemberRole, *model.ErrorResponse)

	Save(serviceMemberRole entity.ServiceMemberRole) (*entity.ServiceMemberRole, *model.ErrorResponse)
}
