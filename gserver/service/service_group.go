package service

import (
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var sgsInstance ServiceGroupService

type ServiceGroupService interface {
	InsertServiceGroup(serviceGroup *entity.ServiceGroup) (*entity.ServiceGroup, *model.ErrorResBody)
}

type ServiceGroupServiceImpl struct {
	serviceGroupRepository data.ServiceGroupRepository
}

func GetServiceGroupServiceInstance() ServiceGroupService {
	if sgsInstance == nil {
		sgsInstance = NewServiceGroupService()
	}
	return sgsInstance
}

func NewServiceGroupService() ServiceGroupService {
	log.Logger.Info("New `ServiceGroupService` instance")
	log.Logger.Info("Inject `ServiceGroupRepository` to `ServiceGroupService`")
	return ServiceGroupServiceImpl{serviceGroupRepository: data.GetServiceGroupRepositoryInstance(driver.Db)}
}

func (sgs ServiceGroupServiceImpl) InsertServiceGroup(serviceGroup *entity.ServiceGroup) (*entity.ServiceGroup, *model.ErrorResBody) {
	return sgs.serviceGroupRepository.Save(*serviceGroup)
}
