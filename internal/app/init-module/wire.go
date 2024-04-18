//go:build wireinject
// +build wireinject

package init_module

import (
	"context"
	"net/http"

	"github.com/google/wire"

	"github.com/jeremykane/go-boilerplate/internal/app/api"
	"github.com/jeremykane/go-boilerplate/internal/app/worker"
	"github.com/jeremykane/go-boilerplate/internal/config"
	"github.com/jeremykane/go-boilerplate/internal/repository"
	"github.com/jeremykane/go-boilerplate/internal/service"
)

var (
	cfgSet = wire.NewSet(
		wire.FieldsOf(new(*config.Config), "Server"),
		wire.FieldsOf(new(*config.Config), "Database"),
		wire.FieldsOf(new(*config.Config), "SchedulerConfig"),
	)

	dependencySet = wire.NewSet(
		InitializeDB,
	)

	repositorySet = wire.NewSet(
		repository.NewFacilityTypeRepository,
		wire.Struct(new(repository.FacilityTypeRepoParam), "*"),
	)

	serviceSet = wire.NewSet(
		service.NewFacilityTypeService,
		wire.Struct(new(service.FacilityTypeServiceParam), "*"),
	)

	appSet = wire.NewSet(
		wire.Struct(new(api.HandlerParams), "*"),
		api.NewHandler,
		wire.Struct(new(worker.WorkerParam), "*"),
		worker.NewWorker,
		NewAPIWrapper,
		NewWorkerWrapper,
		NewAPIGracefulHandler,
		NewWorkerGracefulHandler,
	)

	// List of all instances
	allSet = wire.NewSet(
		cfgSet,
		dependencySet,
		repositorySet,
		serviceSet,
		appSet,
	)
)

func NewAPI(ctx context.Context, cfg *config.Config, httpClient *http.Client) *APIWrapper {
	wire.Build(allSet)
	return &APIWrapper{}
}

func NewWorker(ctx context.Context, cfg *config.Config, httpClient *http.Client) *WorkerWrapper {
	wire.Build(allSet)
	return &WorkerWrapper{}
}
