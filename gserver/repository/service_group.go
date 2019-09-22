package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type ServiceGroupRepository interface {
	FindServiceByGroupId(groupId int) ([]*entity.Service, *model.ErrorResBody)

	FindGroupByServiceId(serviceId int) ([]*entity.Group, *model.ErrorResBody)

	Save(group entity.ServiceGroup) (*entity.ServiceGroup, *model.ErrorResBody)
}
