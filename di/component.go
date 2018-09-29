package di

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/tomoyane/grant-n-z/domain/service"
)

var (
	ProvideUserService      service.UserService
	ProvideTokenService     service.TokenService
	ProvideRoleService      service.RoleService
	ProvidePrincipalService service.PrincipalService
	ProvideServiceService   service.ServiceService
	ProvideMemberService    service.MemberService
)

func InitUserService(repo repository.UserRepository) {
	ProvideUserService = service.UserService{
		UserRepository: repo,
	}
}

func InitTokenService(tRepo repository.TokenRepository, uRepo repository.UserRepository) {
	ProvideTokenService = service.TokenService{
		TokenRepository: tRepo,
		UserRepository:  uRepo,
	}
}

func InitRoleService(repo repository.RoleRepository) {
	ProvideRoleService = service.RoleService{
		RoleRepository: repo,
	}
}

func InitPrincipalService(pRepo repository.PrincipalRepository, uRepo repository.UserRepository,
	sRepo repository.ServiceRepository, mRepo repository.MemberRepository, rRepo repository.RoleRepository) {
	ProvidePrincipalService = service.PrincipalService{
		PrincipalRepository: pRepo,
		UserRepository:      uRepo,
		ServiceRepository:   sRepo,
		MemberRepository:    mRepo,
		RoleRepository:      rRepo,
	}
}

func InitServiceService(repo repository.ServiceRepository) {
	ProvideServiceService = service.ServiceService{
		ServiceRepository: repo,
	}
}

func InitMemberService(repo repository.MemberRepository) {
	ProvideMemberService = service.MemberService{
		MemberRepository: repo,
	}
}
