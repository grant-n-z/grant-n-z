package users

import (
	"net/http"
	"testing"

	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	policy Policy
)

func init() {
	log.InitLogger("info")
	ctx.InitContext()

	policy = PolicyImpl{
		PolicyService: StubPolicyService{},
	}
}

// Test constructor
func TestGetPolicyInstance(t *testing.T) {
	GetPolicyInstance()
}

// Test method not allowed
func TestPolicy_MethodNotAllowed(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodPut}
	policy.Api(response, &request)

	if statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect TestPolicy_MethodNotAllowed test.")
		t.FailNow()
	}
}

// Test get
func TestPolicy_Get(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet}
	policy.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestPolicy_Get test.")
		t.FailNow()
	}
}

// Less than stub struct
// PolicyService
type StubPolicyService struct {
}

func (ps StubPolicyService) GetPolicies() ([]*entity.Policy, *model.ErrorResBody) {
	return []*entity.Policy{}, nil
}

func (ps StubPolicyService) GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody) {
	return []*entity.Policy{}, nil
}

func (ps StubPolicyService) GetPoliciesOfUser() ([]model.PolicyResponse, *model.ErrorResBody) {
	return []model.PolicyResponse{}, nil
}

func (ps StubPolicyService) GetPolicyByUserGroup(userId int, groupId int) (*entity.Policy, *model.ErrorResBody) {
	return &entity.Policy{}, nil
}

func (ps StubPolicyService) GetPolicyById(id int) (entity.Policy, *model.ErrorResBody) {
	return entity.Policy{}, nil
}

func (ps StubPolicyService) UpdatePolicy(policy entity.Policy) (*entity.Policy, *model.ErrorResBody) {
	return &entity.Policy{}, nil
}
