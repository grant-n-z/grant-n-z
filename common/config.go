package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tomoyane/grant-n-z/domain/service"
	"github.com/tomoyane/grant-n-z/domain/repository"
)

var (
	ProvideUserService service.UserService
)

func InitUserService(repo repository.UserRepository) {
	ProvideUserService = service.UserService{
		UserRepository: repo,
	}
}
