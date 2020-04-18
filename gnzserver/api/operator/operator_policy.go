package operator

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var omhInstance OperatorPolicy

type OperatorPolicy interface {
	// Implement operator policy api
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	post(w http.ResponseWriter, r *http.Request)

	// Http PUT method
	put(w http.ResponseWriter, r *http.Request)

	// Http DELETE method
	delete(w http.ResponseWriter, r *http.Request)
}

type OperatorPolicyImpl struct {
	OperatorPolicyService service.OperatorPolicyService
}

func GetOperatorPolicyInstance() OperatorPolicy {
	if omhInstance == nil {
		omhInstance = NewOperatorPolicy()
	}
	return omhInstance
}

func NewOperatorPolicy() OperatorPolicy {
	log.Logger.Info("New `OperatorPolicy` instance")
	return OperatorPolicyImpl{OperatorPolicyService: service.NewOperatorPolicyServiceService()}
}

func (rmrhi OperatorPolicyImpl) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rmrhi.get(w, r)
	case http.MethodPost:
		rmrhi.post(w, r)
	case http.MethodPut:
		rmrhi.put(w, r)
	case http.MethodDelete:
		rmrhi.delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (rmrhi OperatorPolicyImpl) get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(entity.OperatorPolicyUserId.String())

	roleMemberEntities, err := rmrhi.OperatorPolicyService.Get(id)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleMemberEntities)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (rmrhi OperatorPolicyImpl) post(w http.ResponseWriter, r *http.Request) {
	var roleMemberEntity *entity.OperatorPolicy
	if err := middleware.BindBody(w, r, &roleMemberEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, roleMemberEntity); err != nil {
		return
	}

	roleMember, err := rmrhi.OperatorPolicyService.Insert(roleMemberEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleMember)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (rmrhi OperatorPolicyImpl) put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (rmrhi OperatorPolicyImpl) delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
