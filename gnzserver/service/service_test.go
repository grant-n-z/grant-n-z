package service

import (
	"strings"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	service Service
)

// Set up
func init() {
	log.InitLogger("info")
	ctx.InitContext()
	ctx.SetUserId(1)
	ctx.SetServiceId(1)
	ctx.SetApiKey("test_key")

	stubConnection, _ = gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	service = ServiceImpl{
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
			Ctx:        ctx.GetCtx(),
		},
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
	_, err := service.GetServiceById(1)
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
	ctx.SetApiKey("test_key")
	_, err := service.GetServiceOfApiKey()
	if err != nil {
		t.Errorf("Incorrect TestGetServiceOfApiKey_Success test")
		t.FailNow()
	}
}

// Test get service of user
func TestGetServiceOfUser_Success(t *testing.T) {
	_, err := service.GetServiceOfUser()
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

// Test generate api key
func TestGenerateApiKey(t *testing.T) {
	key := service.GenerateApiKey()
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

func (sri StubServiceRepositoryImpl) FindAll() ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) FindById(id int) (*entity.Service, *model.ErrorResBody) {
	var service entity.Service
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindByName(name string) (*entity.Service, *model.ErrorResBody) {
	var service entity.Service
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindByApiKey(apiKey string) (*entity.Service, *model.ErrorResBody) {
	service := entity.Service{Name:"test"}
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindNameById(id int) *string {
	service, _ := sri.FindById(id)
	return &service.Name
}

func (sri StubServiceRepositoryImpl) FindNameByApiKey(name string) *string {
	service, _ := sri.FindByName(name)
	return &service.Name
}

func (sri StubServiceRepositoryImpl) FindServicesByUserId(userId int) ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) Save(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	return &service, nil
}

func (sri StubServiceRepositoryImpl) SaveWithRelationalData(service entity.Service, roles []entity.Role, permissions []entity.Permission) (*entity.Service, *model.ErrorResBody) {
	return &service, nil
}

func (sri StubServiceRepositoryImpl) Update(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	return &service, nil
}
