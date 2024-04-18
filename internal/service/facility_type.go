package service

import (
	"context"

	"github.com/jeremykane/go-boilerplate/internal/entity"
	"github.com/jeremykane/go-boilerplate/internal/repository"
	"github.com/jeremykane/go-boilerplate/pkg/errorx"
	"github.com/jeremykane/go-boilerplate/pkg/logger"
)

type (
	FacilityTypeService interface {
		GetAll(ctx context.Context) ([]entity.FacilityType, *errorx.CustomError)
	}

	FacilityTypeServiceParam struct {
		FacilityTypeRepo repository.FacilityTypeRepository
	}

	facilityTypeService struct {
		facilityTypeRepo repository.FacilityTypeRepository
	}
)

func NewFacilityTypeService(param FacilityTypeServiceParam) FacilityTypeService {
	return &facilityTypeService{
		facilityTypeRepo: param.FacilityTypeRepo,
	}
}

func (s *facilityTypeService) GetAll(ctx context.Context) ([]entity.FacilityType, *errorx.CustomError) {
	const funcName = "[service][facility_type]GetAll"
	facilityType, err := s.facilityTypeRepo.GetAll(ctx)
	if err != nil {
		logger.ErrorWithFields(ctx, funcName, logger.Fields{
			"err": err,
		})
		return nil, err
	}
	return facilityType, nil
}
