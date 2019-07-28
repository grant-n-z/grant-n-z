package service

import (
	"fmt"
	"os"

	"crypto/rand"
	"crypto/rsa"

	"github.com/tomoyane/grant-n-z/server/log"
)

const (
	BitSize = 2048
)

var (
	PrivateKey *rsa.PrivateKey = nil
	PublicKey  *rsa.PublicKey = nil
)

type policyLocalServiceImpl struct {
}

func NewPolicyLocalService() PolicyLocalService {
	log.Logger.Info("")
	return policyLocalServiceImpl{}
}

func (plsi policyLocalServiceImpl) ReadPolicy(basePath string) {
	panic("implement me")
}

func (plsi policyLocalServiceImpl) WritePolicy(basePath string) {
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

func (plsi policyLocalServiceImpl) EncryptData(payload string) (*string, error) {
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

func (plsi policyLocalServiceImpl) DecryptData(data string) (*string, error) {
	decryptedJsonBytes, err := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, []byte(data))
	if err != nil {
		log.Logger.Error("Error decrypt PKCS1v15", err.Error())
		return nil, err
	}

	decryptedPayload := string(decryptedJsonBytes)
	return &decryptedPayload, nil
}
