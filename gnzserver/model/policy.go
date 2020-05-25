package model

import "github.com/google/uuid"

// Builder interface
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

// Policy request struct
type PolicyRequest struct {
	Name           string `validate:"required"json:"name"`
	ToUserEmail    string `validate:"required"json:"to_user_email"`
	RoleUuid       string `validate:"required"json:"role_uuid"`
	PermissionUuid string `validate:"required"json:"permission_uuid"`
}

// The api policy response struct
type PolicyResponse struct {
	Name           string    `json:"policy_name"`
	RoleName       string    `json:"role_name"`
	RoleUuid       string    `json:"role_uuid"`
	PermissionName string    `json:"permission_name"`
	PermissionUuid uuid.UUID `json:"permission_uuid"`
	ServiceName    string    `json:"service_name"`
	ServiceUuid    uuid.UUID `json:"service_uuid"`
	GroupName      string    `json:"group_name"`
	GroupUuid      uuid.UUID `json:"group_uuid"`
}

// The user policy response struct
type UserPolicyOnGroupResponse struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	ServiceName    string `json:"service_name"`
	PolicyName     string `json:"policy_name"`
	RoleName       string `json:"role_name"`
	PermissionName string `json:"permission_name"`
}

// The user policy response struct
type UserPolicyOnServiceResponse struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	PolicyName     string `json:"policy_name"`
	GroupName      string `json:"group_name"`
	GroupUuid      string `json:"group_uuid"`
	RoleName       string `json:"role_name"`
	RoleUuid       string `json:"role_uuid"`
	PermissionName string `json:"permission_name"`
	PermissionUuid string `json:"permission_uuid"`
}

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
