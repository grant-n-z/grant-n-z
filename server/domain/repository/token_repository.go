package repository

import "github.com/tomoyane/grant-n-z/server/domain/entity"

type TokenRepository interface {
	FindByUserUuid(userUuid string) *entity.Token

	Save(token entity.Token) *entity.Token
}
