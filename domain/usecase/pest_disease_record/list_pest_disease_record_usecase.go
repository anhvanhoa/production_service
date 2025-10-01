package pest_disease_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type ListPestDiseaseRecordUsecase interface {
	Execute(ctx context.Context, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error)
}

type listPestDiseaseRecordUsecase struct {
	pestDiseaseRecordRepo repository.PestDiseaseRecordRepository
}

func NewListPestDiseaseRecordUsecase(pestDiseaseRecordRepo repository.PestDiseaseRecordRepository) ListPestDiseaseRecordUsecase {
	return &listPestDiseaseRecordUsecase{
		pestDiseaseRecordRepo: pestDiseaseRecordRepo,
	}
}

func (u *listPestDiseaseRecordUsecase) Execute(ctx context.Context, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error) {
	pestDiseaseRecords, total, err := u.pestDiseaseRecordRepo.List(ctx, pagination, filter)
	if err != nil {
		return nil, 0, err
	}

	return pestDiseaseRecords, total, nil
}
