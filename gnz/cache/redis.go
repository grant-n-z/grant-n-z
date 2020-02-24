package cache

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/go-redis/redis"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var rcInstance RedisClient

type RedisClient interface {
	// Set policy
	SetPolicy(policy entity.Policy)

	// Set policy with expires
	SetPolicyWithExpires(policy entity.Policy, expiresMinutes time.Duration)

	// Set permission
	SetPermission(permission entity.Permission)

	// Set permission with expires
	SetPermissionWithExpires(permission entity.Permission, expiresMinutes time.Duration)

	// Set role
	SetRole(role entity.Role)

	// Set role with expires
	SetRoleWithExpires(role entity.Role, expiresMinutes time.Duration)

	// Set service
	SetService(service entity.Service)

	// Set service with expires
	SetServiceWithExpires(service entity.Service, expiresMinutes time.Duration)

	// Get policy
	GetPolicyById(policyId int) *entity.Policy
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

func (rc RedisClientImpl) SetPolicy(policy entity.Policy) {
	policyJson, _ := json.Marshal(policy)
	rc.Connection.Set(fmt.Sprintf("policy=%d", policy.Id), policyJson, 0)
}

func (rc RedisClientImpl) SetPolicyWithExpires(policy entity.Policy, expiresMinutes time.Duration) {
	policyJson, _ := json.Marshal(policy)
	rc.Connection.Set(fmt.Sprintf("policy=%d", policy.Id), policyJson, expiresMinutes)
}

func (rc RedisClientImpl) SetPermission(permission entity.Permission) {
	permissionJson, _ := json.Marshal(permission)
	rc.Connection.Set(fmt.Sprintf("permission=%d", permission.Id), permissionJson, 0)
}

func (rc RedisClientImpl) SetPermissionWithExpires(permission entity.Permission, expiresMinutes time.Duration) {
	permissionJson, _ := json.Marshal(permission)
	rc.Connection.Set(fmt.Sprintf("permission=%d", permission.Id), permissionJson, expiresMinutes)
}

func (rc RedisClientImpl) SetRole(role entity.Role) {
	roleJson, _ := json.Marshal(role)
	rc.Connection.Set(fmt.Sprintf("role=%d", role.Id), roleJson, 0)
}

func (rc RedisClientImpl) SetRoleWithExpires(role entity.Role, expiresMinutes time.Duration) {
	roleJson, _ := json.Marshal(role)
	rc.Connection.Set(fmt.Sprintf("role=%d", role.Id), roleJson, 0)
}

func (rc RedisClientImpl) SetService(service entity.Service) {
	serviceJson, _ := json.Marshal(service)
	rc.Connection.Set(fmt.Sprintf("service=%d", service.Id), serviceJson, 0)
}

func (rc RedisClientImpl) SetServiceWithExpires(service entity.Service, expiresMinutes time.Duration) {
	serviceJson, _ := json.Marshal(service)
	rc.Connection.Set(fmt.Sprintf("service=%d", service.Id), serviceJson, 0)
}

func (rc RedisClientImpl) GetPolicyById(policyId int) *entity.Policy {
	policyJson := rc.Connection.Get(fmt.Sprintf("policy=%d", policyId))

	var policy entity.Policy
	err := json.Unmarshal([]byte(policyJson.String()), &policy)
	if err != nil {
		log.Logger.Error("Failed to convert json to struct for policy cache", err.Error())
		return nil
	}
	return &policy
}
