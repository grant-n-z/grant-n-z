package router

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/handler"
	"github.com/tomoyane/grant-n-z/gserver/log"
)

type Router struct {
	AuthHandler               handler.AuthHandler
	TokenHandler              handler.TokenHandler
	UserHandler               handler.UserHandler
	ServiceHandler            handler.ServiceHandler
	RoleHandler               handler.RoleHandler
	OperatorMemberRoleHandler handler.OperateMemberRoleHandler
	UserServiceHandler        handler.UserServiceHandler
	PermissionHandler         handler.PermissionHandler
	PolicyHandler             handler.PolicyHandler
	ServiceMemberRoleHandler  handler.ServiceMemberRoleHandler
}

func NewRouter() Router {
	return Router{
		AuthHandler:               handler.GetAuthHandlerInstance(),
		TokenHandler:              handler.GetTokenHandlerInstance(),
		UserHandler:               handler.GetUserHandlerInstance(),
		ServiceHandler:            handler.GetServiceHandlerInstance(),
		RoleHandler:               handler.GetRoleHandlerInstance(),
		OperatorMemberRoleHandler: handler.GetOperateMemberRoleHandlerInstance(),
		UserServiceHandler:        handler.GetUserServiceHandlerInstance(),
		PermissionHandler:         handler.GetPermissionHandlerInstance(),
		PolicyHandler:             handler.GetPolicyHandlerInstance(),
		ServiceMemberRoleHandler:  handler.GetServiceMemberRoleHandlerInstance(),
	}
}

func (r Router) Init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(map[string]string{"message": "Not found."})
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
	})
}

func (r Router) V1() {
	http.HandleFunc("/api/v1/auth", r.AuthHandler.Api)
	http.HandleFunc("/api/v1/token", r.TokenHandler.Api)
	http.HandleFunc("/api/v1/users", r.UserHandler.Api)
	http.HandleFunc("/api/v1/services", r.ServiceHandler.Api)
	http.HandleFunc("/api/v1/roles", r.RoleHandler.Api)
	http.HandleFunc("/api/v1/user-services", r.UserServiceHandler.Api)
	http.HandleFunc("/api/v1/permissions", r.PermissionHandler.Api)
	http.HandleFunc("/api/v1/policies", r.PolicyHandler.Api)
	http.HandleFunc("/api/v1/operator-member-roles", r.OperatorMemberRoleHandler.Api)
	http.HandleFunc("/api/v1/service-member-roles", r.ServiceMemberRoleHandler.Api)

	log.Logger.Info("------ Routing info ------")
	log.Logger.Info("HTTP Method: `POST` Routing: /api/v1/oauth")
	log.Logger.Info("HTTP Method: `POST`, `PUT` Routing: /api/v1/users")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/services`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/roles`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/user-services`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/permissions`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/policies`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/operator-member-roles`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/service-member-roles`")
	log.Logger.Info("------ Routing info ------")
}
