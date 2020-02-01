package data

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/common/constant"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var grInstance GroupRepository

type GroupRepository interface {
	// Get groups all data
	FindAll() ([]*entity.Group, *model.ErrorResBody)

	// Get group by id
	FindById(id int) (*entity.Group, *model.ErrorResBody)

	// Get group from groups.name
	FindByName(name string) (*entity.Group, *model.ErrorResBody)

	// Generate groups, user_groups, service_groups
	// Transaction mode
	SaveWithRelationalData(group entity.Group, roleId int, permissionId int, serviceId int, userId int) (*entity.Group, *model.ErrorResBody)
}

// GroupRepository struct
type GroupRepositoryImpl struct {
	Db *gorm.DB
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetGroupRepositoryInstance(db *gorm.DB) GroupRepository {
	if grInstance == nil {
		grInstance = NewGroupRepository(db)
	}
	return grInstance
}

// Constructor
func NewGroupRepository(db *gorm.DB) GroupRepository {
	log.Logger.Info("New `GroupRepository` instance")
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

func (gr GroupRepositoryImpl) FindById(id int) (*entity.Group, *model.ErrorResBody) {
	var group entity.Group
	if err := gr.Db.Where("id = ?", id).Find(&group).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &group, nil
}

func (gr GroupRepositoryImpl) FindByName(name string) (*entity.Group, *model.ErrorResBody) {
	var group *entity.Group
	if err := gr.Db.Where("name = ?", name).Find(&group).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return group, nil
}

func (gr GroupRepositoryImpl) SaveWithRelationalData(group entity.Group, roleId int, permissionId int, serviceId int, userId int) (*entity.Group, *model.ErrorResBody) {
	tx := gr.Db.Begin()

	// Save groups
	if err := tx.Create(&group).Error; err != nil {
		log.Logger.Warn("Failed to save groups at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit groups data.")
		}

		return nil, model.InternalServerError()
	}

	// Save service_groups
	serviceGroup := entity.ServiceGroup{
		GroupId:   group.Id,
		ServiceId: serviceId,
	}
	if err := tx.Create(&serviceGroup).Error; err != nil {
		log.Logger.Warn("Failed to save service_groups at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit service_groups data.")
		}

		return nil, model.InternalServerError()
	}

	// Save user_groups
	userGroup := entity.UserGroup{
		UserId:  userId,
		GroupId: group.Id,
	}
	if err := tx.Create(&userGroup).Error; err != nil {
		log.Logger.Warn("Failed to save user_groups at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit user_groups data.")
		}

		return nil, model.InternalServerError()
	}

	// Save group_roles
	groupRole := entity.GroupRole{
		RoleId:  roleId,
		GroupId: group.Id,
	}
	if err := tx.Create(&groupRole).Error; err != nil {
		log.Logger.Warn("Failed to save group_roles at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit group_roles data.")
		}

		return nil, model.InternalServerError()
	}

	// Save group_permissions
	groupPermission := entity.GroupPermission{
		PermissionId: permissionId,
		GroupId:      group.Id,
	}
	if err := tx.Create(&groupPermission).Error; err != nil {
		log.Logger.Warn("Failed to save group_permissions at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit group_permissions data.")
		}

		return nil, model.InternalServerError()
	}

	// Save policies
	policy := entity.Policy{
		Name:         constant.AdminPolicy,
		RoleId:       roleId,
		PermissionId: permissionId,
		ServiceId:    serviceId,
		UserGroupId:  userGroup.Id,
	}
	if err := tx.Create(&policy).Error; err != nil {
		log.Logger.Warn("Failed to save policies at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit policies data.")
		}

		return nil, model.InternalServerError()
	}

	tx.Commit()

	return &group, nil
}
