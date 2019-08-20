package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

type OperatorMemberRoleHandlerImpl struct {
	RequestHandler    RequestHandler
	RoleMemberService service.OperatorMemberRoleService
}

func NewOperatorMemberRoleHandler() OperateMemberRoleHandler {
	log.Logger.Info("Inject `RequestHandler`, `OperatorMemberRoleService` to `OperateMemberRoleHandler`")
	return OperatorMemberRoleHandlerImpl{
		RequestHandler: NewRequestHandler(),
		RoleMemberService: service.NewOperatorMemberRoleService(),
	}
}

func (rmrhi OperatorMemberRoleHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (rmrhi OperatorMemberRoleHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET operator_member_roles")
	id := r.URL.Query().Get(entity.OperatorMemberRoleUserId.String())

	roleMemberEntities, err := rmrhi.RoleMemberService.Get(id)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleMemberEntities)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (rmrhi OperatorMemberRoleHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST operator_member_roles")
	var roleMemberEntity *entity.OperatorMemberRole

	body, err := rmrhi.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &roleMemberEntity)
	if err := rmrhi.RequestHandler.ValidateHttpRequest(w, roleMemberEntity); err != nil {
		return
	}

	roleMemberEntity, err = rmrhi.RoleMemberService.Insert(roleMemberEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleMemberEntity)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (rmrhi OperatorMemberRoleHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (rmrhi OperatorMemberRoleHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
