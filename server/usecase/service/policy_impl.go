package service

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"crypto/rand"
	"crypto/rsa"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"

	"github.com/tomoyane/grant-n-z/server/usecase/repository"
)

const BitSize = 2048

var (
	PrivateKey *rsa.PrivateKey = nil
	PublicKey  *rsa.PublicKey = nil
)

type policyServiceImpl struct {
	policyRepository     repository.PolicyRepository
	permissionRepository repository.PermissionRepository
	roleRepository       repository.RoleRepository
}

func NewPolicyService() PolicyService {
	log.Logger.Info("Inject `roleRepository`, `permissionRepository`, `roleRepository` to `PolicyService`")
	return policyServiceImpl{
		policyRepository:     repository.PolicyRepositoryImpl{Db: config.Db},
		permissionRepository: repository.PermissionRepositoryImpl{Db: config.Db},
		roleRepository:       repository.RoleRepositoryImpl{Db: config.Db},
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

	if roleEntity, _ := ps.roleRepository.FindById(policy.RoleId); roleEntity == nil {
		log.Logger.Warn("Not found role id")
		return nil, model.BadRequest("Not found role id")
	}

	return ps.policyRepository.Save(*policy)
}

func (ps policyServiceImpl) ReadLocalPolicy(basePath string) {
	panic("implement me")
}

func (ps policyServiceImpl) WriteLocalPolicy(basePath string) {
	path := fmt.Sprintf("%spolicy.json", basePath)
	file, err := os.Open(path)
	if err != nil {
		file, err = os.Create(path)
		if err != nil {
			log.Logger.Error("Error write policy file", err.Error())
		}
	}
	defer file.Close()

	// TODO: Read policy table, then update policy file
	// TODO: Now, example test data
	output := "{'key': 'value'}"
	_, _ = file.Write(([]byte)(output))
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
