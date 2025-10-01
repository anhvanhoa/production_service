package pest_disease_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type GetPestDiseaseRecordsByPlantingCycleUsecase interface {
	Execute(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error)
}

type getPestDiseaseRecordsByPlantingCycleUsecase struct {
	pestDiseaseRecordRepo repository.PestDiseaseRecordRepository
}

func NewGetPestDiseaseRecordsByPlantingCycleUsecase(pestDiseaseRecordRepo repository.PestDiseaseRecordRepository) GetPestDiseaseRecordsByPlantingCycleUsecase {
	return &getPestDiseaseRecordsByPlantingCycleUsecase{
		pestDiseaseRecordRepo: pestDiseaseRecordRepo,
	}
}

func (u *getPestDiseaseRecordsByPlantingCycleUsecase) Execute(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error) {
	if plantingCycleID == "" {
		return nil, 0, ErrPestDiseaseRecordNotFound
	}

	pestDiseaseRecords, total, err := u.pestDiseaseRecordRepo.GetByPlantingCycleID(ctx, plantingCycleID, pagination, filter)
	if err != nil {
		return nil, 0, err
	}

	return pestDiseaseRecords, total, nil
}
