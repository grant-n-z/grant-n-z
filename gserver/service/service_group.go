package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type ServiceGroupService interface {
	InsertServiceGroup(serviceGroup *entity.ServiceGroup) (*entity.ServiceGroup, *model.ErrorResponse)
}
