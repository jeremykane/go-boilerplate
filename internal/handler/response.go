package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jeremykane/go-boilerplate/internal/entity"
	"github.com/jeremykane/go-boilerplate/pkg/errorx"
	"github.com/labstack/echo/v4"
)

// SetJSONResponse is a helper to set JSON response with a customizable JSON encoder.
func SetJSONResponse(ctx echo.Context, data interface{}) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	ctx.Response().WriteHeader(http.StatusOK)

	response := entity.GeneralAPIResponse{
		Data:  data,
		Error: nil,
	}

	return json.NewEncoder(ctx.Response()).Encode(response)
}
func SetJSONErrorResponse(ctx echo.Context, e *errorx.CustomError, httpStatusCode int32) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	ctx.Response().WriteHeader(int(httpStatusCode))

	response := entity.GeneralAPIResponse{
		Error: e,
		Data:  nil,
	}

	return json.NewEncoder(ctx.Response()).Encode(response)
}

func BadResponseError(c echo.Context, e *errorx.CustomError) error {
	return SetJSONErrorResponse(c, e, http.StatusBadRequest)
}

func SuccessResponse(c echo.Context, data interface{}) error {
	return SetJSONResponse(c, data)
}
