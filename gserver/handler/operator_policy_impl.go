package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var omhInstance OperatePolicyHandler

type OperatorPolicyHandlerImpl struct {
	RequestHandler        RequestHandler
	OperatorPolicyService service.OperatorPolicyService
}

func GetOperatorPolicyHandlerInstance() OperatePolicyHandler {
	if omhInstance == nil {
		omhInstance = NewOperatorPolicyHandler()
	}
	return omhInstance
}

func NewOperatorPolicyHandler() OperatePolicyHandler {
	log.Logger.Info("New `OperatePolicyHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `operatorMemberRoleService` to `OperatePolicyHandler`")
	return OperatorPolicyHandlerImpl{
		RequestHandler:        GetRequestHandlerInstance(),
		OperatorPolicyService: service.NewOperatorPolicyServiceService(),
	}
}

func (rmrhi OperatorPolicyHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := rmrhi.RequestHandler.VerifyToken(w, r, property.AuthOperator)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		rmrhi.Get(w, r)
	case http.MethodPost:
		rmrhi.Post(w, r)
	case http.MethodPut:
		rmrhi.Put(w, r)
	case http.MethodDelete:
		rmrhi.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (rmrhi OperatorPolicyHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(entity.OperatorPolicyUserId.String())

	roleMemberEntities, err := rmrhi.OperatorPolicyService.Get(id)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleMemberEntities)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (rmrhi OperatorPolicyHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	var roleMemberEntity *entity.OperatorPolicy

	body, err := rmrhi.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &roleMemberEntity)
	if err := rmrhi.RequestHandler.ValidateHttpRequest(w, roleMemberEntity); err != nil {
		return
	}

	roleMemberEntity, err = rmrhi.OperatorPolicyService.Insert(roleMemberEntity)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleMemberEntity)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (rmrhi OperatorPolicyHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (rmrhi OperatorPolicyHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
