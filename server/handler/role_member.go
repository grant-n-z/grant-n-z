package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/service"
	"github.com/tomoyane/grant-n-z/server/log"
)

type RoleMemberHandler struct {
	RoleMemberService service.RoleMemberService
}

func NewRoleMemberHandler() RoleMemberHandler {
	log.Logger.Info("inject `RoleMemberService` to `RoleMemberHandler`")
	return RoleMemberHandler{RoleMemberService: service.NewRoleMemberService()}
}

func (rmh RoleMemberHandler) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet: rmh.Get(w, r)
	case http.MethodPost: rmh.Post(w, r)
	case http.MethodPut: rmh.Put(w, r)
	case http.MethodDelete: rmh.Delete(w, r)
	default:
		err := entity.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (rmh RoleMemberHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET role_members")
	id := r.URL.Query().Get(entity.ROLE_MEMBER_USER_ID.String())

	roleMemberEntities, err := rmh.RoleMemberService.Get(id)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleMemberEntities)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (rmh RoleMemberHandler) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST role_members")
	var roleMemberEntity *entity.RoleMember

	body, err := Interceptor(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &roleMemberEntity)
	if err := BodyValidator(w, roleMemberEntity); err != nil {
		return
	}

	roleMemberEntity, err = rmh.RoleMemberService.InsertRoleMember(roleMemberEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleMemberEntity)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (rmh RoleMemberHandler) Put(w http.ResponseWriter, r *http.Request) {
}

func (rmh RoleMemberHandler) Delete(w http.ResponseWriter, r *http.Request) {
}
