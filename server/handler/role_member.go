package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/service"
	"github.com/tomoyane/grant-n-z/server/log"
)

type RoleMemberHandler struct {
	RoleMemberService service.RoleMemberService
	UserService service.UserService
	RoleService service.RoleService
}

func NewRoleMemberHandler() RoleMemberHandler {
	log.Logger.Info("inject `RoleMemberService`,`UserService`,`RoleService` to `RoleMemberHandler`")
	return RoleMemberHandler{
		RoleMemberService: service.NewRoleMemberService(),
		UserService: service.NewUserService(),
		RoleService: service.NewRoleService(),
	}
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
	var result interface{}
	var roleMemberEntities []*entity.RoleMember
	var err *entity.ErrorResponse
	id := r.URL.Query().Get(entity.ROLE_MEMBER_USER_ID.String())

	if !strings.EqualFold(id, "") {
		log.Logger.Info("GET role_members by user_id")

		i, castErr := strconv.Atoi(id)
		if castErr != nil {
			badErr := entity.BadRequest(castErr.Error())
			http.Error(w, badErr.ToJson(), badErr.Code)
			return
		}

		roleMemberEntities, err = rmh.RoleMemberService.GetRoleMemberByUserId(i)
	} else {
		log.Logger.Info("GET role_members list")

		roleMemberEntities, err = rmh.RoleMemberService.GetRoleMembers()
	}

	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	if roleMemberEntities == nil {
		result = []string{}
	} else {
		result = roleMemberEntities
	}

	res, _ := json.Marshal(result)
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

	if userEntity, _ := rmh.UserService.GetUserById(roleMemberEntity.UserId); userEntity == nil {
		err = entity.BadRequest("not found user id")
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	if roleEntity, _ := rmh.RoleService.GetRoleById(roleMemberEntity.RoleId); roleEntity == nil {
		err = entity.BadRequest("not found role id")
		http.Error(w, err.ToJson(), err.Code)
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
