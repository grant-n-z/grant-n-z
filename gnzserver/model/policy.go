package model

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
