package repository

import (
	"context"

	"github.com/jeremykane/go-boilerplate/internal/config"
	"github.com/jeremykane/go-boilerplate/internal/constant/errorCode"
	"github.com/jeremykane/go-boilerplate/internal/entity"
	"github.com/jeremykane/go-boilerplate/pkg/database"
	"github.com/jeremykane/go-boilerplate/pkg/errorx"
	"github.com/jeremykane/go-boilerplate/pkg/logger"
)

type (
	FacilityTypeRepository interface {
		GetAll(ctx context.Context) ([]entity.FacilityType, *errorx.CustomError)
	}

	facilityTypeRepo struct {
		db *database.Replication
	}

	FacilityTypeRepoParam struct {
		DB map[string]*database.Replication
	}
)

func NewFacilityTypeRepository(param FacilityTypeRepoParam) FacilityTypeRepository {
	return &facilityTypeRepo{
		db: param.DB[config.DatabaseGo],
	}
}

func (repo *facilityTypeRepo) GetAll(ctx context.Context) ([]entity.FacilityType, *errorx.CustomError) {
	const funcName = "[repository][facility_type]GetAll"
	var facilityTypes []entity.FacilityType
	result := repo.db.Master.Order("id").Find(&facilityTypes)
	if result.Error != nil {
		logger.ErrorWithFields(ctx, funcName, logger.Fields{
			"err": result.Error.Error(),
		})
		return nil, errorx.NewError(result.Error, errorCode.ErrorQuery)
	}

	return facilityTypes, nil
}
