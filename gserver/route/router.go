package route

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api/v1"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type Router struct {
	Auth           v1.Auth
	Token          v1.Token
	Group          v1.Group
	User           v1.User
	UserGroup      v1.UserGroup
	Service        v1.Service
	ServiceGroup   v1.ServiceGroup
	Role           v1.Role
	OperatorPolicy v1.OperatorPolicy
	UserService    v1.UserService
	Permission     v1.Permission
	Policy         v1.Policy
}

func NewRouter() Router {
	return Router{
		Auth:           v1.GetAuthInstance(),
		Token:          v1.GetTokenInstance(),
		Group:          v1.GetGroupInstance(),
		User:           v1.GetUserInstance(),
		UserGroup:      v1.GetUserGroupInstance(),
		Service:        v1.GetServiceInstance(),
		ServiceGroup:   v1.GetServiceGroupInstance(),
		Role:           v1.GetRoleInstance(),
		OperatorPolicy: v1.GetOperatorPolicyInstance(),
		UserService:    v1.GetUserServiceInstance(),
		Permission:     v1.GetPermissionInstance(),
		Policy:         v1.GetPolicyInstance(),
	}
}

func (r Router) Init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := model.NotFound("Not found resource path.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res.ToJson()))
	})
}

func (r Router) V1() {
	http.HandleFunc("/api/v1/auth", r.Auth.Api)
	http.HandleFunc("/api/v1/token", r.Token.Api)

	http.HandleFunc("/api/v1/groups", r.Group.Api)

	http.HandleFunc("/api/v1/users", r.User.Api)
	http.HandleFunc("/api/v1/user_services", r.UserService.Api)
	http.HandleFunc("/api/v1/user_groups", r.UserGroup.Api)

	http.HandleFunc("/api/v1/services", r.Service.Api)
	http.HandleFunc("/api/v1/service_groups", r.ServiceGroup.Api)

	http.HandleFunc("/api/v1/roles", r.Role.Api)
	http.HandleFunc("/api/v1/permissions", r.Permission.Api)

	http.HandleFunc("/api/v1/policies", r.Policy.Api)
	http.HandleFunc("/api/v1/operator_policies", r.OperatorPolicy.Api)

	log.Logger.Info("------ Routing info ------")
	log.Logger.Info("HTTP Method: `POST` Routing: /api/v1/oauth")
	log.Logger.Info("HTTP Method: `POST`, `PUT` Routing: /api/v1/groups")
	log.Logger.Info("HTTP Method: `POST`, `PUT` Routing: /api/v1/users")
	log.Logger.Info("HTTP Method: `POST`, `PUT` Routing: /api/v1/user_service")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/user_groups`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/services`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/service_groups`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/roles`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/permissions`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/policies`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/operator_policies`")
	log.Logger.Info("------ Routing info ------")
}
