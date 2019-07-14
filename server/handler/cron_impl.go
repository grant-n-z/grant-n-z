package handler

import (
	"fmt"
	"os"

	"github.com/jasonlvhit/gocron"
	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/log"
)

const minutes = 5

type CronHandlerImpl struct {
	cron gocron.Scheduler
	app  config.AppConfig
}

func NewCronHandler() CronHandler {
	return CronHandlerImpl{
		cron: *gocron.NewScheduler(),
		app: config.App,
	}
}

func (ch CronHandlerImpl) RunUpdatePolicy() {
	ch.cron.Every(minutes).Minutes().Do(ch.Run)
	ch.cron.Start()
}

func (ch CronHandlerImpl) Run() {
	log.Logger.Info("Run update policy file")

	path := fmt.Sprintf("%spolicy.json", ch.app.PolicyFilePath)
	file, err := os.Open(path)
	if err != nil {
		file, err = os.Create(path)
		if err != nil {
			log.Logger.Error("Error write policy file", err.Error())
		}
	}
	defer file.Close()

	output := "{'key': 'value'}"
	_, _ = file.Write(([]byte)(output))
}
