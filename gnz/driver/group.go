package driver

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var grInstance GroupRepository

type GroupRepository interface {
	// Get groups all data
	FindAll() ([]*entity.Group, *model.ErrorResBody)

	// Get group by id
	FindById(id int) (*entity.Group, *model.ErrorResBody)

	// Get group from groups.name
	FindByName(name string) (*entity.Group, *model.ErrorResBody)

	// Get all groups by user_id
	// Join user_groups and groups
	FindByUserId(userId int) ([]*entity.Group, *model.ErrorResBody)

	// Get all groups with user_groups with policy that has user
	// Join user_groups and groups and polices
	FindGroupWithUserWithPolicyGroupsByUserId(userId int) ([]*model.GroupWithUserGroupWithPolicy, *model.ErrorResBody)

	// Get user_groups with policies by user id and group id
	// Join user_groups and groups and polices
	FindGroupWithPolicyByUserIdAndGroupId(userId int, groupId int) (*model.GroupWithUserGroupWithPolicy, *model.ErrorResBody)

	// Generate groups, user_groups, service_groups
	// Transaction mode
	SaveWithRelationalData(
		group entity.Group,
		serviceGroup entity.ServiceGroup,
		userGroup entity.UserGroup,
		groupPermission entity.GroupPermission,
		groupRole entity.GroupRole,
		policy entity.Policy) (*entity.Group, *model.ErrorResBody)
}

// GroupRepository struct
type GroupRepositoryImpl struct {
	Connection *gorm.DB
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetGroupRepositoryInstance() GroupRepository {
	if grInstance == nil {
		grInstance = NewGroupRepository()
	}
	return grInstance
}

// Constructor
func NewGroupRepository() GroupRepository {
	log.Logger.Info("New `GroupRepository` instance")
	return GroupRepositoryImpl{Connection: connection}
}

func (gr GroupRepositoryImpl) FindAll() ([]*entity.Group, *model.ErrorResBody) {
	var groups []*entity.Group
	if err := gr.Connection.Find(&groups).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return groups, nil
}

func (gr GroupRepositoryImpl) FindById(id int) (*entity.Group, *model.ErrorResBody) {
	var group entity.Group
	if err := gr.Connection.Where("id = ?", id).Find(&group).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &group, nil
}

func (gr GroupRepositoryImpl) FindByName(name string) (*entity.Group, *model.ErrorResBody) {
	var group *entity.Group
	if err := gr.Connection.Where("name = ?", name).Find(&group).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return group, nil
}

func (gr GroupRepositoryImpl) FindByUserId(userId int) ([]*entity.Group, *model.ErrorResBody) {
	var groups []*entity.Group

	target := entity.GroupTable.String() + "." + entity.GroupId.String() + "," + entity.GroupTable.String() + "." + entity.GroupUuid.String() + "," +
		entity.GroupTable.String() + "." + entity.GroupName.String() + "," + entity.GroupTable.String() + "." + entity.GroupCreatedAt.String() + "," +
		entity.GroupTable.String() + "." + entity.GroupUpdatedAt.String()
	if err := gr.Connection.Table(entity.UserGroupTable.String()).
		Select(target).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.GroupTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupGroupId,
			entity.GroupTable.String(),
			entity.GroupId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupUserId), userId).
		Scan(&groups).Error; err != nil {

		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found group")
		}
		return nil, model.InternalServerError()
	}

	return groups, nil
}

func (gr GroupRepositoryImpl) FindGroupWithUserWithPolicyGroupsByUserId(userId int) ([]*model.GroupWithUserGroupWithPolicy, *model.ErrorResBody) {
	var groupWithUserGroupWithPolicies []*model.GroupWithUserGroupWithPolicy

	if err := gr.Connection.Table(entity.UserGroupTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.GroupTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupGroupId,
			entity.GroupTable.String(),
			entity.GroupId)).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PolicyTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupId,
			entity.PolicyTable.String(),
			entity.PolicyUserGroupId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupUserId), userId).
		Scan(&groupWithUserGroupWithPolicies).Error; err != nil {

		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found group or policy")
		}

		return nil, model.InternalServerError()
	}

	return groupWithUserGroupWithPolicies, nil
}

func (gr GroupRepositoryImpl) FindGroupWithPolicyByUserIdAndGroupId(userId int, groupId int) (*model.GroupWithUserGroupWithPolicy, *model.ErrorResBody) {
	var groupWithUserGroupWithPolicy model.GroupWithUserGroupWithPolicy

	if err := gr.Connection.Table(entity.UserGroupTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.GroupTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupGroupId,
			entity.GroupTable.String(),
			entity.GroupId)).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PolicyTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupId,
			entity.PolicyTable.String(),
			entity.PolicyUserGroupId)).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupUserId), userId).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupGroupId), groupId).
		Scan(&groupWithUserGroupWithPolicy).Error; err != nil {

		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, model.NotFound("Not found group or policy")
		}

		return nil, model.InternalServerError()
	}

	return &groupWithUserGroupWithPolicy, nil

}

func (gr GroupRepositoryImpl) SaveWithRelationalData(
	group entity.Group,
	serviceGroup entity.ServiceGroup,
	userGroup entity.UserGroup,
	groupPermission entity.GroupPermission,
	groupRole entity.GroupRole,
	policy entity.Policy) (*entity.Group, *model.ErrorResBody) {

	tx := gr.Connection.Begin()

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
	serviceGroup.GroupId = group.Id
	if err := tx.Create(&serviceGroup).Error; err != nil {
		log.Logger.Warn("Failed to save service_groups at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit service_groups data.")
		}

		return nil, model.InternalServerError()
	}

	// Save user_groups
	userGroup.GroupId = group.Id
	if err := tx.Create(&userGroup).Error; err != nil {
		log.Logger.Warn("Failed to save user_groups at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit user_groups data.")
		}

		return nil, model.InternalServerError()
	}

	// Save group_roles
	groupRole.GroupId = group.Id
	if err := tx.Create(&groupRole).Error; err != nil {
		log.Logger.Warn("Failed to save group_roles at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit group_roles data.")
		}

		return nil, model.InternalServerError()
	}

	// Save group_permissions
	groupPermission.GroupId = group.Id
	if err := tx.Create(&groupPermission).Error; err != nil {
		log.Logger.Warn("Failed to save group_permissions at transaction process", err.Error())
		tx.Rollback()
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit group_permissions data.")
		}

		return nil, model.InternalServerError()
	}

	// Save policies
	policy.UserGroupId = userGroup.Id
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
