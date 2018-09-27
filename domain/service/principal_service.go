package service

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/handler"
)

type PrincipalService struct {
	PrincipalRepository repository.PrincipalRepository
}

func (s PrincipalService) GetPrincipalByName(name string) *entity.Principal {
	return s.PrincipalRepository.FindByName(name)
}

func (s PrincipalService) InsertPrincipal(principal entity.Principal) *entity.Principal {
	principal.Uuid, _ = uuid.NewV4()
	return s.PrincipalRepository.Save(principal)
}

func (s PrincipalService) PostPrincipalData(principal *entity.Principal) (insertedPrincipal *entity.Principal, errRes *handler.ErrorResponse) {
	principalData := s.GetPrincipalByName(principal.Name)
	if principalData == nil {
		return nil, handler.InternalServerError("")
	}

	if len(principalData.Name) > 0 {
		return nil, handler.Conflict("")
	}

	principalData = s.InsertPrincipal(*principal)
	if principalData == nil {
		return nil, handler.InternalServerError("")
	}

	return principalData, nil
}