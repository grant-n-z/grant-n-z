package timer

import (
	"fmt"

	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzcacher/service"
)

const limit = 100

type Runner interface {
	// Run main process
	Run()
}

type RunnerImpl struct {
	UpdaterService   service.UpdaterService
	ExtractorService service.ExtractorService
}

func NewRunner() Runner {
	return RunnerImpl{
		UpdaterService:   service.NewUpdaterService(),
		ExtractorService: service.NewExtractorService(),
	}
}

func (r RunnerImpl) Run() {
	go r.executePolicy()
	go r.executePermission()
	go r.executeRole()
	go r.executeService()
	go r.executeUserService()
}

func (r RunnerImpl) executePolicy() {
	dataLength := 1
	offset := 0
	for dataLength != 0 {
		policies := r.ExtractorService.GetPolicies(offset, limit)
		r.UpdaterService.UpdatePolicy(policies)
		dataLength = len(policies)
		offset += limit
		log.Logger.Info(fmt.Sprintf("Update policy length = %d", dataLength))
	}
}

func (r RunnerImpl) executePermission() {
	dataLength := 1
	offset := 0
	for dataLength != 0 {
		permissions := r.ExtractorService.GetPermissions(offset, limit)
		r.UpdaterService.UpdatePermission(permissions)
		dataLength = len(permissions)
		offset += limit
		log.Logger.Info(fmt.Sprintf("Update permission length = %d", dataLength))
	}
}

func (r RunnerImpl) executeRole() {
	dataLength := 1
	offset := 0
	for dataLength != 0 {
		roles := r.ExtractorService.GetRoles(offset, limit)
		r.UpdaterService.UpdateRole(roles)
		dataLength = len(roles)
		offset += limit
		log.Logger.Info(fmt.Sprintf("Update role length = %d", dataLength))
	}
}

func (r RunnerImpl) executeService() {
	dataLength := 1
	offset := 0
	for dataLength != 0 {
		services := r.ExtractorService.GetServices(offset, limit)
		r.UpdaterService.UpdateService(services)
		dataLength = len(services)
		offset += limit
		log.Logger.Info(fmt.Sprintf("Update service length = %d", dataLength))
	}
}

func (r RunnerImpl) executeUserService() {
	dataLength := 1
	offset := 0
	for dataLength != 0 {
		userServices := r.ExtractorService.GetUserServices(offset, limit)
		r.UpdaterService.UpdateUserService(userServices)
		dataLength = len(userServices)
		offset += limit
		log.Logger.Info(fmt.Sprintf("Update user_service length = %d", dataLength))
	}
}
