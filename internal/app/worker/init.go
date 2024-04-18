package worker

import (
	"context"
	"fmt"

	"github.com/jeremykane/go-boilerplate/internal/config"
	"github.com/jeremykane/go-boilerplate/internal/service"
	"github.com/jeremykane/go-boilerplate/pkg/logger"
	"github.com/robfig/cron/v3"
)

type (
	Worker struct {
		cronPool *cron.Cron
		config   *config.Config
		cronJobs []worker
	}

	WorkerParam struct {
		Config          *config.Config
		FacilityTypeSvc service.FacilityTypeService
	}

	JobParameter struct {
		Name     string
		TimeSpec string
		Handler  func(ctx context.Context)
	}

	worker interface {
		GenerateWorkerParameters() []JobParameter
	}
)

func NewWorker(param WorkerParam) *Worker {
	cronPool := cron.New(
		cron.WithParser(
			cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor),
		),
	)

	cw := Worker{
		cronPool: cronPool,
		config:   param.Config,
	}

	cw.cronJobs = []worker{
		NewClinicWorker(ClinicWorkerParam{
			Worker:          &cw,
			FacilityTypeSvc: param.FacilityTypeSvc,
			SchedulerConfig: param.Config.SchedulerConfig,
		}),
	}

	cw.registerJobs()

	return &cw
}

func (cw *Worker) registerJobs() {
	fmt.Println("Registering jobs...")
	defer fmt.Println("Jobs are successfully registered")

	for i := range cw.cronJobs {
		for _, cronJob := range cw.cronJobs[i].GenerateWorkerParameters() {
			if _, err := cw.cronPool.AddFunc(cronJob.TimeSpec, cw.GenerateHandler(cronJob.Handler)); err != nil {
				logger.FatalWithFields(context.Background(), "failed to add worker", logger.Fields{
					"jobName":  cronJob.Name,
					"timeSpec": cronJob.TimeSpec,
				})
			}
		}
	}
}

func (cw *Worker) GenerateHandler(f func(ctx context.Context)) func() {
	return func() {
		f(context.Background())
	}
}

// Start starts all the registered jobs
func (cw *Worker) Start() {
	fmt.Println("Starting worker...")
	cw.cronPool.Start()
}

// Stop stops all currently running jobs
func (cw *Worker) Stop() {
	fmt.Println("Stopping worker...")
	defer fmt.Println("Cron is gracefully stopped")

	cw.cronPool.Stop()
}
