package route

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tomoyane/grant-n-z/gserver/api/operator"
	"github.com/tomoyane/grant-n-z/gserver/api/v1"
	"github.com/tomoyane/grant-n-z/gserver/api/v1/groups"
	"github.com/tomoyane/grant-n-z/gserver/api/v1/users"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/middleware"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type Router struct {
	mux         *mux.Router
	interceptor middleware.Interceptor

	Auth  v1.Auth
	Token v1.Token

	UsersRouter     UsersRouter
	GroupsRouter    GroupsRouter
	OperatorsRouter OperatorsRouter
}

type UsersRouter struct {
	Group   users.Group
	User    users.User
	Service users.Service
	Policy  users.Policy
}

type GroupsRouter struct {
	Role       groups.Role
	Permission groups.Permission
}

type OperatorsRouter struct {
	OperatorPolicy operator.OperatorPolicy
	Service        operator.Service
}

func NewRouter() Router {
	usersRouter := UsersRouter{
		Group:   users.GetGroupInstance(),
		User:    users.GetUserInstance(),
		Service: users.GetServiceInstance(),
		Policy:  users.GetPolicyInstance(),
	}

	groupsRouter := GroupsRouter{
		Role:       groups.GetRoleInstance(),
		Permission: groups.GetPermissionInstance(),
	}

	operatorsRouter := OperatorsRouter{
		OperatorPolicy: operator.GetOperatorPolicyInstance(),
		Service:        operator.GetOperatorServiceInstance(),
	}

	return Router{
		mux:         mux.NewRouter(),
		interceptor: middleware.GetInterceptorInstance(),

		Auth:  v1.GetAuthInstance(),
		Token: v1.GetTokenInstance(),

		UsersRouter:     usersRouter,
		GroupsRouter:    groupsRouter,
		OperatorsRouter: operatorsRouter,
	}
}

func (r Router) Run() *mux.Router {
	r.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := model.NotFound("Not found resource path.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res.ToJson()))
	})

	r.v1()
	r.operators()
	r.admin()
	return r.mux
}

func (r Router) v1() {
	r.mux.HandleFunc("/api/v1/auth", r.Auth.Api)
	r.mux.HandleFunc("/api/v1/token", r.interceptor.Intercept(r.Token.Api))

	user := func() {
		r.mux.HandleFunc("/api/v1/users", r.interceptor.Intercept(r.UsersRouter.User.Post))
		r.mux.HandleFunc("/api/v1/users/{user_id:[0-9]+}", r.interceptor.InterceptAuthenticateUser(r.UsersRouter.User.Put))
		r.mux.HandleFunc("/api/v1/users/{user_id:[0-9]+}/group", r.UsersRouter.Group.Api)
		r.mux.HandleFunc("/api/v1/users/{user_id:[0-9]+}/service", r.UsersRouter.Service.Api)
		r.mux.HandleFunc("/api/v1/users/{user_id:[0-9]+}/policy", r.UsersRouter.Policy.Api)
	}

	group := func() {
		r.mux.HandleFunc("/api/v1/groups/{group_id}/user", r.GroupsRouter.Role.Api)
		r.mux.HandleFunc("/api/v1/groups/{group_id}/role", r.GroupsRouter.Role.Api)
		r.mux.HandleFunc("/api/v1/groups/{group_id}/permission", r.GroupsRouter.Permission.Api)
		r.mux.HandleFunc("/api/v1/groups/{group_id}/users/{to_user_id}/policy", r.GroupsRouter.Role.Api)
	}

	user()
	group()

	// TODO: update route info
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

func (r Router) operators() {
	// TODO: update route info
	r.mux.HandleFunc("/api/operator/service", r.OperatorsRouter.Service.Api)
	r.mux.HandleFunc("/api/operators/role", r.OperatorsRouter.Service.Api)
	r.mux.HandleFunc("/api/operators/permission", r.OperatorsRouter.Service.Api)
	r.mux.HandleFunc("/api/operators/policy", r.OperatorsRouter.Service.Api)
}

func (r Router) admin() {
	// TODO
}
