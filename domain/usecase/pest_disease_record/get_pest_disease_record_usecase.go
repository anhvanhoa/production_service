package pest_disease_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"
)

type GetPestDiseaseRecordUsecase interface {
	Execute(ctx context.Context, id string) (*entity.PestDiseaseRecord, error)
}

type getPestDiseaseRecordUsecase struct {
	pestDiseaseRecordRepo repository.PestDiseaseRecordRepository
}

func NewGetPestDiseaseRecordUsecase(pestDiseaseRecordRepo repository.PestDiseaseRecordRepository) GetPestDiseaseRecordUsecase {
	return &getPestDiseaseRecordUsecase{
		pestDiseaseRecordRepo: pestDiseaseRecordRepo,
	}
}

func (u *getPestDiseaseRecordUsecase) Execute(ctx context.Context, id string) (*entity.PestDiseaseRecord, error) {
	if id == "" {
		return nil, ErrPestDiseaseRecordNotFound
	}

	pestDiseaseRecord, err := u.pestDiseaseRecordRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if pestDiseaseRecord == nil {
		return nil, ErrPestDiseaseRecordNotFound
	}

	return pestDiseaseRecord, nil
}
