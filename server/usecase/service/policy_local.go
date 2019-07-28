package service

type PolicyLocalService interface {
	ReadPolicy(basePath string)

	WritePolicy(basePath string)

	EncryptData(data string) (*string, error)

	DecryptData(data string) (*string, error)
}
