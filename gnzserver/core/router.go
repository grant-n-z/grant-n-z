package core

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tomoyane/grant-n-z/gnzserver/api/operator"
	v1 "github.com/tomoyane/grant-n-z/gnzserver/api/v1"
	"github.com/tomoyane/grant-n-z/gnzserver/api/v1/groups"
	"github.com/tomoyane/grant-n-z/gnzserver/api/v1/users"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

type Router struct {
	// Router
	mux *mux.Router

	// Http request Interceptor
	interceptor middleware.Interceptor

	// V1 endpoint
	Auth    v1.Auth
	Token   v1.Token
	Service v1.Service

	// V1 users endpoint
	UsersRouter UsersRouter

	// V1 groups endpoint
	GroupsRouter GroupsRouter

	// Operators endpoint
	OperatorsRouter OperatorsRouter
}

type UsersRouter struct {
	Group   users.Group
	User    users.User
	Service users.Service
	Policy  users.Policy
}

type GroupsRouter struct {
	Group      groups.Group
	User       groups.User
	Policy     groups.Policy
	Role       groups.Role
	Permission groups.Permission
}

type OperatorsRouter struct {
	OperatorPolicy operator.OperatorPolicy
	Service        operator.OperatorService
}

func NewRouter() Router {
	usersRouter := UsersRouter{
		Group:   users.GetGroupInstance(),
		User:    users.GetUserInstance(),
		Service: users.GetServiceInstance(),
		Policy:  users.GetPolicyInstance(),
	}

	groupsRouter := GroupsRouter{
		Group:      groups.GetGroupInstance(),
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
	r.mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	r.mux.HandleFunc("/api/v1/auth", r.interceptor.Intercept(r.Auth.Api))
	r.mux.HandleFunc("/api/v1/services", r.interceptor.Intercept(r.Service.Get)).Methods(http.MethodGet, http.MethodOptions)

	// Not required Client-Secret header
	r.mux.HandleFunc("/api/v1/token", r.interceptor.Intercept(r.Token.Api))

	// Required Client-Secret header
	r.mux.HandleFunc("/api/v1/services/add_user", r.interceptor.InterceptSecret(r.Service.Post)).Methods(http.MethodPost, http.MethodOptions)

	// Not required Client-Secret header
	user := func() {
		r.mux.HandleFunc("/api/v1/users", r.interceptor.Intercept(r.UsersRouter.User.Post)).Methods(http.MethodPost, http.MethodOptions)
		r.mux.HandleFunc("/api/v1/users", r.interceptor.InterceptAuthenticateUser(r.UsersRouter.User.Put)).Methods(http.MethodPut, http.MethodOptions)
		r.mux.HandleFunc("/api/v1/users/group", r.interceptor.InterceptAuthenticateUser(r.UsersRouter.Group.Api))
		r.mux.HandleFunc("/api/v1/users/service", r.interceptor.InterceptAuthenticateUser(r.UsersRouter.Service.Api))
		r.mux.HandleFunc("/api/v1/users/policy", r.interceptor.InterceptAuthenticateUser(r.UsersRouter.Policy.Api))
	}

	// Required Client-Secret and group admin permission
	group := func() {
		r.mux.HandleFunc("/api/v1/groups/{group_uuid}", r.interceptor.InterceptAuthenticateGroupUser(r.GroupsRouter.Group.Get)).Methods(http.MethodGet, http.MethodOptions)
		r.mux.HandleFunc("/api/v1/groups/{group_uuid}/user", r.interceptor.InterceptAuthenticateGroupAdmin(r.GroupsRouter.User.Api))
		r.mux.HandleFunc("/api/v1/groups/{group_uuid}/policy", r.interceptor.InterceptAuthenticateGroupAdmin(r.GroupsRouter.Policy.Api))
		r.mux.HandleFunc("/api/v1/groups/{group_uuid}/role", r.interceptor.InterceptAuthenticateGroupUser(r.GroupsRouter.Role.Get)).Methods(http.MethodGet, http.MethodOptions)
		r.mux.HandleFunc("/api/v1/groups/{group_uuid}/role", r.interceptor.InterceptAuthenticateGroupAdmin(r.GroupsRouter.Role.Post)).Methods(http.MethodPost, http.MethodOptions)
		r.mux.HandleFunc("/api/v1/groups/{group_uuid}/role", r.interceptor.InterceptAuthenticateGroupAdmin(r.GroupsRouter.Role.Delete)).Methods(http.MethodDelete, http.MethodOptions)
		r.mux.HandleFunc("/api/v1/groups/{group_uuid}/permission", r.interceptor.InterceptAuthenticateGroupUser(r.GroupsRouter.Permission.Get)).Methods(http.MethodGet, http.MethodOptions)
		r.mux.HandleFunc("/api/v1/groups/{group_uuid}/permission", r.interceptor.InterceptAuthenticateGroupAdmin(r.GroupsRouter.Permission.Post)).Methods(http.MethodPost, http.MethodOptions)
		r.mux.HandleFunc("/api/v1/groups/{group_uuid}/permission", r.interceptor.InterceptAuthenticateGroupAdmin(r.GroupsRouter.Permission.Delete)).Methods(http.MethodDelete, http.MethodOptions)
	}

	user()
	group()
}

func (r Router) operators() {
	r.mux.HandleFunc("/api/operators/service", r.interceptor.InterceptAuthenticateOperator(r.OperatorsRouter.Service.Api))
	//r.mux.HandleFunc("/api/operators/role", r.OperatorsRouter.OperatorService.Api)
	//r.mux.HandleFunc("/api/operators/permission", r.OperatorsRouter.OperatorService.Api)
	//r.mux.HandleFunc("/api/operators/policy", r.OperatorsRouter.OperatorService.Api)
}
