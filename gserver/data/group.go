package data

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var grInstance GroupRepository

type GroupRepository interface {
	FindAll() ([]*entity.Group, *model.ErrorResBody)

	FindByName(name string) (*entity.Group, *model.ErrorResBody)

	Save(group entity.Group) (*entity.Group, *model.ErrorResBody)
}

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


func (gr GroupRepositoryImpl) FindAll() ([]*entity.Group, *model.ErrorResBody) {
	var groups []*entity.Group
	if err := gr.Db.Find(&groups).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return groups, nil
}

func (gr GroupRepositoryImpl) FindByName(name string) (*entity.Group, *model.ErrorResBody) {
	var groups *entity.Group
	if err := gr.Db.Where("name = ?", name).Find(&groups).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return groups, nil
}

func (gr GroupRepositoryImpl) Save(group entity.Group) (*entity.Group, *model.ErrorResBody) {
	if err := gr.Db.Create(&group).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		}

		return nil, model.InternalServerError()
	}

	return &group, nil
}
