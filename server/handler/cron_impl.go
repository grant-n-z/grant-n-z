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
	cron               gocron.Scheduler
	app                config.AppConfig
	policyLocalService service.PolicyLocalService
}

func NewCronHandlerImpl() CronHandler {
	return CronHandlerImpl{
		cron:               *gocron.NewScheduler(),
		app:                config.App,
		policyLocalService: service.NewPolicyLocalService(),
	}
}

func (ch CronHandlerImpl) RunUpdatePolicy() {
	ch.cron.Every(minutes).Minutes().Do(ch.Run)
	ch.cron.Start()
}

func (ch CronHandlerImpl) Run() {
	log.Logger.Info("Run update policy file")
	a, _ := ch.policyLocalService.EncryptData("Test")
	fmt.Println(*a)

	b, _ := ch.policyLocalService.DecryptData(*a)
	fmt.Println(*b)
}
