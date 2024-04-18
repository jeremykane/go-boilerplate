package common

import (
	"errors"

	"github.com/jeremykane/go-boilerplate/internal/constant/errorCode"
	"github.com/jeremykane/go-boilerplate/pkg/errorx"
)

var (
	ErrInternalServerError = errorx.NewError(errors.New("Internal Server Error"), errorCode.ErrInternalServerError)
	ErrBadRequest          = errorx.NewError(errors.New("Bad request"), errorCode.ErrBadRequest)
)
