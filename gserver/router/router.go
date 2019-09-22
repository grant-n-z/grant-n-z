package router

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/handler"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type Router struct {
	AuthHandler          handler.AuthHandler
	TokenHandler         handler.TokenHandler
	GroupHandler         handler.GroupHandler
	UserHandler          handler.UserHandler
	UserGroupHandler     handler.UserGroupHandler
	ServiceHandler       handler.ServiceHandler
	ServiceGroupHandler  handler.ServiceGroupHandler
	RoleHandler          handler.RoleHandler
	OperatePolicyHandler handler.OperatePolicyHandler
	UserServiceHandler   handler.UserServiceHandler
	PermissionHandler    handler.PermissionHandler
	PolicyHandler        handler.PolicyHandler
}

func NewRouter() Router {
	return Router{
		AuthHandler:          handler.GetAuthHandlerInstance(),
		TokenHandler:         handler.GetTokenHandlerInstance(),
		GroupHandler:         handler.GetGroupHandlerInstance(),
		UserHandler:          handler.GetUserHandlerInstance(),
		UserGroupHandler:     handler.GetUserGroupHandlerInstance(),
		ServiceHandler:       handler.GetServiceHandlerInstance(),
		ServiceGroupHandler:  handler.GetServiceGroupHandlerInstance(),
		RoleHandler:          handler.GetRoleHandlerInstance(),
		OperatePolicyHandler: handler.GetOperatorPolicyHandlerInstance(),
		UserServiceHandler:   handler.GetUserServiceHandlerInstance(),
		PermissionHandler:    handler.GetPermissionHandlerInstance(),
		PolicyHandler:        handler.GetPolicyHandlerInstance(),
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
	http.HandleFunc("/api/v1/auth", r.AuthHandler.Api)
	http.HandleFunc("/api/v1/token", r.TokenHandler.Api)

	http.HandleFunc("/api/v1/groups", r.GroupHandler.Api)

	http.HandleFunc("/api/v1/users", r.UserHandler.Api)
	http.HandleFunc("/api/v1/user_services", r.UserServiceHandler.Api)
	http.HandleFunc("/api/v1/user_groups", r.UserGroupHandler.Api)

	http.HandleFunc("/api/v1/services", r.ServiceHandler.Api)
	http.HandleFunc("/api/v1/service_groups", r.ServiceGroupHandler.Api)

	http.HandleFunc("/api/v1/roles", r.RoleHandler.Api)
	http.HandleFunc("/api/v1/permissions", r.PermissionHandler.Api)

	http.HandleFunc("/api/v1/policies", r.PolicyHandler.Api)
	http.HandleFunc("/api/v1/operator_policies", r.OperatePolicyHandler.Api)

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
