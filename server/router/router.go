package router

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/server/handler"
	"github.com/tomoyane/grant-n-z/server/log"
)

type Router struct {
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
		TokenHandler:              handler.NewTokenHandler(),
		UserHandler:               handler.NewUserHandler(),
		ServiceHandler:            handler.NewServiceHandler(),
		RoleHandler:               handler.NewRoleHandler(),
		OperatorMemberRoleHandler: handler.NewOperatorMemberRoleHandler(),
		UserServiceHandler:        handler.NewUserServiceHandler(),
		PermissionHandler:         handler.NewPermissionHandler(),
		PolicyHandler:             handler.NewPolicyHandlerHandler(),
		ServiceMemberRoleHandler:  handler.NewServiceMemberRoleHandler(),
	}
}

func (r Router) V1() {
	http.HandleFunc("/api/v1/oauth", r.TokenHandler.Post)
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
	log.Logger.Info("HTTP Method: `POST` Routing: /api/v1/users")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/services`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/roles`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/user-services`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/permissions`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/policies`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/operator-member-roles`")
	log.Logger.Info("HTTP Method: `POST`, `GET` Routing: `/api/v1/service-member-roles`")
}
