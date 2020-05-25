package middleware

import (
	"github.com/google/uuid"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/common"
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
	// Create only first time
	operatorUser := entity.User{
		Uuid:     uuid.New(),
		Username: common.OperatorRole,
		Email:    "operator@gmail.com",
		Password: "grant_n_z_operator",
	}
	savedOperatorUser, userErr := m.userService.InsertUser(operatorUser)
	if userErr != nil {
		if userErr.Code != http.StatusConflict {
			panic("Failed to generate user for migration")
		}
	}
	log.Logger.Info("Generate to user for migration")

	// Generate operator role
	operatorRole := entity.Role{
		Uuid: uuid.New(),
		Name: common.OperatorRole,
	}
	_, roleErr1 := m.roleService.InsertRole(&operatorRole)
	if roleErr1 != nil {
		if roleErr1.Code != http.StatusConflict {
			panic("Failed to generate operator role for migration")
		}
	}

	// Generate admin role
	adminRole := entity.Role{
		Uuid: uuid.New(),
		Name: common.AdminRole,
	}
	_, roleErr2 := m.roleService.InsertRole(&adminRole)
	if roleErr2 != nil {
		if roleErr2.Code != http.StatusConflict {
			panic("Failed to generate admin role for migration")
		}
	}

	// Generate user role
	userRole := entity.Role{
		Uuid: uuid.New(),
		Name: common.UserRole,
	}
	_, roleErr3 := m.roleService.InsertRole(&userRole)
	if roleErr3 != nil {
		if roleErr3.Code != http.StatusConflict {
			panic("Failed to generate user role for migration")
		}
	}
	log.Logger.Info("Generate to role for migration")

	// Generate admin permission
	adminPermission := entity.Permission{
		Uuid: uuid.New(),
		Name: common.AdminPermission,
	}
	_, permissionErr01 := m.permissionService.InsertPermission(&adminPermission)
	if permissionErr01 != nil {
		if permissionErr01.Code != http.StatusConflict {
			panic("Failed to generate admin permission for migration")
		}
	}

	// Generate read permission
	readPermission := entity.Permission{
		Uuid: uuid.New(),
		Name: common.ReadPermission,
	}
	_, permissionErr02 := m.permissionService.InsertPermission(&readPermission)
	if permissionErr02 != nil {
		if permissionErr02.Code != http.StatusConflict {
			panic("Failed to generate read permission for migration")
		}
	}

	// Generate write permission
	writePermission := entity.Permission{
		Uuid: uuid.New(),
		Name: common.WritePermission,
	}
	_, permissionErr03 := m.permissionService.InsertPermission(&writePermission)
	if permissionErr03 != nil {
		if permissionErr03.Code != http.StatusConflict {
			panic("Failed to generate write permission for migration")
		}
	}
	log.Logger.Info("Generate to role for migration")

	// Generate operator operator_member_role
	operatorMemberRole := entity.OperatorPolicy{
		UserUuid: savedOperatorUser.Uuid,
		RoleUuid: operatorRole.Uuid,
	}
	_, operatorRoleMemberErr := m.operatorPolicyService.Insert(&operatorMemberRole)
	if operatorRoleMemberErr != nil {
		if operatorRoleMemberErr.Code != http.StatusConflict {
			panic("Error generate operator policies for migration")
		}
	}
	log.Logger.Info("Generate to operator_policies for migration")
}

func (m Migration) checkV1Migration() bool {
	operatorAdminRole, err := m.roleService.GetRoleByName(common.OperatorRole)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Info("Not found operator role")
		panic(failedMigrationMsg)
	}

	adminRole, err := m.roleService.GetRoleByName(common.AdminRole)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Info("Not found admin role")
		panic(failedMigrationMsg)
	}

	adminPermission, err := m.permissionService.GetPermissionByName(common.AdminPermission)
	if err != nil && err.Code != http.StatusNotFound {
		log.Logger.Info("Not found admin permission")
		panic(failedMigrationMsg)
	}
		var operatorPolicy []*entity.OperatorPolicy
		operatorPolicy, err = m.operatorPolicyService.GetAll()
		if err != nil && err.Code != http.StatusNotFound {
			panic(failedMigrationMsg)
		}

	if operatorAdminRole != nil && adminRole != nil && adminPermission != nil && len(operatorPolicy) != 0 {
		log.Logger.Info("Skip to database migration")
		return false
	}

	return true
}
