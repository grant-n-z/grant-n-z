package users

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var plhInstance Policy

type Policy interface {
	// Implement policy api
	// Endpoint is `/api/v1/users/policy`
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	// Required query param what it is group name
	get(w http.ResponseWriter, r *http.Request)
}

// Policy api struct
type PolicyImpl struct {
	PolicyService service.PolicyService
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetPolicyInstance() Policy {
	if plhInstance == nil {
		plhInstance = NewPolicy()
	}
	return plhInstance
}

// Constructor
func NewPolicy() Policy {
	return PolicyImpl{PolicyService: service.GetPolicyServiceInstance()}
}

func (ph PolicyImpl) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ph.get(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (ph PolicyImpl) get(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value(middleware.ScopeJwt).(model.JwtPayload)
	policyResponses, err := ph.PolicyService.GetPoliciesByUser(jwt.UserUuid)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(policyResponses)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
