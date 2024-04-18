package logrus

import (
	"context"

	"github.com/jeremykane/go-boilerplate/pkg/logger"
)

type (
// Fields is standard key - value pair type for log

)

var (
	ContextFields = []string{
		"trace_id",
		"span_id",
	}
)

// MergeFields will merge fld2 and fld1 into new fields.
// fld1 will be prioritized, If any same key on both fields than value on fld1 will be used.
func MergeFields(fld1 logger.Fields, fld2 logger.Fields) logger.Fields {
	res := make(logger.Fields, 0)
	for key, value := range fld1 {
		res[key] = value
	}

	for key, value := range fld2 {
		if _, ok := res[key]; !ok {
			res[key] = value
		}
	}

	return res
}

// MergeContextWithFields will merge specific context value as a new fields
func MergeContextWithFields(ctx context.Context, fields logger.Fields) logger.Fields {
	res := make(logger.Fields, 0)
	for key, value := range fields {
		res[key] = value
	}

	for _, v := range ContextFields {
		if value := ctx.Value(v); value != nil {
			res[v] = value
		}
	}
	return res
}
