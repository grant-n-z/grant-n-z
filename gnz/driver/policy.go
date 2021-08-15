package driver

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var plrInstance PolicyRepository

type PolicyRepository interface {
	// FindAll
	// Find all policy
	FindAll() ([]*entity.Policy, error)

	// FindOffSetAndLimit
	// Find policy for offset and limit
	FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Policy, error)

	// FindByRoleUuid
	// Find policy by role uuid
	FindByRoleUuid(roleUuid string) ([]*entity.Policy, error)

	// FindByUuid
	// Find by uuid
	FindByUuid(uuid string) (entity.Policy, error)

	// FindPolicyOfUserGroupByUserUuidAndGroupUuid
	// Find policy data by user uuid and group uuid
	FindPolicyOfUserGroupByUserUuidAndGroupUuid(userUuid string, groupUuid string) (model.UserPolicyOnGroupResponse, error)

	// FindPolicyOfUserServiceByUserUuidAndServiceUuid
	// Find policy data by user uuid and group uuid
	FindPolicyOfUserServiceByUserUuidAndServiceUuid(userUuid string) ([]model.UserPolicyOnServiceResponse, error)

	Update(policy entity.Policy) (*entity.Policy, error)
}

type RdbmsPolicyRepository struct {
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
	return RdbmsPolicyRepository{Connection: connection}
}

func (pri RdbmsPolicyRepository) FindAll() ([]*entity.Policy, error) {
	var policies []*entity.Policy
	if err := pri.Connection.Find(&policies).Error; err != nil {
		return nil, err
	}

	return policies, nil
}

func (pri RdbmsPolicyRepository) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Policy, error) {
	var policies []*entity.Policy
	if err := pri.Connection.Limit(limitCnt).Offset(offsetCnt).Find(&policies).Error; err != nil {
		return nil, err
	}

	return policies, nil
}

func (pri RdbmsPolicyRepository) FindByRoleUuid(roleUuid string) ([]*entity.Policy, error) {
	var policies []*entity.Policy
	if err := pri.Connection.Where("role_uuid = ?", roleUuid).Find(&policies).Error; err != nil {
		return nil, err
	}

	return policies, nil
}

func (pri RdbmsPolicyRepository) FindByUuid(uuid string) (entity.Policy, error) {
	var policy entity.Policy
	if err := pri.Connection.Where("uuid = ?", uuid).Find(&policy).Error; err != nil {
		return entity.Policy{}, err
	}

	return policy, nil
}

func (pri RdbmsPolicyRepository) FindPolicyOfUserGroupByUserUuidAndGroupUuid(userUuid string, groupUuid string) (model.UserPolicyOnGroupResponse, error) {
	var policy model.UserPolicyOnGroupResponse

	target := entity.UserTable.String() + "." +
		entity.UserUsername.String() + "," +
		entity.UserTable.String() + "." +
		entity.UserEmail.String() + "," +
		entity.PolicyTable.String() + "." +
		entity.PolicyName.String() + " AS policy_name," +
		entity.RoleTable.String() + "." +
		entity.RoleName.String() + " AS role_name," +
		entity.PermissionTable.String() + "." +
		entity.PermissionName.String() + " AS permission_name," +
		entity.ServiceTable.String() + "." +
		entity.ServiceName.String() + " AS service_name"

	if err := pri.Connection.Table(entity.UserGroupTable.String()).
		Select(target).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PolicyTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupUuid.String(),
			entity.PolicyTable.String(),
			entity.PolicyUserGroupUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.RoleTable.String(),
			entity.RoleTable.String(),
			entity.RoleUuid.String(),
			entity.PolicyTable.String(),
			entity.PolicyRoleUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PermissionTable.String(),
			entity.PermissionTable.String(),
			entity.PermissionUuid.String(),
			entity.PolicyTable.String(),
			entity.PolicyPermissionUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupUserUuid.String(),
			entity.UserTable.String(),
			entity.UserUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PolicyTable.String(),
			entity.PolicyTable.String(),
			entity.PolicyServiceUuid.String(),
			entity.ServiceTable.String(),
			entity.ServiceUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupUserUuid.String()), userUuid).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserGroupTable.String(),
			entity.UserGroupGroupUuid.String()), groupUuid).
		Scan(&policy).Error; err != nil {

		return model.UserPolicyOnGroupResponse{}, err
	}

	return policy, nil
}

func (pri RdbmsPolicyRepository) FindPolicyOfUserServiceByUserUuidAndServiceUuid(userUuid string) ([]model.UserPolicyOnServiceResponse, error) {
	var policy []model.UserPolicyOnServiceResponse

	target := entity.UserTable.String() + "." +
		entity.UserUsername.String() + "," +
		entity.UserTable.String() + "." +
		entity.UserEmail.String() + "," +
		entity.PolicyTable.String() + "." +
		entity.PolicyName.String() + " AS policy_name," +
		entity.RoleTable.String() + "." +
		entity.RoleName.String() + " AS role_name," +
		entity.RoleTable.String() + "." +
		entity.RoleUuid.String() + " AS role_uuid," +
		entity.PermissionTable.String() + "." +
		entity.PermissionName.String() + " AS permission_name," +
		entity.PermissionTable.String() + "." +
		entity.PermissionUuid.String() + " AS permission_uuid," +
		entity.GroupTable.String() + "." +
		entity.GroupName.String() + " AS group_name," +
		entity.GroupTable.String() + "." +
		entity.GroupUuid.String() + " AS group_uuid"

	if err := pri.Connection.Table(entity.UserTable.String()).
		Select(target).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserServiceTable.String(),
			entity.UserServiceTable.String(),
			entity.UserServiceUserUuid.String(),
			entity.UserTable.String(),
			entity.UserUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PolicyTable.String(),
			entity.PolicyTable.String(),
			entity.PolicyServiceUuid.String(),
			entity.UserServiceTable.String(),
			entity.UserServiceServiceUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.PermissionTable.String(),
			entity.PermissionTable.String(),
			entity.PermissionUuid.String(),
			entity.PolicyTable.String(),
			entity.PolicyPermissionUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.RoleTable.String(),
			entity.RoleTable.String(),
			entity.RoleUuid.String(),
			entity.PolicyTable.String(),
			entity.PolicyRoleUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.UserGroupTable.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupUuid.String(),
			entity.PolicyTable.String(),
			entity.PolicyUserGroupUuid.String())).
		Joins(fmt.Sprintf("LEFT JOIN %s ON %s.%s = %s.%s",
			entity.GroupTable.String(),
			entity.GroupTable.String(),
			entity.GroupUuid.String(),
			entity.UserGroupTable.String(),
			entity.UserGroupGroupUuid.String())).
		Where(fmt.Sprintf("%s.%s = ?",
			entity.UserTable.String(),
			entity.UserUuid.String()), userUuid).
		Scan(&policy).Error; err != nil {

		return []model.UserPolicyOnServiceResponse{}, err
	}

	return policy, nil
}

func (pri RdbmsPolicyRepository) Update(policy entity.Policy) (*entity.Policy, error) {
	if err := pri.Connection.Where("user_group_uuid = ?", policy.UserGroupUuid).Assign(policy).FirstOrCreate(&policy).Error; err != nil {
		return nil, err
	}

	return &policy, nil
}
