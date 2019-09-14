package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var grInstance GroupRepository

type GroupRepositoryImpl struct {
	Db *gorm.DB
}

func GetGroupRepositoryInstance(db *gorm.DB) GroupRepository {
	if grInstance == nil {
		grInstance = NewGroupRepository(db)
	}
	return grInstance
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	log.Logger.Info("New `GroupRepository` instance")
	log.Logger.Info("Inject `gorm.DB` to `GroupRepository`")
	return GroupRepositoryImpl{Db: db}
}


func (gr GroupRepositoryImpl) FindAll() ([]*entity.Group, *model.ErrorResponse) {

}

func (gr GroupRepositoryImpl) FindByName(name string) ([]*entity.Group, *model.ErrorResponse) {

}

func (gr GroupRepositoryImpl) Save(group entity.Group) (*entity.Group, *model.ErrorResponse) {

}
