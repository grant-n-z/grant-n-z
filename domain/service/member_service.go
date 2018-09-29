package service

import (
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
)

type MemberService struct {
	MemberRepository    repository.MemberRepository
}

func (s MemberService) GetByUserUuidAndServiceUuid(userUuid uuid.UUID, serviceUuid uuid.UUID) *entity.Member {
	return s.MemberRepository.FindByUserUuidAndServiceUuid(userUuid, serviceUuid)
}

func (s MemberService) Insert(member entity.Member) *entity.Member {
	return s.MemberRepository.Save(member)
}