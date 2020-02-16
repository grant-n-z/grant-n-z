package service

type UpdaterService interface {

}

type UpdaterServiceImpl struct {

}

func NewUpdaterService() UpdaterService {
	return UpdaterServiceImpl{}
}

