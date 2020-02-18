package service

type ExtractorService interface {
	GetPolicy()

	GetPermission()

	GetRole()

	GetService()
}

type ExtractorServiceImpl struct {
}

func NewExtractorService() ExtractorService {
	return ExtractorServiceImpl{}
}

func (us ExtractorServiceImpl) GetPolicy() {
	panic("implement me")
}

func (us ExtractorServiceImpl) GetPermission() {
	panic("implement me")
}

func (us ExtractorServiceImpl) GetRole() {
	panic("implement me")
}

func (us ExtractorServiceImpl) GetService() {
	panic("implement me")
}
