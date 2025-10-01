package pest_disease_record

import (
	"context"
	"production_service/domain/repository"
)

type DeletePestDiseaseRecordUsecase interface {
	Execute(ctx context.Context, id string) error
}

type deletePestDiseaseRecordUsecase struct {
	pestDiseaseRecordRepo repository.PestDiseaseRecordRepository
}

func NewDeletePestDiseaseRecordUsecase(pestDiseaseRecordRepo repository.PestDiseaseRecordRepository) DeletePestDiseaseRecordUsecase {
	return &deletePestDiseaseRecordUsecase{
		pestDiseaseRecordRepo: pestDiseaseRecordRepo,
	}
}

func (u *deletePestDiseaseRecordUsecase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return ErrPestDiseaseRecordNotFound
	}

	_, err := u.pestDiseaseRecordRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = u.pestDiseaseRecordRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
