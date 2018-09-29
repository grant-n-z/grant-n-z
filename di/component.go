package di

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/tomoyane/grant-n-z/domain/service"
)

var (
	ProviderUserService service.UserService
	ProviderTokenService service.TokenService
	ProviderRoleService service.RoleService
	ProviderPrincipalService service.PrincipalService
	ProviderServiceService service.ServiceService
)

func InitUserService(repo repository.UserRepository) {
	ProviderUserService = service.UserService{
		UserRepository: repo,
	}
}

func InitTokenService(tRepo repository.TokenRepository, uRepo repository.UserRepository) {
	ProviderTokenService = service.TokenService{
		TokenRepository: tRepo,
		UserRepository: uRepo,
	}
}

func InitRoleService(repo repository.RoleRepository) {
	ProviderRoleService = service.RoleService{
		RoleRepository: repo,
	}
}

func InitPrincipalService(pRepo repository.PrincipalRepository, uRepo repository.UserRepository,
	sRepo repository.ServiceRepository, mRepo repository.MemberRepository, rRepo repository.RoleRepository) {
	ProviderPrincipalService = service.PrincipalService{
		PrincipalRepository: pRepo,
		UserRepository: uRepo,
		ServiceRepository: sRepo,
		MemberRepository: mRepo,
		RoleRepository: rRepo,
	}
}

func InitServiceService(repo repository.ServiceRepository) {
	ProviderServiceService = service.ServiceService{
		ServiceRepository: repo,
	}
}
