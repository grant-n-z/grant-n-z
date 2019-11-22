package service

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

const BitSize = 2048

var (
	PrivateKey  *rsa.PrivateKey
	PublicKey   *rsa.PublicKey
	plsInstance PolicyService
)

type PolicyService interface {
	GetPolicies() ([]*entity.Policy, *model.ErrorResBody)

	GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody)

	GetPolicyByOfUser() ([]map[string]entity.PolicyResponse, *model.ErrorResBody)

	InsertPolicy(policy *entity.Policy) (*entity.Policy, *model.ErrorResBody)

	EncryptData(data string) (*string, error)

	DecryptData(data string) (*string, error)
}

type policyServiceImpl struct {
	policyRepository     data.PolicyRepository
	permissionRepository data.PermissionRepository
	roleRepository       data.RoleRepository
	userGroupRepository  data.UserGroupRepository
}

func GetPolicyServiceInstance() PolicyService {
	if plsInstance == nil {
		plsInstance = NewPolicyService()
	}
	return plsInstance
}

func NewPolicyService() PolicyService {
	log.Logger.Info("New `PolicyService` instance")
	log.Logger.Info("Inject `PolicyRepository`, `PermissionRepository`, `RoleRepository`, `UserGroupRepository` to `PolicyService`")
	return policyServiceImpl{
		policyRepository:     data.GetPolicyRepositoryInstance(driver.Db),
		permissionRepository: data.GetPermissionRepositoryInstance(driver.Db),
		roleRepository:       data.GetRoleRepositoryInstance(driver.Db),
		userGroupRepository:  data.GetUserGroupRepositoryInstance(driver.Db),
	}
}

func (ps policyServiceImpl) GetPolicies() ([]*entity.Policy, *model.ErrorResBody) {
	return ps.policyRepository.FindAll()
}

func (ps policyServiceImpl) GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody) {
	return ps.policyRepository.FindByRoleId(roleId)
}

func (ps policyServiceImpl) GetPolicyByOfUser() ([]map[string]entity.PolicyResponse, *model.ErrorResBody) {
	if ctx.GetUserId().(int) == 0 {
		return nil, model.BadRequest("Required user id")
	}

	joinObj, err := ps.userGroupRepository.FindGroupWithUserWithPolicyGroupsByUserId(ctx.GetUserId().(int))
	if err != nil {
		return nil, err
	}

	var groupPolicyMaps []map[string]entity.PolicyResponse
	for _, joinData := range joinObj {
		// TODO: Read redis cache, roles and permissions
		policy := entity.NewPolicyResponse().
			SetName(&joinData.Policy.Name).
			SetRoleName(ps.roleRepository.FindNameById(joinData.Policy.RoleId)).
			SetPermissionName(ps.permissionRepository.FindNameById(joinData.Policy.PermissionId)).
			Build()

		response := map[string]entity.PolicyResponse{joinData.Group.Name: policy}
		groupPolicyMaps = append(groupPolicyMaps, response)
	}

	return groupPolicyMaps, nil
}

func (ps policyServiceImpl) InsertPolicy(policy *entity.Policy) (*entity.Policy, *model.ErrorResBody) {
	if permissionEntity, _ := ps.permissionRepository.FindById(policy.PermissionId); permissionEntity == nil {
		log.Logger.Warn("Not found permission id")
		return nil, model.BadRequest("Not found permission id")
	}

	return ps.policyRepository.Save(*policy)
}

func (ps policyServiceImpl) EncryptData(payload string) (*string, error) {
	if PrivateKey == nil {
		generatedPri, err := rsa.GenerateKey(rand.Reader, BitSize)
		if err != nil {
			log.Logger.Error("Failed to generateSignedInToken private key", err.Error())
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
		log.Logger.Error("Failed to encrypt PKCS1v15", err.Error())
		return nil, err
	}

	cipherPayload := string(cipherJsonBytes)
	return &cipherPayload, nil
}

func (ps policyServiceImpl) DecryptData(data string) (*string, error) {
	decryptedJsonBytes, err := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, []byte(data))
	if err != nil {
		log.Logger.Error("Failed to decrypt PKCS1v15", err.Error())
		return nil, err
	}

	decryptedPayload := string(decryptedJsonBytes)
	return &decryptedPayload, nil
}
