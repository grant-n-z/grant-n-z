package users

import (
	"context"
	"github.com/google/uuid"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"net/http"
	"testing"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	policy Policy
)

func init() {
	log.InitLogger("info")

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

	jwt := model.JwtPayload{
		UserUuid: uuid.New(),
		Username: "user",
	}
	policy.Api(response, request.WithContext(context.WithValue(request.Context(), middleware.ScopeJwt, jwt)))

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

func (ps StubPolicyService) GetPoliciesByRoleUuid(roleUuid string) ([]*entity.Policy, *model.ErrorResBody) {
	return []*entity.Policy{}, nil
}

func (ps StubPolicyService) GetPoliciesByUser(userUuid string) ([]model.PolicyResponse, *model.ErrorResBody) {
	return []model.PolicyResponse{}, nil
}

func (ps StubPolicyService) GetPolicyByUserGroup(userUuid string, groupUuid string) (*entity.Policy, *model.ErrorResBody) {
	return &entity.Policy{}, nil
}

func (ps StubPolicyService) GetPoliciesByUserGroup(groupUuid string) ([]model.UserPolicyOnGroupResponse, *model.ErrorResBody) {
	return []model.UserPolicyOnGroupResponse{}, nil
}

func (ps StubPolicyService) GetPolicyByUuid(uuid string) (entity.Policy, *model.ErrorResBody) {
	return entity.Policy{}, nil
}

func (ps StubPolicyService) UpdatePolicy(policyRequest model.PolicyRequest, secret string, groupUuid string) (*entity.Policy, *model.ErrorResBody) {
	return &entity.Policy{}, nil
}
