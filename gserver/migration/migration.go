package migration

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/service"
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
	if !m.checkAdminUser() {
		return
	}

	// Generate operator user
	operatorUser := entity.User{
		Id:       1,
		Username: "operator",
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
		Name: "operator",
	}
	_, roleErr := m.RoleService.InsertRole(&operatorRole)
	if roleErr != nil {
		if userErr.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate user for migration")
		}
	}
	log.Logger.Info("Generate to role for migration")

	// Generate operator operator_member_role
	operatorMemberRole := entity.OperatorPolicy{
		UserId: 1,
		RoleId: 1,
	}
	_, operatorRoleMemberErr := m.OperatorPolicyService.Insert(&operatorMemberRole)
	if operatorRoleMemberErr != nil {
		if userErr.Code != http.StatusConflict {
			log.Logger.Fatal("Error generate operator_policies for migration")
		}
	}
	log.Logger.Info("Generate to operator_policies for migration")
}

func (m Migration) checkAdminUser() bool {
	operatorAdminUser, err := m.UserService.GetUserById(1)
	if err != nil {
		log.Logger.Fatal("Failed to not valid grant_n_z schema or data is broken for migration")
	}
	operatorAdminRole, err := m.RoleService.GetRoleById(1)
	if err != nil {
		log.Logger.Fatal("Failed to not valid grant_n_z schema or data is broken for migration")
	}

	var operatorAdminMemberRole []*entity.OperatorPolicy
	operatorAdminMemberRole, err = m.OperatorPolicyService.GetByUserId(1)
	if err != nil {
		log.Logger.Fatal("Failed to not valid grant_n_z schema or data is broken for migration")
	}

	if operatorAdminUser != nil && operatorAdminRole != nil && len(operatorAdminMemberRole) != 0 {
		log.Logger.Info("Skip to database migration")
		return false
	}

	return true
}
