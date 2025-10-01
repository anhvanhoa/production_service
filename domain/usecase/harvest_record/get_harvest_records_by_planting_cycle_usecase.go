package harvest_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type GetHarvestRecordsByPlantingCycleUsecase interface {
	Execute(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error)
}

type getHarvestRecordsByPlantingCycleUsecase struct {
	harvestRecordRepo repository.HarvestRecordRepository
}

func NewGetHarvestRecordsByPlantingCycleUsecase(harvestRecordRepo repository.HarvestRecordRepository) GetHarvestRecordsByPlantingCycleUsecase {
	return &getHarvestRecordsByPlantingCycleUsecase{
		harvestRecordRepo: harvestRecordRepo,
	}
}

func (u *getHarvestRecordsByPlantingCycleUsecase) Execute(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error) {
	if plantingCycleID == "" {
		return nil, 0, ErrHarvestRecordNotFound
	}

	harvestRecords, total, err := u.harvestRecordRepo.GetByPlantingCycleID(ctx, plantingCycleID, pagination, filter)
	if err != nil {
		return nil, 0, err
	}

	return harvestRecords, total, nil
}
