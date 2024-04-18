package worker

import (
	"context"

	"github.com/jeremykane/go-boilerplate/internal/config"
	"github.com/jeremykane/go-boilerplate/internal/service"
)

type (
	ClinicWorker struct {
		*Worker
		facilityTypeSvc service.FacilityTypeService
		schedulerConfig config.SchedulerConfig
	}

	ClinicWorkerParam struct {
		*Worker
		FacilityTypeSvc service.FacilityTypeService
		SchedulerConfig config.SchedulerConfig
	}
)

func NewClinicWorker(param ClinicWorkerParam) *ClinicWorker {
	cw := &ClinicWorker{
		Worker:          param.Worker,
		facilityTypeSvc: param.FacilityTypeSvc,
		schedulerConfig: param.SchedulerConfig,
	}
	return cw
}

func (cw *ClinicWorker) GenerateWorkerParameters() []JobParameter {
	return []JobParameter{
		{
			Name:     "Placeholder",
			TimeSpec: cw.config.PlaceholderSchedulerTime,
			Handler:  cw.PlaceholderHandler,
		},
	}
}

func (cw *ClinicWorker) PlaceholderHandler(ctx context.Context) {
	const funcName = "[worker][clinic]PlaceholderHandler"
	println("Placeholder Running")
	return
}
