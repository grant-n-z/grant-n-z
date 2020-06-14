package driver

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var grInstance GroupRepository

type GroupRepository interface {
	// Get groups all data
	FindAll() ([]*entity.Group, error)

	// Get group by uuid
	FindByUuid(uuid string) (*entity.Group, error)

	// Get group from groups.name
	FindByName(name string) (*entity.Group, error)

	// Get all groups by user uuid
	// Join user_groups and groups
	FindByUserUuid(userUuid string) ([]*entity.Group, error)

	// Get all groups by service uuid
	// Join service_groups and groups
	FindByServiceUuid(serviceUuid string) ([]*entity.Group, error)

	// Get all groups with user_groups with policy that has user
	// Join user_groups and groups and polices
	FindGroupWithUserWithPolicyGroupsByUserUuid(userUuid string) ([]*model.GroupWithUserGroupWithPolicy, error)

	// Get user_groups with policies by user id and group id
	// Join user_groups and groups and polices
	FindGroupWithPolicyByUserUuidAndGroupUuid(userUuid string, groupUuid string) (*model.GroupWithUserGroupWithPolicy, error)

	// Generate groups, user_groups, service_groups
	// Transaction mode
	SaveWithRelationalData(
		group entity.Group,
		serviceGroup entity.ServiceGroup,
		userGroup entity.UserGroup,
		groupPermission entity.GroupPermission,
		groupRole entity.GroupRole,
		policy entity.Policy) (*entity.Group, error)
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

func (gr GroupRepositoryImpl) FindAll() ([]*entity.Group, error) {
	var groups []*entity.Group
	if err := gr.Connection.Find(&groups).Error; err != nil {
		return nil, err
	}

	return groups, nil
}

func (gr GroupRepositoryImpl) FindByUuid(uuid string) (*entity.Group, error) {
	var group entity.Group
	if err := gr.Connection.Where("uuid = ?", uuid).Find(&group).Error; err != nil {
		return nil, err
	}

	return &group, nil
}

func (gr GroupRepositoryImpl) FindByName(name string) (*entity.Group, error) {
	var group *entity.Group
	if err := gr.Connection.Where("name = ?", name).Find(&group).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (gr GroupRepositoryImpl) FindByUserUuid(userUuid string) ([]*entity.Group, error) {
	var groups []*entity.Group

	target := entity.GroupTable.String() + "." +
		entity.GroupUuid.String() + "," +
		entity.GroupTable.String() + "." +
		entity.GroupName.String() + "," +
		entity.GroupTable.String() + "." +
		entity.GroupCreatedAt.String() + "," +
		entity.GroupTable.String() + "." +
		entity.GroupUpdatedAt.String()

	if err := gr.Connection.Table(entity.UserGroupTable.String()).
		Select(target).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.GroupTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupGroupUuid.String(),
			entity.GroupTable.String(),
			entity.GroupUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupUserUuid.String()), userUuid).
		Scan(&groups).Error; err != nil {

		return nil, err
	}

	return groups, nil
}

func (gr GroupRepositoryImpl) FindByServiceUuid(serviceUuid string) ([]*entity.Group, error) {
	var groups []*entity.Group

	target := entity.GroupTable.String() + "." +
		entity.GroupId.String() + "," +
		entity.GroupTable.String() + "." +
		entity.GroupInternalId.String() + "," +
		entity.GroupTable.String() + "." +
		entity.GroupUuid.String() + "," +
		entity.GroupTable.String() + "." +
		entity.GroupName.String() + "," +
		entity.GroupTable.String() + "." +
		entity.GroupCreatedAt.String() + "," +
		entity.GroupTable.String() + "." +
		entity.GroupUpdatedAt.String()

	if err := gr.Connection.Table(entity.ServiceGroupTable.String()).
		Select(target).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.GroupTable.String(),
			entity.GroupTable.String(),
			entity.GroupUuid.String(),
			entity.ServiceGroupTable.String(),
			entity.ServiceGroupGroupUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.ServiceGroupTable.String(),
			entity.ServiceGroupServiceUuid.String()), serviceUuid).
		Scan(&groups).Error; err != nil {

		return nil, err
	}

	return groups, nil
}

func (gr GroupRepositoryImpl) FindGroupWithUserWithPolicyGroupsByUserUuid(userUuid string) ([]*model.GroupWithUserGroupWithPolicy, error) {
	var groupWithUserGroupWithPolicies []*model.GroupWithUserGroupWithPolicy

	if err := gr.Connection.Table(entity.UserGroupTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.GroupTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupGroupUuid.String(),
			entity.GroupTable.String(),
			entity.GroupUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PolicyTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupUuid.String(),
			entity.PolicyTable.String(),
			entity.PolicyUserGroupUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupUserUuid.String()), userUuid).
		Scan(&groupWithUserGroupWithPolicies).Error; err != nil {

		return nil, err
	}

	return groupWithUserGroupWithPolicies, nil
}

func (gr GroupRepositoryImpl) FindGroupWithPolicyByUserUuidAndGroupUuid(userUuid string, groupUuid string) (*model.GroupWithUserGroupWithPolicy, error) {
	var groupWithUserGroupWithPolicy model.GroupWithUserGroupWithPolicy

	if err := gr.Connection.Table(entity.UserGroupTable.String()).
		Select("*").
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.GroupTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupGroupUuid.String(),
			entity.GroupTable.String(),
			entity.GroupUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PolicyTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupUuid.String(),
			entity.PolicyTable.String(),
			entity.PolicyUserGroupUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupUserUuid.String()), userUuid).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupGroupUuid.String()), groupUuid).
		Scan(&groupWithUserGroupWithPolicy).Error; err != nil {

		return nil, err
	}

	return &groupWithUserGroupWithPolicy, nil

}

func (gr GroupRepositoryImpl) SaveWithRelationalData(
	group entity.Group,
	serviceGroup entity.ServiceGroup,
	userGroup entity.UserGroup,
	groupPermission entity.GroupPermission,
	groupRole entity.GroupRole,
	policy entity.Policy) (*entity.Group, error) {

	tx := gr.Connection.Begin()

	// Save groups
	if err := tx.Create(&group).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save service_groups
	serviceGroup.GroupUuid = group.Uuid
	if err := tx.Create(&serviceGroup).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save user_groups
	userGroup.GroupUuid = group.Uuid
	if err := tx.Create(&userGroup).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save group_roles
	groupRole.GroupUuid = group.Uuid
	if err := tx.Create(&groupRole).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save group_permissions
	groupPermission.GroupUuid = group.Uuid
	if err := tx.Create(&groupPermission).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save policies
	policy.UserGroupUuid = userGroup.Uuid
	if err := tx.Create(&policy).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &group, nil
}
