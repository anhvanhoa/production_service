package harvest_record

import (
	"context"
	"production_service/domain/repository"
)

type DeleteHarvestRecordUsecase interface {
	Execute(ctx context.Context, id string) error
}

type deleteHarvestRecordUsecase struct {
	harvestRecordRepo repository.HarvestRecordRepository
}

func NewDeleteHarvestRecordUsecase(harvestRecordRepo repository.HarvestRecordRepository) DeleteHarvestRecordUsecase {
	return &deleteHarvestRecordUsecase{
		harvestRecordRepo: harvestRecordRepo,
	}
}

func (u *deleteHarvestRecordUsecase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return ErrHarvestRecordNotFound
	}

	_, err := u.harvestRecordRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = u.harvestRecordRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
