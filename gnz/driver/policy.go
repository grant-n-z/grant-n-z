package driver

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var plrInstance PolicyRepository

type PolicyRepository interface {
	// Find all policy
	FindAll() ([]*entity.Policy, *model.ErrorResBody)

	// Find policy for offset and limit
	FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Policy, *model.ErrorResBody)

	// Find policy by role id
	FindByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody)

	// Find by id
	FindById(id int) (entity.Policy, *model.ErrorResBody)

	// Find policy data by user id
	FindPolicyResponseOfUserByUserIdAndGroupId(userId int, groupId int) (model.UserPolicyOnGroupResponse, *model.ErrorResBody)

	// Update
	Update(policy entity.Policy) (*entity.Policy, *model.ErrorResBody)
}

type PolicyRepositoryImpl struct {
	Connection *gorm.DB
}

func GetPolicyRepositoryInstance() PolicyRepository {
	if plrInstance == nil {
		plrInstance = NewPolicyRepository()
	}
	return plrInstance
}

func NewPolicyRepository() PolicyRepository {
	log.Logger.Info("New `PolicyRepository` instance")
	return PolicyRepositoryImpl{Connection: connection}
}

func (pri PolicyRepositoryImpl) FindAll() ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	if err := pri.Connection.Find(&policies).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return policies, nil
}

func (pri PolicyRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	if err := pri.Connection.Limit(limitCnt).Offset(offsetCnt).Find(&policies).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return policies, nil
}

func (pri PolicyRepositoryImpl) FindByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	if err := pri.Connection.Where("role_id = ?", roleId).Find(&policies).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return policies, nil
}

func (pri PolicyRepositoryImpl) FindById(id int) (entity.Policy, *model.ErrorResBody) {
	var policy entity.Policy
	if err := pri.Connection.Where("id = ?", id).Find(&policy).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return policy, nil
		}

		return policy, model.InternalServerError(err.Error())
	}

	return policy, nil
}

func (pri PolicyRepositoryImpl) FindPolicyResponseOfUserByUserIdAndGroupId(userId int, groupId int) (model.UserPolicyOnGroupResponse, *model.ErrorResBody) {
	var policy model.UserPolicyOnGroupResponse
	target := entity.UserTable.String() + "." + entity.UserUsername.String() + "," + entity.UserTable.String() + "." + entity.UserEmail.String() +
		"," + entity.PolicyTable.String() + "." + entity.PolicyName.String() + " AS policy_name," + entity.RoleTable.String() + "." + entity.RoleName.String() +
		" AS role_name," + entity.PermissionTable.String() + "." + entity.PermissionName.String() + " AS permission_name"

	if err := pri.Connection.Table(entity.UserGroupTable.String()).
		Select(target).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PolicyTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupId.String(),
			entity.PolicyTable.String(),
			entity.PolicyUserGroupId.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.RoleTable.String(),
			entity.RoleTable.String(),
			entity.RoleId.String(),
			entity.PolicyTable.String(),
			entity.PolicyRoleId.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PermissionTable.String(),
			entity.PermissionTable.String(),
			entity.PermissionId.String(),
			entity.PolicyTable.String(),
			entity.PolicyPermissionId.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupUserId.String(),
			entity.UserTable.String(),
			entity.UserId.String())).
		Where(fmt.Sprintf("%s.%s = ? AND %s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupUserId.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupGroupId.String()), userId, groupId).
		Scan(&policy).Error; err != nil {

		log.Logger.Warn(err.Error())
		return policy, model.InternalServerError()
	}

	return policy, nil
}

func (pri PolicyRepositoryImpl) Update(policy entity.Policy) (*entity.Policy, *model.ErrorResBody) {
	if err := pri.Connection.Where("user_group_id = ?", policy.UserGroupId).Assign(policy).FirstOrCreate(&policy).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		}

		return nil, model.InternalServerError()
	}

	return &policy, nil
}
