package harvest_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type ListHarvestRecordUsecase interface {
	Execute(ctx context.Context, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error)
}

type listHarvestRecordUsecase struct {
	harvestRecordRepo repository.HarvestRecordRepository
}

func NewListHarvestRecordUsecase(harvestRecordRepo repository.HarvestRecordRepository) ListHarvestRecordUsecase {
	return &listHarvestRecordUsecase{
		harvestRecordRepo: harvestRecordRepo,
	}
}

func (u *listHarvestRecordUsecase) Execute(ctx context.Context, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error) {
	harvestRecords, total, err := u.harvestRecordRepo.List(ctx, pagination, filter)
	if err != nil {
		return nil, 0, err
	}
	return harvestRecords, total, nil
}
