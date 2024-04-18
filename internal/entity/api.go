package entity

import "github.com/jeremykane/go-boilerplate/pkg/errorx"

type (
	GeneralAPIResponse struct {
		Data  interface{}         `json:"data,omitempty"`
		Error *errorx.CustomError `json:"error,omitempty"`
	}
)
