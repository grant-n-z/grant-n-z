package repository

import "github.com/tomoyane/grant-n-z/domain/entity"

type TokenRepository interface {
	FindByUserId(userId string) *entity.Token

	Save(token entity.Token) *entity.Token
}
