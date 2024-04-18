package api

import (
	"github.com/jeremykane/go-boilerplate/internal/handler"
	"github.com/jeremykane/go-boilerplate/internal/service"
)

type (
	// Handler is the unified API module which stores all available API handlers.
	Handler struct {
		*handler.FacilityTypeHandler
	}

	HandlerParams struct {
		FacilityTypeSvc service.FacilityTypeService
	}
)

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		FacilityTypeHandler: handler.NewFacilityTypeHandler(func() handler.FacilityTypeHandlerParam {
			return handler.FacilityTypeHandlerParam{
				FacilityTypeSvc: params.FacilityTypeSvc,
			}
		}()),
	}
}
