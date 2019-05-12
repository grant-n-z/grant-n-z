package router

import (
	"github.com/tomoyane/grant-n-z/server/handler"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/log"
)

type Router struct {
	UserHandler        handler.UserHandler
	ServiceHandler     handler.ServiceHandler
	RoleHandler        handler.RoleHandler
	RoleMemberHandler  handler.RoleMemberHandler
	UserServiceHandler handler.UserServiceHandler
	PermissionHandler  handler.PermissionHandler
	PolicyHandler      handler.PolicyHandler
}

func (r Router) V1() {
	http.HandleFunc("/api/v1/users", r.UserServiceHandler.Api)
	http.HandleFunc("/api/v1/services", r.ServiceHandler.Api)
	http.HandleFunc("/api/v1/roles", r.RoleHandler.Api)
	http.HandleFunc("/api/v1/role-members", r.RoleMemberHandler.Api)
	http.HandleFunc("/api/v1/user-services", r.UserServiceHandler.Api)
	http.HandleFunc("/api/v1/permissions", r.PermissionHandler.Api)
	http.HandleFunc("/api/v1/policies", r.PolicyHandler.Api)

	log.Logger.Debug("____ routing info ____")
	log.Logger.Debug("Method: `POST` routing: /api/v1/users")
	log.Logger.Debug("Method: `POST`, `GET` Routing: `/api/v1/services`")
	log.Logger.Debug("Method: `POST`, `GET` Routing: `/api/v1/roles`")
	log.Logger.Debug("Method: `POST`, `GET` Routing: `/api/v1/role-members`")
	log.Logger.Debug("Method: `POST`, `GET` Routing: `/api/v1/user-services`")
	log.Logger.Debug("Method: `POST`, `GET` Routing: `/api/v1/permissions`")
	log.Logger.Debug("Method: `POST`, `GET` Routing: `/api/v1/policies`")
}

func (r Router) Run(port string) {
	log.Logger.Fatal(http.ListenAndServe(port, nil))
}
