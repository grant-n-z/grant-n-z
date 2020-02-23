package model

import "github.com/tomoyane/grant-n-z/gnz/entity"

// The table `groups` and `user_groups` and `policy` struct
type GroupWithUserGroupWithPolicy struct {
	entity.Group
	entity.UserGroup
	entity.Policy
}
