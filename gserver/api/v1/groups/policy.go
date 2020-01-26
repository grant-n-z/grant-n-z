package groups

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
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
	// TODO
}

func GetPolicyInstance() Policy {
	if pInstance == nil {
		pInstance = NewPolicy()
	}
	return pInstance
}

func NewPolicy() Policy {
	log.Logger.Info("New `groups.Policy` instance")
	return PolicyImpl{}
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
	//res, _ := json.Marshal(permissionEntities)
	//w.WriteHeader(http.StatusOK)
	//w.Write(res)
}
