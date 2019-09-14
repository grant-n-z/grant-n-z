package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type ServiceGroupRepository interface {
	FindServiceByGroupId(groupId int) ([]*entity.Service, *model.ErrorResponse)

	FindGroupByServiceId(serviceId int) ([]*entity.Group, *model.ErrorResponse)

	Save(group entity.ServiceGroup) (*entity.ServiceGroup, *model.ErrorResponse)
}
