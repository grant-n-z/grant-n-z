package service

import (
	"strconv"
	"strings"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"

	"github.com/tomoyane/grant-n-z/server/usecase/repository"
)

type userServiceServiceImpl struct {
	userServiceRepository repository.UserServiceRepository
	userRepository        repository.UserRepository
	serviceRepository     repository.ServiceRepository
}

func NewUserServiceService() UserServiceService {
	log.Logger.Info("Inject `userServiceRepository`, `userRepository`, `serviceRepository` to `UserServiceService`")
	return userServiceServiceImpl{
		userServiceRepository: repository.UserServiceRepositoryImpl{Db: config.Db},
		userRepository:        repository.UserRepositoryImpl{Db: config.Db},
		serviceRepository:     repository.ServiceRepositoryImpl{Db: config.Db},
	}
}

func (uss userServiceServiceImpl) Get(queryParam string) (interface{}, *model.ErrorResponse) {
	var result interface{}

	if strings.EqualFold(queryParam, "") {
		return uss.GetUserServices()
	}

	i, castErr := strconv.Atoi(queryParam)
	if castErr != nil {
		log.Logger.Warn("The user_id is only integer")
		return nil, model.BadRequest(castErr.Error())
	}

	userServiceEntities, err := uss.GetUserServicesByUserId(i)
	if err != nil {
		return nil, err
	}

	if userServiceEntities == nil {
		result = new([]string)
	} else {
		result = userServiceEntities
	}

	return result, nil
}

func (uss userServiceServiceImpl) GetUserServices() ([]*entity.UserService, *model.ErrorResponse) {
	return uss.userServiceRepository.FindAll()
}

func (uss userServiceServiceImpl) GetUserServicesByUserId(userId int) ([]*entity.UserService, *model.ErrorResponse) {
	return uss.userServiceRepository.FindByUserId(userId)
}

func (uss userServiceServiceImpl) InsertUserService(userService *entity.UserService) (*entity.UserService, *model.ErrorResponse) {
	if userEntity, _ := uss.userRepository.FindById(userService.UserId); userEntity == nil {
		log.Logger.Warn("Not found user id")
		return nil, model.BadRequest("Not found user id")
	}

	if serviceEntity, _ := uss.serviceRepository.FindById(userService.ServiceId); serviceEntity == nil {
		log.Logger.Warn("Not found service id")
		return nil, model.BadRequest("Not found service id")
	}

	return uss.userServiceRepository.Save(*userService)
}
