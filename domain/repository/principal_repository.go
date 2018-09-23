package repository

import "github.com/tomoyane/grant-n-z/domain/entity"

type PrincipalRepository interface {
	FindByName(name string) *entity.Principal

	Save(principal entity.Principal) *entity.Principal
}
