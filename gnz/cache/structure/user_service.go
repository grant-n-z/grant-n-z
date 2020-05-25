package structure

// The `user_service` struct in etcd
type UserService struct {
	ServiceUUid string `json:"service_uuid"`
	ServiceName string `json:"service_name"`
}
