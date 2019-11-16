package route

import (
	"github.com/tomoyane/grant-n-z/gserver/api/admin"
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
	Service        v1.Service
	ServiceGroup   v1.ServiceGroup
	Role           v1.Role
	OperatorPolicy v1.OperatorPolicy
	UserService    v1.UserService
	Permission     v1.Permission
	Policy         v1.Policy

	AdminService admin.AdminService
}

func NewRouter() Router {
	return Router{
		Auth:           v1.GetAuthInstance(),
		Token:          v1.GetTokenInstance(),
		Group:          v1.GetGroupInstance(),
		User:           v1.GetUserInstance(),
		Service:        v1.GetServiceInstance(),
		ServiceGroup:   v1.GetServiceGroupInstance(),
		Role:           v1.GetRoleInstance(),
		OperatorPolicy: v1.GetOperatorPolicyInstance(),
		UserService:    v1.GetUserServiceInstance(),
		Permission:     v1.GetPermissionInstance(),
		Policy:         v1.GetPolicyInstance(),
		AdminService:   admin.GetAdminServiceInstance(),
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
	// Verify token
	http.HandleFunc("/api/v1/auth", r.Auth.Api)

	// Generate user token, operator token
	http.HandleFunc("/api/v1/token", r.Token.Api)

	// Control group of user
	http.HandleFunc("/api/v1/groups", r.Group.Api)

	// Control create to user, update to user
	http.HandleFunc("/api/v1/users", r.User.Api)

	http.HandleFunc("/api/v1/services", r.Service.Api)

	http.HandleFunc("/api/v1/user_services", r.UserService.Api)

	http.HandleFunc("/api/v1/roles", r.Role.Api)

	http.HandleFunc("/api/v1/permissions", r.Permission.Api)

	http.HandleFunc("/api/v1/policies", r.Policy.Api)

	http.HandleFunc("/api/v1/operator_policies", r.OperatorPolicy.Api)

	log.Logger.Info("------ Routing info ------")
	log.Logger.Info("Routing: /api/v1/oauth")
	log.Logger.Info("Routing: /api/v1/groups")
	log.Logger.Info("Routing: /api/v1/users")
	log.Logger.Info("Routing: /api/v1/user_service")
	log.Logger.Info("Routing: /api/v1/user_groups")
	log.Logger.Info("Routing: /api/v1/services")
	log.Logger.Info("Routing: /api/v1/service_groups")
	log.Logger.Info("Routing: /api/v1/roles")
	log.Logger.Info("Routing: /api/v1/permissions")
	log.Logger.Info("Routing: /api/v1/policies")
	log.Logger.Info("Routing: /api/v1/operator_policies")
	log.Logger.Info("------ Routing info ------")
}

func (r Router) Admin() {
	http.HandleFunc("/api/admin/services", r.AdminService.Api)
}
