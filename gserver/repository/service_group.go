package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var sgrInstance ServiceGroupRepository

type ServiceGroupRepository interface {
	FindServiceByGroupId(groupId int) ([]*entity.Service, *model.ErrorResBody)

	FindGroupByServiceId(serviceId int) ([]*entity.Group, *model.ErrorResBody)

	Save(group entity.ServiceGroup) (*entity.ServiceGroup, *model.ErrorResBody)
}

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

func (sgr ServiceGroupRepositoryImpl) FindServiceByGroupId(groupId int) ([]*entity.Service, *model.ErrorResBody) {
	return nil, nil
}

func (sgr ServiceGroupRepositoryImpl) FindGroupByServiceId(serviceId int) ([]*entity.Group, *model.ErrorResBody) {
	return nil, nil
}

func (sgr ServiceGroupRepositoryImpl) Save(serviceGroup entity.ServiceGroup) (*entity.ServiceGroup, *model.ErrorResBody) {
	if err := sgr.Db.Create(&serviceGroup).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		}

		return nil, model.InternalServerError()
	}

	return &serviceGroup, nil
}
