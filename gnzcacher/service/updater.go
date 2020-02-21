package service

type UpdaterService interface {
	UpdatePolicy() error

	UpdatePermission() error

	UpdateRole() error

	UpdateService() error
}

type UpdaterServiceImpl struct {

}

func NewUpdaterService() UpdaterService {
	return UpdaterServiceImpl{}
}

func (us UpdaterServiceImpl) UpdatePolicy() error {
	panic("implement me")
}

func (us UpdaterServiceImpl) UpdatePermission() error {
	panic("implement me")
}

func (us UpdaterServiceImpl) UpdateRole() error {
	panic("implement me")
}

func (us UpdaterServiceImpl) UpdateService() error {
	panic("implement me")
}
