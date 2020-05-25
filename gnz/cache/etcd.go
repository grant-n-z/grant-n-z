package cache

import (
	"context"
	"errors"
	"fmt"

	"encoding/json"

	"go.etcd.io/etcd/clientv3"

	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

const retryCnt = 5

var eInstance EtcdClient

type EtcdClient interface {
	// Set permission with expires
	// key: permission={uuid}
	// value: {"name":"{name}"}
	SetPermission(permissionUuid string, permission structure.Permission)

	// Set role with expires
	// key: role={uuid}
	// value: {"name":"{name}"}
	SetRole(roleUuid string, role structure.Role)

	// Set service with expires
	// key: service={uuid}
	// value: {"name":"{name}"}
	SetService(serviceUuid string, service structure.Service)

	// Set policy with expires
	// key: user_policy={user_uuid}
	// value: [{"service_uuid":"{uuid}","group_uuid":"{uuid}","role_name":"{name}","permission_name":"{name}"}]
	SetUserPolicy(userUuid string, policy []structure.UserPolicy)

	// Set user_service with expires
	// key: user_service={user_uuid}
	// value: [{"service_name":"{name}","service_uuid":"{uuid}"}]
	SetUserService(userUuid string, userServices []structure.UserService)

	// Set user_group with expires
	// key: user_group={user_uuid}
	// value: [{"group_name":"{name}","group_uuid":"{uuid}"}]
	SetUserGroup(userUuid string, userGroups []structure.UserGroup)

	// Get policy by user uuid
	GetUserPolicy(userUuid string) []structure.UserPolicy

	// Get permission by uuid
	GetPermission(permissionUuid string) *structure.Permission

	// Get role by uuid
	GetRole(roleUuid string) *structure.Role

	// Get service by uuid
	GetService(serviceUuid string) *structure.Service

	// Get user_service by user uuid
	GetUserService(userUuid string) []structure.UserService

	// Get user_group by user uuid
	GetUserGroup(userUuid string) []structure.UserGroup

	// Delete policy by user uuid
	DeleteUserPolicy(userUuid string)
}

type EtcdClientImpl struct {
	Connection *clientv3.Client
	Ctx        context.Context
}

func GetEtcdClientInstance() EtcdClient {
	if eInstance == nil {
		eInstance = NewEtcdClient()
	}
	return eInstance
}

// Constructor
// Need to initial ctx.InitContext method
func NewEtcdClient() EtcdClient {
	log.Logger.Info("New `EtcdClient` instance")
	return EtcdClientImpl{
		Connection: connection,
		Ctx:        context.Background(),
	}
}

func (e EtcdClientImpl) SetUserPolicy(userUuid string, policy []structure.UserPolicy) {
	policyJson, _ := json.Marshal(policy)
	e.set([]string{fmt.Sprintf("user_policy=%s", userUuid)}, policyJson)
}

func (e EtcdClientImpl) SetPermission(permissionUuid string, permission structure.Permission) {
	permissionJson, _ := json.Marshal(permission)
	e.set([]string{fmt.Sprintf("permission=%s", permissionUuid)}, permissionJson)
}

func (e EtcdClientImpl) SetRole(roleUuid string, role structure.Role) {
	roleJson, _ := json.Marshal(role)
	e.set([]string{fmt.Sprintf("role=%s", roleUuid)}, roleJson)
}

func (e EtcdClientImpl) SetService(serviceUuid string, service structure.Service) {
	serviceJson, _ := json.Marshal(service)
	e.set([]string{fmt.Sprintf("service=%s", serviceUuid)}, serviceJson)
}

func (e EtcdClientImpl) SetUserService(userUuid string, userServices []structure.UserService) {
	userServiceJson, _ := json.Marshal(userServices)
	e.set([]string{fmt.Sprintf("user_service=%s", userUuid)}, userServiceJson)
}

func (e EtcdClientImpl) SetUserGroup(userUuid string, userGroups []structure.UserGroup) {
	userGroupJson, _ := json.Marshal(userGroups)
	e.set([]string{fmt.Sprintf("user_group=%s", userUuid)}, userGroupJson)
}

func (e EtcdClientImpl) GetUserPolicy(userUuid string) []structure.UserPolicy {
	var policy []structure.UserPolicy
	err := e.get(fmt.Sprintf("user_policy=%s", userUuid), &policy)
	if err != nil {
		return nil
	}
	return policy
}

func (e EtcdClientImpl) GetPermission(permissionUuid string) *structure.Permission {
	var permission structure.Permission
	err := e.get(fmt.Sprintf("permission=%s", permissionUuid), &permission)
	if err != nil {
		return nil
	}
	return &permission
}

func (e EtcdClientImpl) GetRole(roleUuid string) *structure.Role {
	var role structure.Role
	err := e.get(fmt.Sprintf("role=%s", roleUuid), &role)
	if err != nil {
		return nil
	}
	return &role
}

func (e EtcdClientImpl) GetService(serviceUuid string) *structure.Service {
	var service structure.Service
	err := e.get(fmt.Sprintf("service=%s", serviceUuid), &service)
	if err != nil {
		return nil
	}
	return &service
}

func (e EtcdClientImpl) GetUserService(userUuid string) []structure.UserService {
	var userServices []structure.UserService
	err := e.get(fmt.Sprintf("user_service=%s", userUuid), &userServices)
	if err != nil {
		return nil
	}
	return userServices
}

func (e EtcdClientImpl) GetUserGroup(userUuid string) []structure.UserGroup {
	var userGroups []structure.UserGroup
	err := e.get(fmt.Sprintf("user_group=%s", userUuid), &userGroups)
	if err != nil {
		return nil
	}
	return userGroups
}

func (e EtcdClientImpl) DeleteUserPolicy(userUuid string) {
	e.delete([]string{fmt.Sprintf("user_policy=%s", userUuid)})
}

// Get cache shared method
func (e EtcdClientImpl) get(key string, structData interface{}) error {
	if e.Connection == nil {
		return errors.New("Not connected etcd")
	}
	response, err := e.Connection.Get(e.Ctx, key)
	if err != nil || len(response.Kvs) == 0 {
		return errors.New(fmt.Sprintf("Cache is null. key = %v", key))
	}
	kvs := response.Kvs[0]
	err = json.Unmarshal(kvs.Value, &structData)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to convert json to struct for cache. %v", err.Error()))
	}
	return nil
}

// Set cache shared method
func (e EtcdClientImpl) set(keys []string, json []byte) {
	if e.Connection == nil {
		return
	}
	for _, key := range keys {
		_, err := e.Connection.Put(e.Ctx, key, string(json))
		if err != nil {
			log.Logger.Error(fmt.Sprintf("Failed to put data. key = %v. err = %s", key, err.Error()))
		}
	}
}

// Delete cache shared method
func (e EtcdClientImpl) delete(keys []string) {
	if e.Connection == nil {
		return
	}
	for _, key := range keys {
		for i := 0; i < retryCnt; i++ {
			_, err := e.Connection.Delete(e.Ctx, key)
			if err != nil {
				fmt.Println(err)
				log.Logger.Error(fmt.Sprintf("Failed to delete data. key = %v. err = %s", key, err.Error()))
			} else {
				break
			}
		}
	}
}
