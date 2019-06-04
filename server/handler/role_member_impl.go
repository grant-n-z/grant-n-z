package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"

	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

type RoleMemberHandlerImpl struct {
	RoleMemberService service.RoleMemberService
}

func NewRoleMemberHandler() RoleMemberHandler {
	log.Logger.Info("Inject `RoleMemberService` to `RoleMemberHandler`")
	return RoleMemberHandlerImpl{RoleMemberService: service.NewRoleMemberService()}
}

func (rmh RoleMemberHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet: rmh.Get(w, r)
	case http.MethodPost: rmh.Post(w, r)
	case http.MethodPut: rmh.Put(w, r)
	case http.MethodDelete: rmh.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (rmh RoleMemberHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
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

func (rmh RoleMemberHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
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

func (rmh RoleMemberHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (rmh RoleMemberHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
