package harvest_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"
)

type GetHarvestRecordUsecase interface {
	Execute(ctx context.Context, id string) (*entity.HarvestRecord, error)
}

type getHarvestRecordUsecase struct {
	harvestRecordRepo repository.HarvestRecordRepository
}

func NewGetHarvestRecordUsecase(harvestRecordRepo repository.HarvestRecordRepository) GetHarvestRecordUsecase {
	return &getHarvestRecordUsecase{
		harvestRecordRepo: harvestRecordRepo,
	}
}

func (u *getHarvestRecordUsecase) Execute(ctx context.Context, id string) (*entity.HarvestRecord, error) {
	if id == "" {
		return nil, ErrHarvestRecordNotFound
	}

	harvestRecord, err := u.harvestRecordRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if harvestRecord == nil {
		return nil, ErrHarvestRecordNotFound
	}

	return harvestRecord, nil
}
