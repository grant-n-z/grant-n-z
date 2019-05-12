package service

import (
	"strconv"
	"strings"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/usecase/repository"
)

type roleMemberServiceImpl struct {
	roleMemberRepository repository.RoleMemberRepository
	userRepository       repository.UserRepository
	roleRepository       repository.RoleRepository
}

func NewRoleMemberService() RoleMemberService {
	log.Logger.Info("Inject `roleMemberRepository`,`userRepository`,`roleRepository` to `RoleMemberService`")
	return roleMemberServiceImpl{
		roleMemberRepository: repository.RoleMemberRepositoryImpl{Db: config.Db},
		userRepository:       repository.UserRepositoryImpl{Db: config.Db},
		roleRepository:       repository.RoleRepositoryImpl{Db: config.Db},
	}
}

func (rms roleMemberServiceImpl) Get(queryParam string) ([]*entity.RoleMember, *entity.ErrorResponse) {
	if strings.EqualFold(queryParam, "") {
		return rms.GetRoleMembers()
	}

	i, castErr := strconv.Atoi(queryParam)
	if castErr != nil {
		log.Logger.Warn("The user_id is only integer")
		return nil, entity.BadRequest(castErr.Error())
	}

	roleMemberEntities, err := rms.GetRoleMemberByUserId(i)
	if err != nil {
		return nil, err
	}

	if roleMemberEntities == nil {
		return []*entity.RoleMember{}, nil
	}

	return roleMemberEntities, nil
}

func (rms roleMemberServiceImpl) GetRoleMembers() ([]*entity.RoleMember, *entity.ErrorResponse) {
	return rms.roleMemberRepository.FindAll()
}

func (rms roleMemberServiceImpl) GetRoleMemberByUserId(userId int) ([]*entity.RoleMember, *entity.ErrorResponse) {
	return rms.roleMemberRepository.FindByUserId(userId)
}

func (rms roleMemberServiceImpl) InsertRoleMember(roleMember *entity.RoleMember) (*entity.RoleMember, *entity.ErrorResponse) {
	if userEntity, _ := rms.userRepository.FindById(roleMember.UserId); userEntity == nil {
		return nil, entity.BadRequest("Not found user id")
	}

	if roleEntity, _ := rms.roleRepository.FindById(roleMember.RoleId); roleEntity == nil {
		return nil, entity.BadRequest("Not found role id")
	}

	return rms.roleMemberRepository.Save(*roleMember)
}
