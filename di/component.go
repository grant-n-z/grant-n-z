package di

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tomoyane/grant-n-z/domain/service"
	"github.com/tomoyane/grant-n-z/domain/repository"
)

var (
	ProviderUserService service.UserService
	ProviderTokenService service.TokenService
)

func InitUserService(repo repository.UserRepository) {
	ProviderUserService = service.UserService{
		UserRepository: repo,
	}
}

func InitTokenService(repo repository.TokenRepository) {
	ProviderTokenService = service.TokenService{
		TokenRepository: repo,
	}
}

