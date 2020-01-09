package users

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var plhInstance Policy

type Policy interface {
	// Implement policy api
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
	log.Logger.Info("New `Policy` instance")
	log.Logger.Info("Inject `PolicyService` to `Policy`")
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
	policyResponses, err := ph.PolicyService.GetPoliciesOfUser()
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(policyResponses)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
