package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jeremykane/go-boilerplate/internal/constant/errorCode"
	"github.com/jeremykane/go-boilerplate/internal/entity"
	"github.com/jeremykane/go-boilerplate/internal/repository"
	"github.com/jeremykane/go-boilerplate/pkg/errorx"
	"github.com/jeremykane/go-boilerplate/pkg/logger"
	"github.com/jeremykane/go-boilerplate/pkg/logger/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	paramLogs := logrus.LogrusParam{
		Level: "DEBUG",
	}
	logs := logrus.NewLogrus(paramLogs)
	logger.NewLogger(logs)
}

func TestFacilityTypeService_GetAll(t *testing.T) {
	tcs := []struct {
		name                 string
		mockFacilityTypeRepo func(ctrl *gomock.Controller) repository.FacilityTypeRepository
		expectedResult       []entity.FacilityType
		expectedError        *errorx.CustomError
	}{
		{
			name: "Error case getAll error should return errorx object",
			mockFacilityTypeRepo: func(ctrl *gomock.Controller) repository.FacilityTypeRepository {
				mock := repository.NewMockFacilityTypeRepository(ctrl)
				mock.EXPECT().GetAll(gomock.Any()).Return(nil, errorx.NewError(errors.New("test error"), errorCode.ErrorQuery))
				return mock
			},
			expectedResult: nil,
			expectedError:  errorx.NewError(errors.New("test error"), errorCode.ErrorQuery),
		},
		{
			name: "success should return object of facility types",
			mockFacilityTypeRepo: func(ctrl *gomock.Controller) repository.FacilityTypeRepository {
				mock := repository.NewMockFacilityTypeRepository(ctrl)
				mock.EXPECT().GetAll(gomock.Any()).Return([]entity.FacilityType{}, nil)
				return mock
			},
			expectedResult: []entity.FacilityType{},
			expectedError:  nil,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			svc := &facilityTypeService{
				facilityTypeRepo: tc.mockFacilityTypeRepo(ctrl),
			}
			res, err := svc.GetAll(context.Background())
			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
