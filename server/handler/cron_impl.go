package handler

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

const minutes = 5

type CronHandlerImpl struct {
	cron          gocron.Scheduler
	app           config.AppConfig
	policyService service.PolicyService
}

func NewCronHandlerImpl() CronHandler {
	return CronHandlerImpl{
		cron:          *gocron.NewScheduler(),
		app:           config.App,
		policyService: service.NewPolicyService(),
	}
}

func (ch CronHandlerImpl) RunUpdatePolicy() {
	ch.cron.Every(minutes).Minutes().Do(ch.Run)
	ch.cron.Start()
}

func (ch CronHandlerImpl) Run() {
	log.Logger.Info("Run update policy file")
	a, _ := ch.policyService.EncryptData("Test")
	fmt.Println(*a)

	b, _ := ch.policyService.DecryptData(*a)
	fmt.Println(*b)
}
