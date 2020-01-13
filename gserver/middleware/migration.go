package middleware

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/common/constant"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

const failedMigrationMsg = "Failed to not valid grant_n_z schema or data is broken for migration"

type Migration struct {
	userService           service.UserService
	roleService           service.RoleService
	operatorPolicyService service.OperatorPolicyService
	permissionService     service.PermissionService
}

func NewMigration() Migration {
	return Migration{
		userService:           service.GetUserServiceInstance(),
		roleService:           service.GetRoleServiceInstance(),
		operatorPolicyService: service.GetOperatorPolicyServiceInstance(),
		permissionService:     service.GetPermissionServiceInstance(),
	}
}

func (m Migration) V1() {
	if !m.checkV1Migration() {
		return
	}

	// Generate operator user
	operatorUser := entity.User{
		Id:        1,
		Username:  constant.Operator,
			Email: "operator@gmail.com",
		Password:  "grant_n_z_operator",
	}
	_, userErr := m.userService.InsertUser(operatorUser)
	if userErr != nil {
		if userErr.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate user for migration")
		}
	}
	log.Logger.Info("Generate to user for migration")

	// Generate operator role
	operatorRole := entity.Role{
		Id:   1,
		Name: constant.Operator,
	}
	_, roleErr1 := m.roleService.InsertRole(&operatorRole)
	if roleErr1 != nil {
		if roleErr1.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate operator role for migration")
		}
	}

	// Generate admin role
	adminRole := entity.Role{
		Id:   2,
		Name: constant.Admin,
	}
	_, roleErr2 := m.roleService.InsertRole(&adminRole)
	if roleErr2 != nil {
		if roleErr2.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate admin role for migration")
		}
	}
	log.Logger.Info("Generate to role for migration")

	// Generate admin permission
	adminPermission := entity.Permission{
		Id:   1,
		Name: constant.Admin,
	}
	_, permissionErr := m.permissionService.InsertPermission(&adminPermission)
	if permissionErr != nil {
		if permissionErr.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate admin permission for migration")
		}
	}
	log.Logger.Info("Generate to role for migration")

	// Generate operator operator_member_role
	operatorMemberRole := entity.OperatorPolicy{
		UserId: 1,
		RoleId: 1,
	}
	_, operatorRoleMemberErr := m.operatorPolicyService.Insert(&operatorMemberRole)
	if operatorRoleMemberErr != nil {
		if operatorRoleMemberErr.Code != http.StatusConflict {
			log.Logger.Fatal("Error generate operator policies for migration")
		}
	}
	log.Logger.Info("Generate to operator_policies for migration")
}

func (m Migration) checkV1Migration() bool {
	operatorAdminUser, err := m.userService.GetUserById(1)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Fatal(failedMigrationMsg)
	}

	operatorAdminRole, err := m.roleService.GetRoleByName(constant.Operator)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Info("Not found operator role")
		log.Logger.Fatal(failedMigrationMsg)
	}

	adminRole, err := m.roleService.GetRoleByName(constant.Admin)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Info("Not found admin role")
		log.Logger.Fatal(failedMigrationMsg)
	}

	adminPermission, err := m.permissionService.GetPermissionByName(constant.Admin)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Info("Not found admin permission")
		log.Logger.Fatal(failedMigrationMsg)
	}

	var operatorPolicy []*entity.OperatorPolicy
	operatorPolicy, err = m.operatorPolicyService.GetByUserId(1)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Fatal(failedMigrationMsg)
	}

	if operatorAdminUser != nil && operatorAdminRole != nil && adminRole != nil && adminPermission != nil && len(operatorPolicy) != 0 {
		log.Logger.Info("Skip to database migration")
		return false
	}

	return true
}
