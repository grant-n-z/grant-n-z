package groups

import (
	"bytes"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	policy Policy
)

func init() {
	log.InitLogger("info")
	ctx.InitContext()

	policy = PolicyImpl{
		PolicyService: StubPolicyService{},
		UserService: StubUserService{},
		RoleService: StubRoleService{},
		PermissionService: StubPermissionService{},
	}
}

// Test constructor
func TestGetPolicyInstance(t *testing.T) {
	GetPolicyInstance()
}

// Test method not allowed
func TestPolicy_MethodNotAllowed(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet}
	policy.Api(response, &request)

	if statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect TestPolicy_MethodNotAllowed test.")
		t.FailNow()
	}
}

// Test put bad request
func TestPolicy_Put_BadRequest_Body(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"name\":\"test\",\"to_user_email\":\"test@gmail.com\",\"role_id\":0,\"permission_id\":\"\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPut, Body: invalid}
	policy.Api(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestPolicy_Put_BadRequest_Body test.")
		t.FailNow()
	}
}

// Test put bad request
func TestPolicy_Put_BadRequest_QueryParam(t *testing.T) {
	response := StubResponseWriter{}
	body := ioutil.NopCloser(bytes.NewReader([]byte("{\"name\":\"test\",\"to_user_email\":\"test@gmail.com\",\"role_id\":10,\"permission_id\":10}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPut, Body: body}
	policy.Api(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestPolicy_Put_BadRequest_QueryParam test.")
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
