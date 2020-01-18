package data

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var ugrInstance UserGroupRepository

type UserGroupRepository interface {
	// Get all groups that has user
	FindGroupsByUserId(userId int) ([]*entity.Group, *model.ErrorResBody)

	// Get all groups with user_groups with policy that has user
	FindGroupWithUserWithPolicyGroupsByUserId(userId int) ([]*entity.GroupWithUserGroupWithPolicy, *model.ErrorResBody)

	// Get user_groups with policies by user id and group id
	FindGroupWithPolicyByUserIdAndGroupId(userId int, groupId int) (*entity.GroupWithUserGroupWithPolicy, *model.ErrorResBody)
}

type UserGroupRepositoryImpl struct {
	Db *gorm.DB
}

func GetUserGroupRepositoryInstance(db *gorm.DB) UserGroupRepository {
	if ugrInstance == nil {
		ugrInstance = NewUserGroupRepository(db)
	}
	return ugrInstance
}

func NewUserGroupRepository(db *gorm.DB) UserGroupRepository {
	log.Logger.Info("New `UserGroupRepository` instance")
	return UserGroupRepositoryImpl{Db: db}
}

func (ugr UserGroupRepositoryImpl) FindGroupsByUserId(userId int) ([]*entity.Group, *model.ErrorResBody) {
	var groups []*entity.Group

	if err := ugr.Db.Table(entity.UserGroupTable.String()).
		Select("*").
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
			return nil, model.InternalServerError()
	}

	return groups, nil
}

func (ugr UserGroupRepositoryImpl) FindGroupWithUserWithPolicyGroupsByUserId(userId int) ([]*entity.GroupWithUserGroupWithPolicy, *model.ErrorResBody) {
	var groupWithUserGroupWithPolicies []*entity.GroupWithUserGroupWithPolicy

	if err := ugr.Db.Table(entity.UserGroupTable.String()).
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
			return nil, model.InternalServerError()
	}

	return groupWithUserGroupWithPolicies, nil
}

func (ugr UserGroupRepositoryImpl) FindGroupWithPolicyByUserIdAndGroupId(userId int, groupId int) (*entity.GroupWithUserGroupWithPolicy, *model.ErrorResBody) {
	var groupWithUserGroupWithPolicy *entity.GroupWithUserGroupWithPolicy

	if err := ugr.Db.Table(entity.UserGroupTable.String()).
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
			return nil, model.InternalServerError()
	}

	return groupWithUserGroupWithPolicy, nil

}
