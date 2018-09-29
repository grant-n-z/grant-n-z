package service

import (
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/tomoyane/grant-n-z/handler"
)

type PrincipalService struct {
	PrincipalRepository repository.PrincipalRepository
	UserRepository      repository.UserRepository
	ServiceRepository   repository.ServiceRepository
	MemberRepository    repository.MemberRepository
	RoleRepository      repository.RoleRepository
}

func (s PrincipalService) GetPrincipalByName(name string) *entity.Principal {
	return s.PrincipalRepository.FindByName(name)
}

func (s PrincipalService) InsertPrincipal(principal entity.Principal) *entity.Principal {
	return s.PrincipalRepository.Save(principal)
}

func (s PrincipalService) PostPrincipalData(principalRequest entity.PrincipalRequest) (
	insertedPrincipal *entity.Principal, errRes *handler.ErrorResponse) {

	user := s.UserRepository.FindByUserName(principalRequest.UserName)
	if user == nil {
		return nil, handler.InternalServerError("")
	}

	if len(user.Uuid) == 0 {
		return nil, handler.NotFound("")
	}

	service := s.ServiceRepository.FindByName(principalRequest.ServiceName)
	if service == nil {
		return nil, handler.InternalServerError("")
	}

	if len(service.Uuid) == 0 {
		return nil, handler.NotFound("")
	}

	role := s.RoleRepository.FindByPermission(principalRequest.RolePermission)
	if role == nil {
		return nil, handler.InternalServerError("")
	}

	if len(role.Uuid) == 0 {
		return nil, handler.NotFound("")
	}

	member := s.MemberRepository.FindByUserUuidAndServiceUuid(user.Uuid, service.Uuid)
	if member == nil {
		return nil, handler.InternalServerError("")
	}

	if len(member.Uuid) == 0 {
		return nil, handler.NotFound("")
	}

	principalUuid, _ := uuid.NewV4()
	principal := entity.Principal{
		Uuid:       principalUuid,
		MemberUuid: member.Uuid,
		RoleUuid:   role.Uuid,
	}

	principalData := s.InsertPrincipal(principal)
	if principalData == nil {
		return nil, handler.InternalServerError("")
	}

	return principalData, nil
}
