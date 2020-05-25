package structure

// The `user_policy` struct in etcd
type UserPolicy struct {
	ServiceUuid    string `json:"service_uuid"`
	GroupUuid      string `json:"group_uuid"`
	RoleName       string `json:"role_name"`
	PermissionName string `json:"permission_name"`
}
