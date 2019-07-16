package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"

	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

type OperatorMemberRoleHandlerImpl struct {
	RoleMemberService service.OperatorMemberRoleService
}

func NewOperatorMemberRoleHandler() OperateMemberRoleHandler {
	log.Logger.Info("Inject `OperatorMemberRoleService` to `OperateMemberRoleHandler`")
	return OperatorMemberRoleHandlerImpl{RoleMemberService: service.NewOperatorMemberRoleService()}
}

func (rmrhi OperatorMemberRoleHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet: rmrhi.Get(w, r)
	case http.MethodPost: rmrhi.Post(w, r)
	case http.MethodPut: rmrhi.Put(w, r)
	case http.MethodDelete: rmrhi.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (rmrhi OperatorMemberRoleHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET role_members")
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
	log.Logger.Info("POST role_members")
	var roleMemberEntity *entity.OperatorMemberRole

	body, err := Interceptor(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &roleMemberEntity)
	if err := ValidateHttpRequest(w, roleMemberEntity); err != nil {
		return
	}

	roleMemberEntity, err = rmrhi.RoleMemberService.InsertRoleMember(roleMemberEntity)
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
