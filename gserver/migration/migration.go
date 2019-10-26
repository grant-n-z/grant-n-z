package migration

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

const (
	Operator = "operator"
	Admin    = "admin"
)

type Migration struct {
	UserService           service.UserService
	RoleService           service.RoleService
	OperatorPolicyService service.OperatorPolicyService
}

func NewMigration() Migration {
	return Migration{
		UserService:           service.GetUserServiceInstance(),
		RoleService:           service.GetRoleServiceInstance(),
		OperatorPolicyService: service.GetOperatorPolicyServiceInstance(),
	}
}

func (m Migration) V1() {
	if !m.checkMigrationData() {
		return
	}

	// Generate operator user
	operatorUser := entity.User{
		Id:       1,
		Username: Operator,
		Email:    "operator@gmail.com",
		Password: "grant_n_z_operator",
	}
	_, userErr := m.UserService.InsertUser(&operatorUser)
	if userErr != nil {
		if userErr.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate user for migration")
		}
	}
	log.Logger.Info("Generate to user for migration")

	// Generate operator role
	operatorRole := entity.Role{
		Id:   1,
		Name: Operator,
	}
	_, roleErr1 := m.RoleService.InsertRole(&operatorRole)
	if roleErr1 != nil {
		if userErr.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate operator role for migration")
		}
	}

	// Generate admin role
	adminRole := entity.Role{
		Id:   2,
		Name: Admin,
	}
	_, roleErr2 := m.RoleService.InsertRole(&adminRole)
	if roleErr2 != nil {
		if userErr.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate admin role for migration")
		}
	}
	log.Logger.Info("Generate to role for migration")

	// Generate operator operator_member_role
	// TODO: Get role id
	operatorMemberRole := entity.OperatorPolicy{
		UserId: 1,
		RoleId: 1,
	}
	_, operatorRoleMemberErr := m.OperatorPolicyService.Insert(&operatorMemberRole)
	if operatorRoleMemberErr != nil {
		if userErr.Code != http.StatusConflict {
			log.Logger.Fatal("Error generate operator policies for migration")
		}
	}
	log.Logger.Info("Generate to operator_policies for migration")
}

func (m Migration) checkMigrationData() bool {
	operatorAdminUser, err := m.UserService.GetUserById(1)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Fatal("Failed to not valid grant_n_z schema or data is broken for migration")
	}

	operatorAdminRole, err := m.RoleService.GetRoleByName(Operator)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Info("Not found operator role")
		log.Logger.Fatal("Failed to not valid grant_n_z schema or data is broken for migration")
	}

	adminRole, err := m.RoleService.GetRoleByName(Admin)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Info("Not found admin role")
		log.Logger.Fatal("Failed to not valid grant_n_z schema or data is broken for migration")
	}

	var operatorPolicy []*entity.OperatorPolicy
	operatorPolicy, err = m.OperatorPolicyService.GetByUserId(1)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Fatal("Failed to not valid grant_n_z schema or data is broken for migration")
	}

	if operatorAdminUser != nil && operatorAdminRole != nil && adminRole != nil && len(operatorPolicy) != 0 {
		log.Logger.Info("Skip to database migration")
		return false
	}

	return true
}
