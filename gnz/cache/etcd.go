package cache

import (
	"errors"
	"fmt"
	"time"

	"encoding/json"

	"go.etcd.io/etcd/clientv3"

	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var eInstance EtcdClient

type EtcdClient interface {
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

type EtcdClientImpl struct {
	Connection *clientv3.Client
}

func GetEtcdClientInstance() EtcdClient {
	if eInstance == nil {
		eInstance = NewEtcdClient()
	}
	return eInstance
}

func NewEtcdClient() EtcdClient {
	log.Logger.Info("New `EtcdClient` instance")
	return EtcdClientImpl{Connection: connection}
}

func (e EtcdClientImpl) SetPolicy(policy entity.Policy, expiresMinutes time.Duration) {
	policyJson, _ := json.Marshal(policy)
	e.set(policyJson, []string{fmt.Sprintf("policy=%d", policy.Id), fmt.Sprintf("policy=%s", policy.Name)}, expiresMinutes)
}

func (e EtcdClientImpl) SetPermission(permission entity.Permission, expiresMinutes time.Duration) {
	permissionJson, _ := json.Marshal(permission)
	e.set(permissionJson, []string{fmt.Sprintf("permission=%d", permission.Id), fmt.Sprintf("permission=%s", permission.Name)}, expiresMinutes)
}

func (e EtcdClientImpl) SetRole(role entity.Role, expiresMinutes time.Duration) {
	roleJson, _ := json.Marshal(role)
	e.set(roleJson, []string{fmt.Sprintf("role=%d", role.Id), fmt.Sprintf("role=%s", role.Name)}, expiresMinutes)
}

func (e EtcdClientImpl) SetService(service entity.Service, expiresMinutes time.Duration) {
	serviceJson, _ := json.Marshal(service)
	e.set(serviceJson, []string{fmt.Sprintf("service=%d", service.Id), fmt.Sprintf("service=%s", service.Name)}, expiresMinutes)
}

func (e EtcdClientImpl) GetPolicy(data interface{}) *entity.Policy {
	if e.Connection == nil {
		return nil
	}
	var policy entity.Policy
	err := e.get(fmt.Sprintf("policy=%s", data), policy)
	if err != nil {
		log.Logger.Info("Cloud not get cache", err.Error())
		return nil
	}
	return &policy
}
func (e EtcdClientImpl) GetPolicyByNames(names []string) []entity.Policy {
	if e.Connection == nil {
		return nil
	}
	var policies []entity.Policy
	for _, name := range names {
		var policy entity.Policy
		err := e.get(fmt.Sprintf("policy=%s", name), policy)
		if err != nil {
			log.Logger.Info("Cloud not get cache", err.Error())
			continue
		}
		policies = append(policies, policy)
	}
	return policies
}

func (e EtcdClientImpl) GetPermission(data interface{}) *entity.Permission {
	if e.Connection == nil {
		return nil
	}
	var permission entity.Permission
	err := e.get(fmt.Sprintf("permission=%s", data), permission)
	if err != nil {
		log.Logger.Info("Cloud not get cache", err.Error())
		return nil
	}
	return &permission
}

func (e EtcdClientImpl) GetPermissionByNames(names []string) []entity.Permission {
	if e.Connection == nil {
		return nil
	}
	var permissions []entity.Permission
	for _, name := range names {
		var permission entity.Permission
		err := e.get(fmt.Sprintf("permission=%s", name), permission)
		if err != nil {
			log.Logger.Info("Cloud not get cache", err.Error())
			continue
		}
		permissions = append(permissions, permission)
	}
	return permissions
}

func (e EtcdClientImpl) GetRole(data interface{}) *entity.Role {
	if e.Connection == nil {
		return nil
	}
	var role entity.Role
	err := e.get(fmt.Sprintf("role=%s", data), role)
	if err != nil {
		log.Logger.Info("Cloud not get cache", err.Error())
		return nil
	}
	return &role
}

func (e EtcdClientImpl) GetRoleByNames(names []string) []entity.Role {
	if e.Connection == nil {
		return nil
	}
	var roles []entity.Role
	for _, name := range names {
		var role entity.Role
		err := e.get(fmt.Sprintf("role=%s", name), role)
		if err != nil {
			log.Logger.Info("Cloud not get cache", err.Error())
			continue
		}
		roles = append(roles, role)
	}
	return roles
}

func (e EtcdClientImpl) GetService(data interface{}) *entity.Service {
	if e.Connection == nil {
		return nil
	}
	var service entity.Service
	err := e.get(fmt.Sprintf("service=%s", data), service)
	if err != nil {
		log.Logger.Info("Cloud not get cache", err.Error())
		return nil
	}
	return &service
}

func (e EtcdClientImpl) GetServiceByNames(names []string) []entity.Service {
	if e.Connection == nil {
		return nil
	}
	var services []entity.Service
	for _, name := range names {
		var service entity.Service
		err := e.get(fmt.Sprintf("service=%s", name), service)
		if err != nil {
			log.Logger.Info("Cloud not get cache", err.Error())
			continue
		}
		services = append(services, service)
	}
	return services
}

// Get cache shared method
func (e EtcdClientImpl) get(key string, structData interface{}) error {
	start := time.Millisecond
	response, err := e.Connection.Get(ctx.GetCtx(), key)
	if err != nil {
		return errors.New(fmt.Sprintf("Cache is null. key = %s", key))
	}
	end := time.Millisecond
	kvs := response.Kvs[0]
	log.Logger.Info(fmt.Sprintf("[%sms] GET etcd key %s", end - start, key))
	err = json.Unmarshal(kvs.Value, &structData)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to convert json to struct for cache. %s", err.Error()))
	}
	return nil
}

// Set cache shared method
func (e EtcdClientImpl) set(json []byte, keys []string, expiresMinutes time.Duration) {
	for _, key := range keys {
		start := time.Millisecond
		_, err := e.Connection.Put(ctx.GetCtx(), key, string(json))
		if err != nil {
			log.Logger.Error(fmt.Sprintf("Cache is null. key = %s", key))
		}
		end := time.Millisecond
		log.Logger.Info(fmt.Sprintf("[%s, %sms] PUT etcd key %s", end, start, key))
	}
}
