package cache

import (
	"fmt"
	"strings"
	"time"

	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var rcInstance RedisClient

type RedisClient interface {
	// Set policy with expires
	SetPolicy(policy entity.Policy, expiresMinutes time.Duration)

	// Set permission with expires
	SetPermission(permission entity.Permission, expiresMinutes time.Duration)

	// Set role with expires
	SetRole(role entity.Role, expiresMinutes time.Duration)

	// Set service with expires
	SetService(service entity.Service, expiresMinutes time.Duration)

	// Get policy by id
	GetPolicyById(id int) *entity.Policy

	// Get policy by name
	GetPolicyByName(name string) *entity.Policy

	// Get policy by names
	GetPolicyByNames(names []string) []*entity.Policy

	// Get permission by id
	GetPermissionById(id int) *entity.Policy

	// Get permission by name
	GetPermissionByName(name string) *entity.Policy

	// Get permission by names
	GetPermissionByNames(names []string) []*entity.Policy

	// Get role by id
	GetRoleById(id int) *entity.Policy

	// Get role by name
	GetRoleByName(name string) *entity.Policy

	// Get role by names
	GetRoleByNames(names []string) []*entity.Policy

	// Get service by id
	GetServiceById(id int) *entity.Policy

	// Get service by name
	GetServiceByName(name string) *entity.Policy

	// Get service by names
	GetServiceByNames(names []string) []*entity.Policy
}

type RedisClientImpl struct {
	Connection *redis.Client
}

func GetRedisClientInstance() RedisClient {
	if rcInstance == nil {
		rcInstance = NewRedisClient()
	}
	return rcInstance
}

func NewRedisClient() RedisClient {
	log.Logger.Info("New `RedisClient` instance")
	return RedisClientImpl{Connection: connection}
}

func (rc RedisClientImpl) SetPolicy(policy entity.Policy, expiresMinutes time.Duration) {
	policyJson, _ := json.Marshal(policy)
	rc.Connection.Set(fmt.Sprintf("policy=%d", policy.Id), policyJson, expiresMinutes)
	rc.Connection.Set(fmt.Sprintf("policy=%s", policy.Name), policyJson, expiresMinutes)
}

func (rc RedisClientImpl) SetPermission(permission entity.Permission, expiresMinutes time.Duration) {
	permissionJson, _ := json.Marshal(permission)
	rc.Connection.Set(fmt.Sprintf("permission=%d", permission.Id), permissionJson, expiresMinutes)
	rc.Connection.Set(fmt.Sprintf("permission=%s", permission.Name), permissionJson, expiresMinutes)
}

func (rc RedisClientImpl) SetRole(role entity.Role, expiresMinutes time.Duration) {
	roleJson, _ := json.Marshal(role)
	rc.Connection.Set(fmt.Sprintf("role=%d", role.Id), roleJson, expiresMinutes)
	rc.Connection.Set(fmt.Sprintf("role=%s", role.Name), roleJson, expiresMinutes)
}

func (rc RedisClientImpl) SetService(service entity.Service, expiresMinutes time.Duration) {
	serviceJson, _ := json.Marshal(service)
	rc.Connection.Set(fmt.Sprintf("service=%d", service.Id), serviceJson, expiresMinutes)
	rc.Connection.Set(fmt.Sprintf("service=%s", service.Name), serviceJson, expiresMinutes)
}

func (rc RedisClientImpl) GetPolicyById(id int) *entity.Policy {
	policyJson := rc.Connection.Get(fmt.Sprintf("policy=%d", id)).String()
	if strings.EqualFold(policyJson, ""){
		return nil
	}

	var policy entity.Policy
	err := json.Unmarshal([]byte(policyJson), &policy)
	if err != nil {
		log.Logger.Error("Failed to convert json to struct for policy cache", err.Error())
		return nil
	}
	return &policy
}

func (rc RedisClientImpl) GetPolicyByName(name string) *entity.Policy {
	panic("implement me")
}

func (rc RedisClientImpl) GetPolicyByNames(names []string) []*entity.Policy {
	panic("implement me")
}

func (rc RedisClientImpl) GetPermissionById(id int) *entity.Policy {
	panic("implement me")
}

func (rc RedisClientImpl) GetPermissionByName(name string) *entity.Policy {
	panic("implement me")
}

func (rc RedisClientImpl) GetPermissionByNames(names []string) []*entity.Policy {
	panic("implement me")
}

func (rc RedisClientImpl) GetRoleById(id int) *entity.Policy {
	panic("implement me")
}

func (rc RedisClientImpl) GetRoleByName(name string) *entity.Policy {
	panic("implement me")
}

func (rc RedisClientImpl) GetRoleByNames(names []string) []*entity.Policy {
	panic("implement me")
}

func (rc RedisClientImpl) GetServiceById(id int) *entity.Policy {
	panic("implement me")
}

func (rc RedisClientImpl) GetServiceByName(name string) *entity.Policy {
	panic("implement me")
}

func (rc RedisClientImpl) GetServiceByNames(names []string) []*entity.Policy {
	panic("implement me")
}
