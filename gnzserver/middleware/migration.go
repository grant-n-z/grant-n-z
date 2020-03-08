package middleware

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/config"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

const failedMigrationMsg = "Failed to not valid grant_n_z schema or data is broken for migration"

type Migration struct {
	service               service.Service
	userService           service.UserService
	roleService           service.RoleService
	operatorPolicyService service.OperatorPolicyService
	permissionService     service.PermissionService
}

func NewMigration() Migration {
	return Migration{
		service:               service.GetServiceInstance(),
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
		Id:       1,
		Username: config.OperatorRole,
		Email:    "operator@gmail.com",
		Password: "grant_n_z_operator",
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
		Name: config.OperatorRole,
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
		Name: config.AdminRole,
	}
	_, roleErr2 := m.roleService.InsertRole(&adminRole)
	if roleErr2 != nil {
		if roleErr2.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate admin role for migration")
		}
	}

	// Generate user role
	userRole := entity.Role{
		Id:   3,
		Name: config.UserRole,
	}
	_, roleErr3 := m.roleService.InsertRole(&userRole)
	if roleErr3 != nil {
		if roleErr3.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate user role for migration")
		}
	}
	log.Logger.Info("Generate to role for migration")

	// Generate admin permission
	adminPermission := entity.Permission{
		Id:   1,
		Name: config.AdminPermission,
	}
	_, permissionErr01 := m.permissionService.InsertPermission(&adminPermission)
	if permissionErr01 != nil {
		if permissionErr01.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate admin permission for migration")
		}
	}

	// Generate read permission
	readPermission := entity.Permission{
		Id:   2,
		Name: config.ReadPermission,
	}
	_, permissionErr02 := m.permissionService.InsertPermission(&readPermission)
	if permissionErr02 != nil {
		if permissionErr02.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate read permission for migration")
		}
	}

	// Generate write permission
	writePermission := entity.Permission{
		Id:   3,
		Name: config.WritePermission,
	}
	_, permissionErr03 := m.permissionService.InsertPermission(&writePermission)
	if permissionErr03 != nil {
		if permissionErr03.Code != http.StatusConflict {
			log.Logger.Fatal("Failed to generate write permission for migration")
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

	operatorAdminRole, err := m.roleService.GetRoleByName(config.OperatorRole)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Info("Not found operator role")
		log.Logger.Fatal(failedMigrationMsg)
	}

	adminRole, err := m.roleService.GetRoleByName(config.AdminRole)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Info("Not found admin role")
		log.Logger.Fatal(failedMigrationMsg)
	}

	adminPermission, err := m.permissionService.GetPermissionByName(config.AdminPermission)
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
