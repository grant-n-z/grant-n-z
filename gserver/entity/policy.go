package entity

import (
	"time"
)

const (
	PolicyTable PolicyTableConfig = iota
	PolicyId
	PolicyName
	PolicyRoleId
	PolicyPermissionId
	PolicyUserGroupId
	PolicyCreatedAt
	PolicyUpdatedAt
)

type PolicyResponseBuilder interface {
	// Set name at response data
	SetName(name *string) PolicyResponseBuilder

	// Set role_name at response data
	SetRoleName(roleName *string) PolicyResponseBuilder

	// Set permission_name at response data
	SetPermissionName(permissionName *string) PolicyResponseBuilder

	// Build PolicyResponse struct
	Build() PolicyResponse
}

// The table `policy` struct
type Policy struct {
	Id           int       `json:"id"`
	Name         string    `validate:"required"json:"name"`
	RoleId       int       `validate:"required"json:"role_id"`
	PermissionId int       `validate:"required"json:"permission_id"`
	UserGroupId  int       `validate:"required"json:"user_group_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// The api policy response struct
type PolicyResponse struct {
	Name           string `json:"name"`
	RoleName       string `json:"role_name"`
	PermissionName string `json:"permission_name"`
}

// Policy table config struct
type PolicyTableConfig int

// PolicyResponse constructor
func NewPolicyResponse() PolicyResponseBuilder {
	return &PolicyResponse{}
}

func (p PolicyResponse) SetName(name *string) PolicyResponseBuilder {
	if name == nil {
		p.Name = ""
	} else {
		p.Name = *name
	}
	return p
}

func (p PolicyResponse) SetRoleName(roleName *string) PolicyResponseBuilder {
	if roleName == nil {
		p.RoleName = ""
	} else {
		p.RoleName = *roleName
	}
	return p
}

func (p PolicyResponse) SetPermissionName(permissionName *string) PolicyResponseBuilder {
	if permissionName == nil {
		p.PermissionName = ""
	} else {
		p.PermissionName = *permissionName
	}
	return p
}

func (p PolicyResponse) Build() PolicyResponse {
	return PolicyResponse{
		Name:           p.Name,
		RoleName:       p.RoleName,
		PermissionName: p.PermissionName,
	}
}

func (pc PolicyTableConfig) String() string {
	switch pc {
	case PolicyTable:
		return "policies"
	case PolicyId:
		return "id"
	case PolicyName:
		return "name"
	case PolicyRoleId:
		return "role_id"
	case PolicyPermissionId:
		return "permission_id"
	case PolicyUserGroupId:
		return "user_group_id"
	case PolicyCreatedAt:
		return "created_at"
	case PolicyUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
