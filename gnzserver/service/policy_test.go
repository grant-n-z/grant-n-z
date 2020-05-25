package service

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

var (
	policyService PolicyService
)

// Set up
func init() {
	log.InitLogger("info")

	stubConnection, _ = gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	policyService = PolicyServiceImpl{
		EtcdClient:           cache.EtcdClientImpl{Connection: stubEtcdConnection},
		PolicyRepository:     StubPolicyRepositoryImpl{Connection: stubConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
		RoleRepository:       StubRoleRepositoryImpl{Connection: stubConnection},
		ServiceRepository:    StubServiceRepositoryImpl{Connection: stubConnection},
		GroupRepository:      StubGroupRepositoryImpl{Connection: stubConnection},
		UserRepository:       StubUserRepositoryImpl{Connection: stubConnection},
	}
}

// Test constructor
func TestGetPolicyServiceInstance(t *testing.T) {
	GetPolicyServiceInstance()
}

// Test get policies
func TestGetPolicies_Success(t *testing.T) {
	_, err := policyService.GetPolicies()
	if err != nil {
		t.Errorf("Incorrect TestGetPolicies_Success test")
		t.FailNow()
	}
}

// Test get policies by role id
func TestGetPoliciesByRoleId_Success(t *testing.T) {
	_, err := policyService.GetPoliciesByRoleUuid(uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetPoliciesByRoleId_Success test")
		t.FailNow()
	}
}

// Test get policies of request user
func TestGetPoliciesOfUser_Success(t *testing.T) {
	_, err := policyService.GetPoliciesByUser(uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetPoliciesOfUser_Success test")
		t.FailNow()
	}
}

// Test get policy by user_group
func TestGetPolicyByUserGroup_Success(t *testing.T) {
	_, err := policyService.GetPolicyByUserGroup(uuid.New().String(), uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetPolicyByUserGroup_Success test")
		t.FailNow()
	}
}

// Test get policies by user_group
func TestGetPoliciesOfUserGroup_Success(t *testing.T) {
	_, err := policyService.GetPoliciesByUserGroup(uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetPoliciesOfUserGroup_Success test")
		t.FailNow()
	}
}

// Test get policies by id
func TestGetPolicyById_Success(t *testing.T) {
	_, err := policyService.GetPolicyByUuid(uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetPolicyById_Success test")
		t.FailNow()
	}
}

// Test update
func TestUpdatePolicy_Success(t *testing.T) {
	_, err := policyService.UpdatePolicy(model.PolicyRequest{}, "", "")
	if err != nil {
		t.Errorf("Incorrect TestUpdatePolicy_Success test")
		t.FailNow()
	}
}

// Less than stub struct
// Policy repository
type StubPolicyRepositoryImpl struct {
	Connection *gorm.DB
}

func (pri StubPolicyRepositoryImpl) FindAll() ([]*entity.Policy, error) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Policy, error) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindByRoleUuid(roleUuid string) ([]*entity.Policy, error) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindByUuid(uuid string) (entity.Policy, error) {
	var policy entity.Policy
	return policy, nil
}

func (pri StubPolicyRepositoryImpl) FindPolicyOfUserGroupByUserUuidAndGroupUuid(userUuid string, groupUuid string) (model.UserPolicyOnGroupResponse, error) {
	var policy model.UserPolicyOnGroupResponse
	return policy, nil
}

func (pri StubPolicyRepositoryImpl) FindPolicyOfUserServiceByUserUuidAndServiceUuid(userUuid string) ([]model.UserPolicyOnServiceResponse, error) {
	var policy []model.UserPolicyOnServiceResponse
	return policy, nil
}

func (pri StubPolicyRepositoryImpl) Update(policy entity.Policy) (*entity.Policy, error) {
	return &policy, nil
}
