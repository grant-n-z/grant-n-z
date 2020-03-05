package cache

import (
	"errors"
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

	// Get policy by id or name
	GetPolicy(data interface{}) *entity.Policy

	// Get policy by names
	GetPolicyByNames(names []string) []entity.Policy

	// Get permission by id or name
	GetPermission(data interface{}) *entity.Permission

	// Get permission by names
	GetPermissionByNames(names []string) []entity.Permission

	// Get role by id or name
	GetRole(data interface{}) *entity.Role

	// Get role by names
	GetRoleByNames(names []string) []entity.Role

	// Get service by id or name
	GetService(data interface{}) *entity.Service

	// Get service by names
	GetServiceByNames(names []string) []entity.Service
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
	rc.set(policyJson, []string{fmt.Sprintf("policy=%d", policy.Id), fmt.Sprintf("policy=%s", policy.Name)}, expiresMinutes)
}

func (rc RedisClientImpl) SetPermission(permission entity.Permission, expiresMinutes time.Duration) {
	permissionJson, _ := json.Marshal(permission)
	rc.set(permissionJson, []string{fmt.Sprintf("permission=%d", permission.Id), fmt.Sprintf("permission=%s", permission.Name)}, expiresMinutes)
}

func (rc RedisClientImpl) SetRole(role entity.Role, expiresMinutes time.Duration) {
	roleJson, _ := json.Marshal(role)
	rc.set(roleJson, []string{fmt.Sprintf("role=%d", role.Id), fmt.Sprintf("role=%s", role.Name)}, expiresMinutes)
}

func (rc RedisClientImpl) SetService(service entity.Service, expiresMinutes time.Duration) {
	serviceJson, _ := json.Marshal(service)
	rc.set(serviceJson, []string{fmt.Sprintf("service=%d", service.Id), fmt.Sprintf("service=%s", service.Name)}, expiresMinutes)
}

func (rc RedisClientImpl) GetPolicy(data interface{}) *entity.Policy {
	var policy entity.Policy
	err := rc.get(fmt.Sprintf("policy=%s", data), policy)
	if err != nil {
		log.Logger.Info("Cloud not get cache", err.Error())
		return nil
	}
	return &policy
}
func (rc RedisClientImpl) GetPolicyByNames(names []string) []entity.Policy {
	var policies []entity.Policy
	for _, name := range names {
		var policy entity.Policy
		err := rc.get(fmt.Sprintf("policy=%s", name), policy)
		if err != nil {
			log.Logger.Info("Cloud not get cache", err.Error())
			continue
		}
		policies = append(policies, policy)
	}
	return policies
}

func (rc RedisClientImpl) GetPermission(data interface{}) *entity.Permission {
	var permission entity.Permission
	err := rc.get(fmt.Sprintf("permission=%s", data), permission)
	if err != nil {
		log.Logger.Info("Cloud not get cache", err.Error())
		return nil
	}
	return &permission
}

func (rc RedisClientImpl) GetPermissionByNames(names []string) []entity.Permission {
	var permissions []entity.Permission
	for _, name := range names {
		var permission entity.Permission
		err := rc.get(fmt.Sprintf("permission=%s", name), permission)
		if err != nil {
			log.Logger.Info("Cloud not get cache", err.Error())
			continue
		}
		permissions = append(permissions, permission)
	}
	return permissions
}

func (rc RedisClientImpl) GetRole(data interface{}) *entity.Role {
	var role entity.Role
	err := rc.get(fmt.Sprintf("role=%s", data), role)
	if err != nil {
		log.Logger.Info("Cloud not get cache", err.Error())
		return nil
	}
	return &role
}

func (rc RedisClientImpl) GetRoleByNames(names []string) []entity.Role {
	var roles []entity.Role
	for _, name := range names {
		var role entity.Role
		err := rc.get(fmt.Sprintf("role=%s", name), role)
		if err != nil {
			log.Logger.Info("Cloud not get cache", err.Error())
			continue
		}
		roles = append(roles, role)
	}
	return roles
}

func (rc RedisClientImpl) GetService(data interface{}) *entity.Service {
	var service entity.Service
	err := rc.get(fmt.Sprintf("service=%s", data), service)
	if err != nil {
		log.Logger.Info("Cloud not get cache", err.Error())
		return nil
	}
	return &service
}

func (rc RedisClientImpl) GetServiceByNames(names []string) []entity.Service {
	var services []entity.Service
	for _, name := range names {
		var service entity.Service
		err := rc.get(fmt.Sprintf("service=%s", name), service)
		if err != nil {
			log.Logger.Info("Cloud not get cache", err.Error())
			continue
		}
		services = append(services, service)
	}
	return services
}

// Get cache shared method
func (rc RedisClientImpl) get(key string, structData interface{}) error {
	start := time.Millisecond
	jsonData := rc.Connection.Get(key).String()
	end := time.Millisecond
	log.Logger.Info(fmt.Sprintf("[%sms] GET Redis key %s", end - start, key))
	if strings.EqualFold(jsonData, ""){
		return errors.New(fmt.Sprintf("Cache is null. key = %s", key))
	}

	err := json.Unmarshal([]byte(jsonData), &structData)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to convert json to struct for cache. %s", err.Error()))
	}
	return nil
}

// Set cache shared method
func (rc RedisClientImpl) set(json []byte, keys []string, expiresMinutes time.Duration) {
	for _, key := range keys {
		start := time.Millisecond
		rc.Connection.Set(key, json, expiresMinutes)
		end := time.Millisecond
		log.Logger.Info(fmt.Sprintf("[%s, %sms] SET Redis key %s", end, start, key))
	}
}
