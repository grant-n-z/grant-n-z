package service

import (
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/repository"
)

type RoleMemberService struct {
	RoleMemberRepository repository.RoleMemberRepository
}

func NewRoleMemberService() RoleMemberService {
	return RoleMemberService{RoleMemberRepository: repository.RoleMemberRepositoryImpl{}}
}

func (rms RoleMemberService) GetRoleMembers() ([]*entity.RoleMember, *entity.ErrorResponse)  {
	return rms.RoleMemberRepository.FindAll()
}

func (rms RoleMemberService) GetRoleMemberByUserId(userId int) ([]*entity.RoleMember, *entity.ErrorResponse)  {
	return rms.RoleMemberRepository.FindByUserId(userId)
}

func (rms RoleMemberService) InsertRoleMember(roleMember *entity.RoleMember) (*entity.RoleMember, *entity.ErrorResponse) {
	return rms.RoleMemberRepository.Save(*roleMember)
}
