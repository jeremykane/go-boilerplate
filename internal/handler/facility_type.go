package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/jeremykane/go-boilerplate/internal/service"
	"github.com/jeremykane/go-boilerplate/pkg/logger"
)

type (
	FacilityTypeHandler struct {
		facilityTypeSvc service.FacilityTypeService
	}

	FacilityTypeHandlerParam struct {
		FacilityTypeSvc service.FacilityTypeService
	}
)

func NewFacilityTypeHandler(param FacilityTypeHandlerParam) *FacilityTypeHandler {
	return &FacilityTypeHandler{
		facilityTypeSvc: param.FacilityTypeSvc,
	}
}

// GetAll @Summary Get All facility Type
// @Description get facility type order by id
// @Tags [GET] facility type
// @Accept */*
// @Produce json
// @Success 200 {object} []entity.FacilityType
// @Failure 500 {object} entity.GeneralAPIResponse
// @Router /facilityTypes [get]
func (h *FacilityTypeHandler) GetAll(ctx echo.Context) error {
	const funcName = "[handler][facility_type]GetAll"

	context := ctx.Request().Context()
	result, errx := h.facilityTypeSvc.GetAll(context)
	if errx != nil {
		logger.ErrorWithFields(context, funcName, logger.Fields{
			"err": errx.Error(),
		})
		return BadResponseError(ctx, errx)
	}

	return SuccessResponse(ctx, result)
}
