package service

import (
	"strings"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var (
	service Service
)

// Set up
func init() {
	log.InitLogger("info")

	stubConnection, _ = gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	service = ServiceImpl{
		EtcdClient:           cache.EtcdClientImpl{Connection: stubEtcdConnection},
		ServiceRepository:    StubServiceRepositoryImpl{Connection: stubConnection},
		RoleRepository:       StubRoleRepositoryImpl{Connection: stubConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}
}

// Test constructor
func TestGetServiceInstance(t *testing.T) {
	GetServiceInstance()
}

// Test get services
func TestGetServices_Success(t *testing.T) {
	_, err := service.GetServices()
	if err != nil {
		t.Errorf("Incorrect TestGetServices_Success test")
		t.FailNow()
	}
}

// Test get service by id
func TestGetServiceById_Success(t *testing.T) {
	_, err := service.GetServiceByUuid(uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetServiceById_Success test")
		t.FailNow()
	}
}

// Test get service by name
func TestGetServiceByName_Success(t *testing.T) {
	_, err := service.GetServiceByName("name")
	if err != nil {
		t.Errorf("Incorrect TestGetServiceByName_Success test")
		t.FailNow()
	}
}

// Test get service of user
func TestGetServiceOfApiKey_Success(t *testing.T) {
	_, err := service.GetServiceBySecret("secret")
	if err != nil {
		t.Errorf("Incorrect TestGetServiceOfApiKey_Success test")
		t.FailNow()
	}
}

// Test get service of user
func TestGetServiceOfUser_Success(t *testing.T) {
	_, err := service.GetServiceByUser(uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetServiceOfUser_Success test")
		t.FailNow()
	}
}

// Test insert service
func TestInsertService_Success(t *testing.T) {
	_, err := service.InsertService(entity.Service{Name: "name"})
	if err != nil {
		t.Errorf("Incorrect TestInsertService_Success test")
		t.FailNow()
	}
}

// Test insert service with relational data
func TestInsertServiceWithRelationalData_Success(t *testing.T) {
	_, err := service.InsertServiceWithRelationalData(&entity.Service{Name: "name"})
	if err != nil {
		t.Errorf("Incorrect TestInsertServiceWithRelationalData_Success test")
		t.FailNow()
	}
}

// Test generate secret
func TestGenerateApiKey(t *testing.T) {
	key := service.GenerateSecret()
	if strings.EqualFold(key, "") {
		t.Errorf("Incorrect TestGenerateApiKey test")
		t.FailNow()
	}
}

// Less than stub struct
// Service repository
type StubServiceRepositoryImpl struct {
	Connection *gorm.DB
}

func (sri StubServiceRepositoryImpl) FindAll() ([]*entity.Service, error) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Service, error) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) FindByUuid(uuid string) (*entity.Service, error) {
	var service entity.Service
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindByName(name string) (*entity.Service, error) {
	var service entity.Service
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindBySecret(secret string) (*entity.Service, error) {
	service := entity.Service{Name: "test"}
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindNameByUuid(uuid string) *string {
	service, err := sri.FindByUuid(uuid)
	if err != nil {
		return nil
	}
	return &service.Name
}

func (sri StubServiceRepositoryImpl) FindServicesByUserUuid(userUuid string) ([]*entity.Service, error) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) Save(service entity.Service) (*entity.Service, error) {
	return &service, nil
}

func (sri StubServiceRepositoryImpl) SaveWithRelationalData(service entity.Service, roles []entity.Role, permissions []entity.Permission) (*entity.Service, error) {
	return &service, nil
}

func (sri StubServiceRepositoryImpl) Update(service entity.Service) (*entity.Service, error) {
	return &service, nil
}
