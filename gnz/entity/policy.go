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
	PolicyServiceId
	PolicyUserGroupId
	PolicyCreatedAt
	PolicyUpdatedAt
)

type PolicyResponseBuilder interface {
	// Set policy name at response data
	SetName(name *string) PolicyResponseBuilder

	// Set role_name at response data
	SetRoleName(roleName *string) PolicyResponseBuilder

	// Set permission_name at response data
	SetPermissionName(permissionName *string) PolicyResponseBuilder

	// Set service_name at response data
	SetServiceName(serviceName *string) PolicyResponseBuilder

	// Set group_name at response data
	SetGroupName(groupName *string) PolicyResponseBuilder

	// Build PolicyResponse struct
	Build() PolicyResponse
}

// The table `policy` struct
type Policy struct {
	Id           int       `json:"id"`
	Name         string    `validate:"required"json:"name"`
	RoleId       int       `validate:"required"json:"role_id"`
	PermissionId int       `validate:"required"json:"permission_id"`
	ServiceId    int       `validate:"required"json:"service_id"`
	UserGroupId  int       `validate:"required"json:"user_group_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Policy request struct
type PolicyRequest struct {
	Name         string `validate:"required"json:"name"`
	ToUserEmail  string `validate:"required"json:"to_user_email"`
	RoleId       int    `validate:"required"json:"role_id"`
	PermissionId int    `validate:"required"json:"permission_id"`
}

// The api policy response struct
type PolicyResponse struct {
	Name           string `json:"policy_name"`
	RoleName       string `json:"role_name"`
	PermissionName string `json:"permission_name"`
	ServiceName    string `json:"service_name"`
	GroupName      string `json:"group_name"`
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

func (p PolicyResponse) SetServiceName(serviceName *string) PolicyResponseBuilder {
	if serviceName == nil {
		p.ServiceName = ""
	} else {
		p.ServiceName = *serviceName
	}
	return p
}

func (p PolicyResponse) SetGroupName(groupName *string) PolicyResponseBuilder {
	if groupName == nil {
		p.GroupName = ""
	} else {
		p.GroupName = *groupName
	}
	return p
}

func (p PolicyResponse) Build() PolicyResponse {
	return PolicyResponse{
		Name:           p.Name,
		RoleName:       p.RoleName,
		PermissionName: p.PermissionName,
		ServiceName:    p.ServiceName,
		GroupName:      p.GroupName,
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
	case PolicyServiceId:
		return "service_id"
	case PolicyUserGroupId:
		return "user_group_id"
	case PolicyCreatedAt:
		return "created_at"
	case PolicyUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
