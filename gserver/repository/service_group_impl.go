package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var sgrInstance ServiceGroupRepository

type ServiceGroupRepositoryImpl struct {
	Db *gorm.DB
}

func GetServiceGroupRepositoryInstance(db *gorm.DB) ServiceGroupRepository {
	if sgrInstance == nil {
		sgrInstance = NewServiceGroupRepository(db)
	}
	return sgrInstance
}

func NewServiceGroupRepository(db *gorm.DB) ServiceGroupRepository {
	log.Logger.Info("New `ServiceGroupRepository` instance")
	log.Logger.Info("Inject `gorm.DB` to `ServiceGroupRepository`")
	return ServiceGroupRepositoryImpl{Db: db}
}

func (sgr ServiceGroupRepositoryImpl) FindServiceByGroupId(groupId int) ([]*entity.Service, *model.ErrorResponse) {

}

func (sgr ServiceGroupRepositoryImpl) FindGroupByServiceId(serviceId int) ([]*entity.Group, *model.ErrorResponse) {

}

func (sgr ServiceGroupRepositoryImpl) Save(group entity.ServiceGroup) (*entity.ServiceGroup, *model.ErrorResponse) {

}
