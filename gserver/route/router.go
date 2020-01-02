package route

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api/operator"
	"github.com/tomoyane/grant-n-z/gserver/api/v1"
	"github.com/tomoyane/grant-n-z/gserver/api/v1/groups"
	"github.com/tomoyane/grant-n-z/gserver/api/v1/users"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type Router struct {
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
		Auth:  v1.GetAuthInstance(),
		Token: v1.GetTokenInstance(),

		UsersRouter:     usersRouter,
		GroupsRouter:    groupsRouter,
		OperatorsRouter: operatorsRouter,
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

	user := func() {
		http.HandleFunc("/api/v1/users", r.UsersRouter.User.Api)
		http.HandleFunc("/api/v1/users/group", r.UsersRouter.Group.Api)
		http.HandleFunc("/api/v1/users/service", r.UsersRouter.Service.Api)
		http.HandleFunc("/api/v1/users/policy", r.UsersRouter.Policy.Api)
	}

	group := func() {
		http.HandleFunc("/api/v1/groups/role", r.GroupsRouter.Role.Api)
		http.HandleFunc("/api/v1/groups/permission", r.GroupsRouter.Permission.Api)
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

func (r Router) Operators() {
	// TODO: update route info
	http.HandleFunc("/api/operator/service", r.OperatorsRouter.Service.Api)
	http.HandleFunc("/api/operators/role", r.OperatorsRouter.Service.Api)
	http.HandleFunc("/api/operators/permission", r.OperatorsRouter.Service.Api)
	http.HandleFunc("/api/operators/policy", r.OperatorsRouter.Service.Api)
}

func (r Router) Admin() {
	// TODO
}
