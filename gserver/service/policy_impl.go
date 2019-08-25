package service

import (
	"strconv"
	"strings"

	"crypto/rand"
	"crypto/rsa"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"

	"github.com/tomoyane/grant-n-z/gserver/repository"
)

const BitSize = 2048

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	plsInstance PolicyService
)

type policyServiceImpl struct {
	policyRepository            repository.PolicyRepository
	permissionRepository        repository.PermissionRepository
	roleRepository              repository.RoleRepository
	serviceMemberRoleRepository repository.ServiceMemberRoleRepository
}

func GetPolicyServiceInstance() PolicyService {
	if plsInstance == nil {
		plsInstance = NewPolicyService()
	}
	return plsInstance
}

func NewPolicyService() PolicyService {
	log.Logger.Info("New `RequestHandler` instance")
	log.Logger.Info("Inject `PolicyRepository`, `PermissionRepository`, `RoleRepository`, `ServiceMemberRoleRepository` to `PolicyService`")
	return policyServiceImpl{
		policyRepository:            repository.GetPolicyRepositoryInstance(driver.Db),
		permissionRepository:        repository.NewPermissionRepository(driver.Db),
		roleRepository:              repository.GetRoleRepositoryInstance(driver.Db),
		serviceMemberRoleRepository: repository.GetServiceMemberRoleRepositoryInstance(driver.Db),
	}
}

func (ps policyServiceImpl) Get(queryParam string) ([]*entity.Policy, *model.ErrorResponse) {
	if strings.EqualFold(queryParam, "") {
		return ps.GetPolicies()
	}

	i, castErr := strconv.Atoi(queryParam)
	if castErr != nil {
		log.Logger.Warn("The role_id is only integer")
		return nil, model.BadRequest(castErr.Error())
	}

	policyEntities, err := ps.GetPoliciesByRoleId(i)
	if err != nil {
		return nil, err
	}

	if policyEntities == nil {
		return []*entity.Policy{}, nil
	}

	return policyEntities, nil
}

func (ps policyServiceImpl) GetPolicies() ([]*entity.Policy, *model.ErrorResponse) {
	return ps.policyRepository.FindAll()
}

func (ps policyServiceImpl) GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResponse) {
	return ps.policyRepository.FindByRoleId(roleId)
}

func (ps policyServiceImpl) InsertPolicy(policy *entity.Policy) (*entity.Policy, *model.ErrorResponse) {
	if permissionEntity, _ := ps.permissionRepository.FindById(policy.PermissionId); permissionEntity == nil {
		log.Logger.Warn("Not found permission id")
		return nil, model.BadRequest("Not found permission id")
	}

	if roleEntity, _ := ps.serviceMemberRoleRepository.FindById(policy.ServiceMemberRoleId); roleEntity == nil {
		log.Logger.Warn("Not found service member role id")
		return nil, model.BadRequest("Not found service member role id")
	}

	return ps.policyRepository.Save(*policy)
}

func (ps policyServiceImpl) EncryptData(payload string) (*string, error) {
	if PrivateKey == nil {
		generatedPri, err := rsa.GenerateKey(rand.Reader, BitSize)
		if err != nil {
			log.Logger.Error("Error generate private key", err.Error())
			return nil, err
		}
		PrivateKey = generatedPri
	}

	if PublicKey == nil {
		generatedPub := &PrivateKey.PublicKey
		PublicKey = generatedPub
	}

	cipherJsonBytes, err := rsa.EncryptPKCS1v15(rand.Reader, PublicKey, []byte(payload))
	if err != nil {
		log.Logger.Error("Error encrypt PKCS1v15", err.Error())
		return nil, err
	}

	cipherPayload := string(cipherJsonBytes)
	return &cipherPayload, nil
}

func (ps policyServiceImpl) DecryptData(data string) (*string, error) {
	decryptedJsonBytes, err := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, []byte(data))
	if err != nil {
		log.Logger.Error("Error decrypt PKCS1v15", err.Error())
		return nil, err
	}

	decryptedPayload := string(decryptedJsonBytes)
	return &decryptedPayload, nil
}
