package timer

import "github.com/tomoyane/grant-n-z/gnzcacher/service"

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
		UpdaterService: service.NewUpdaterService(),
		ExtractorService: service.NewExtractorService(),
	}
}

func (RunnerImpl) Run() {

}
