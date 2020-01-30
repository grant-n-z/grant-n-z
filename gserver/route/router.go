package route

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tomoyane/grant-n-z/gserver/api/operator"
	"github.com/tomoyane/grant-n-z/gserver/api/v1"
	"github.com/tomoyane/grant-n-z/gserver/api/v1/groups"
	"github.com/tomoyane/grant-n-z/gserver/api/v1/users"
	"github.com/tomoyane/grant-n-z/gserver/middleware"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type Router struct {
	mux         *mux.Router
	interceptor middleware.Interceptor

	Auth    v1.Auth
	Token   v1.Token
	Service v1.Service

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
	User       groups.User
	Policy     groups.Policy
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
		User:       groups.GetUserInstance(),
		Policy:     groups.GetPolicyInstance(),
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

		Auth:    v1.GetAuthInstance(),
		Token:   v1.GetTokenInstance(),
		Service: v1.GetServiceInstance(),

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
	return r.mux
}

func (r Router) v1() {
	// No restriction
	r.mux.HandleFunc("/api/v1/auth", r.interceptor.InterceptHeader(r.Auth.Api))
	r.mux.HandleFunc("/api/v1/services", r.interceptor.InterceptHeader(r.Service.Get)).Methods(http.MethodGet)

	// Required Api-Key header
	r.mux.HandleFunc("/api/v1/token", r.interceptor.Intercept(r.Token.Api))
	r.mux.HandleFunc("/api/v1/services/add_user", r.interceptor.Intercept(r.Service.Post)).Methods(http.MethodPost)

	// Required Api-Key header
	user := func() {
		r.mux.HandleFunc("/api/v1/users", r.interceptor.Intercept(r.UsersRouter.User.Post)).Methods(http.MethodPost)
		r.mux.HandleFunc("/api/v1/users", r.interceptor.InterceptAuthenticateUser(r.UsersRouter.User.Put)).Methods(http.MethodPut)
		r.mux.HandleFunc("/api/v1/users/group", r.interceptor.InterceptAuthenticateUser(r.UsersRouter.Group.Api))
		r.mux.HandleFunc("/api/v1/users/policy", r.interceptor.InterceptAuthenticateUser(r.UsersRouter.Policy.Api))
	}

	// Required Api-Key and group admin
	group := func() {
		r.mux.HandleFunc("/api/v1/groups/{group_id}/user", r.interceptor.InterceptAuthenticateGroupAdmin(r.GroupsRouter.User.Api))
		r.mux.HandleFunc("/api/v1/groups/{group_id}/policy", r.interceptor.InterceptAuthenticateGroupAdmin(r.GroupsRouter.Policy.Api))
		r.mux.HandleFunc("/api/v1/groups/{group_id}/role", r.interceptor.InterceptAuthenticateGroupAdmin(r.GroupsRouter.Role.Api))
		r.mux.HandleFunc("/api/v1/groups/{group_id}/permission", r.interceptor.InterceptAuthenticateGroupAdmin(r.GroupsRouter.Permission.Api))
	}

	user()
	group()
}

func (r Router) operators() {
	// TODO: update route info
	r.mux.HandleFunc("/api/operators/service", r.interceptor.InterceptAuthenticateOperator(r.OperatorsRouter.Service.Api))
	r.mux.HandleFunc("/api/operators/role", r.OperatorsRouter.Service.Api)
	r.mux.HandleFunc("/api/operators/permission", r.OperatorsRouter.Service.Api)
	r.mux.HandleFunc("/api/operators/policy", r.OperatorsRouter.Service.Api)
}
