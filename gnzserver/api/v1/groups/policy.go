package groups

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/ctx"
	"github.com/tomoyane/grant-n-z/gnzserver/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var pInstance Policy

type Policy interface {
	// Implement permission api
	Api(w http.ResponseWriter, r *http.Request)

	// Http PUT method
	// Update user's policy
	put(w http.ResponseWriter, r *http.Request)
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
	log.Logger.Info("New `groups.Policy` instance")
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
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (p PolicyImpl) put(w http.ResponseWriter, r *http.Request) {
	var policyRequest *entity.PolicyRequest
	if err := middleware.BindBody(w, r, &policyRequest); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, policyRequest); err != nil {
		return
	}

	id, err := middleware.ParamGroupId(r)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	user, errUser := p.UserService.GetUserByEmail(policyRequest.ToUserEmail)
	if errUser != nil {
		model.WriteError(w, errUser.ToJson(), errUser.Code)
		return
	}

	userGroup, errGroup := p.UserService.GetUserGroupByUserIdAndGroupId(user.Id, id)
	if errGroup != nil {
		model.WriteError(w, errGroup.ToJson(), errGroup.Code)
		return
	}

	role, errRole := p.RoleService.GetRoleById(policyRequest.RoleId)
	if errRole != nil {
		model.WriteError(w, errRole.ToJson(), errRole.Code)
		return
	}

	permission, errPermission := p.PermissionService.GetPermissionById(policyRequest.PermissionId)
	if errPermission != nil {
		model.WriteError(w, errPermission.ToJson(), errPermission.Code)
		return
	}

	policy := entity.Policy{
		Name:         policyRequest.Name,
		RoleId:       role.Id,
		PermissionId: permission.Id,
		ServiceId:    ctx.GetServiceId().(int),
		UserGroupId:  userGroup.Id,
	}
	insertedPolicy, errPolicy := p.PolicyService.UpdatePolicy(policy)
	if errPolicy != nil {
		model.WriteError(w, errPolicy.ToJson(), errPolicy.Code)
		return
	}

	res, _ := json.Marshal(insertedPolicy)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
