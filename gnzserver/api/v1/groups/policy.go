package groups

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var pInstance Policy

type Policy interface {
	// Implement permission api
	// Endpoint is `/api/v1/groups/{id}/policy`
	Api(w http.ResponseWriter, r *http.Request)

	// Http PUT method
	// Update user's policy
	put(w http.ResponseWriter, r *http.Request)

	// Http GET method
	// Update user's policy
	get(w http.ResponseWriter, r *http.Request)
}

type PolicyImpl struct {
	PolicyService     service.PolicyService
	UserService       service.UserService
	RoleService       service.RoleService
	PermissionService service.PermissionService
}

func GetPolicyInstance() Policy {
	if pInstance == nil {
		pInstance = NewPolicy()
	}
	return pInstance
}

func NewPolicy() Policy {
	return PolicyImpl{
		PolicyService:     service.GetPolicyServiceInstance(),
		UserService:       service.GetUserServiceInstance(),
		RoleService:       service.GetRoleServiceInstance(),
		PermissionService: service.GetPermissionServiceInstance(),
	}
}

func (p PolicyImpl) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		p.put(w, r)
	case http.MethodGet:
		p.get(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (p PolicyImpl) put(w http.ResponseWriter, r *http.Request) {
	var policyRequest *model.PolicyRequest
	if err := middleware.BindBody(w, r, &policyRequest); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, policyRequest); err != nil {
		return
	}

	secret := r.Context().Value(middleware.ScopeSecret).(string)
	insertedPolicy, errPolicy := p.PolicyService.UpdatePolicy(*policyRequest, secret, middleware.ParamGroupUuid(r))
	if errPolicy != nil {
		model.WriteError(w, errPolicy.ToJson(), errPolicy.Code)
		return
	}

	res, _ := json.Marshal(insertedPolicy)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (p PolicyImpl) get(w http.ResponseWriter, r *http.Request) {
	policies, err := p.PolicyService.GetPoliciesByUserGroup(middleware.ParamGroupUuid(r))
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(policies)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
