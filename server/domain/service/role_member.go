package service

import (
	"strconv"
	"strings"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/repository"
	"github.com/tomoyane/grant-n-z/server/log"
)

type RoleMemberService struct {
	RoleMemberRepository repository.RoleMemberRepository
	UserRepository       repository.UserRepository
	RoleRepository       repository.RoleRepository
}

func NewRoleMemberService() RoleMemberService {
	log.Logger.Info("inject `RoleMemberRepository`,`UserRepository`,`RoleRepository` to `RoleMemberService`")
	return RoleMemberService{
		RoleMemberRepository: repository.RoleMemberRepositoryImpl{},
		UserRepository:       repository.UserRepositoryImpl{},
		RoleRepository:       repository.RoleRepositoryImpl{},
	}
}

func (rms RoleMemberService) Get(queryParam string) (interface{}, *entity.ErrorResponse) {
	var result interface{}

	if strings.EqualFold(queryParam, "") {
		return rms.GetRoleMembers()
	}

	i, castErr := strconv.Atoi(queryParam)
	if castErr != nil {
		log.Logger.Warn("user_id is only integer")
		return nil, entity.BadRequest(castErr.Error())
	}

	roleMemberEntities, err := rms.GetRoleMemberByUserId(i)
	if err != nil {
		return nil, err
	}

	if roleMemberEntities == nil {
		result = []string{}
	} else {
		result = roleMemberEntities
	}

	return result, nil
}

func (rms RoleMemberService) GetRoleMembers() ([]*entity.RoleMember, *entity.ErrorResponse) {
	return rms.RoleMemberRepository.FindAll()
}

func (rms RoleMemberService) GetRoleMemberByUserId(userId int) ([]*entity.RoleMember, *entity.ErrorResponse) {
	return rms.RoleMemberRepository.FindByUserId(userId)
}

func (rms RoleMemberService) InsertRoleMember(roleMember *entity.RoleMember) (*entity.RoleMember, *entity.ErrorResponse) {
	if userEntity, _ := rms.UserRepository.FindById(roleMember.UserId); userEntity == nil {
		return nil, entity.BadRequest("not found user id")
	}

	if roleEntity, _ := rms.RoleRepository.FindById(roleMember.RoleId); roleEntity == nil {
		return nil, entity.BadRequest("not found role id")
	}

	return rms.RoleMemberRepository.Save(*roleMember)
}
