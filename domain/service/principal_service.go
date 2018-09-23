package service

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/satori/go.uuid"
)

type PrincipalService struct {
	PrincipalRepository repository.PrincipalRepository
}

func (s PrincipalService) GetPrincipalByName(name string) *entity.Principal {
	return s.PrincipalRepository.FindByName(name)
}

func (s PrincipalService) InsertRole(principal entity.Principal) *entity.Principal {
	principal.Uuid, _ = uuid.NewV4()
	return s.PrincipalRepository.Save(principal)
}
