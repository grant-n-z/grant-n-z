package service

import (
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
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
	ctx.InitContext()
	ctx.SetUserId(1)
	ctx.SetServiceId(1)

	stubConnection, _ = gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	policyService = PolicyServiceImpl{
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
			Ctx:        ctx.GetCtx(),
		},
		PolicyRepository:     StubPolicyRepositoryImpl{Connection: stubConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
		RoleRepository:       StubRoleRepositoryImpl{Connection: stubConnection},
		ServiceRepository:    StubServiceRepositoryImpl{Connection: stubConnection},
		GroupRepository:      StubGroupRepositoryImpl{Connection: stubConnection},
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
	_, err := policyService.GetPoliciesByRoleId(1)
	if err != nil {
		t.Errorf("Incorrect TestGetPoliciesByRoleId_Success test")
		t.FailNow()
	}
}

// Test get policies of request user
func TestGetPoliciesOfUser_Success(t *testing.T) {
	_, err := policyService.GetPoliciesOfUser()
	if err != nil {
		t.Errorf("Incorrect TestGetPoliciesOfUser_Success test")
		t.FailNow()
	}
}

// Test get policies by user_group
func TestGetPolicyByUserGroup_Success(t *testing.T) {
	_, err := policyService.GetPolicyByUserGroup(1, 1)
	if err != nil {
		t.Errorf("Incorrect TestGetPolicyByUserGroup_Success test")
		t.FailNow()
	}
}

// Test get policies by id
func TestGetPolicyById_Success(t *testing.T) {
	_, err := policyService.GetPolicyById(1)
	if err != nil {
		t.Errorf("Incorrect TestGetPolicyById_Success test")
		t.FailNow()
	}
}

// Test update
func TestUpdatePolicy_Success(t *testing.T) {
	_, err := policyService.UpdatePolicy(entity.Policy{})
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

func (pri StubPolicyRepositoryImpl) FindAll() ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindById(id int) (entity.Policy, *model.ErrorResBody) {
	var policy entity.Policy
	return policy, nil
}

func (pri StubPolicyRepositoryImpl) Update(policy entity.Policy) (*entity.Policy, *model.ErrorResBody) {
	return &policy, nil
}
