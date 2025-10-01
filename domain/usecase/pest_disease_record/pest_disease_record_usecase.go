package pest_disease_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type PestDiseaseRecordUsecase interface {
	Create(ctx context.Context, req *CreatePestDiseaseRecordRequest) (*entity.PestDiseaseRecord, error)
	GetByID(ctx context.Context, id string) (*entity.PestDiseaseRecord, error)
	Update(ctx context.Context, req *UpdatePestDiseaseRecordRequest) (*entity.PestDiseaseRecord, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error)
	GetByPlantingCycleID(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error)
}

type pestDiseaseRecordUsecase struct {
	createUsecase             CreatePestDiseaseRecordUsecase
	getUsecase                GetPestDiseaseRecordUsecase
	updateUsecase             UpdatePestDiseaseRecordUsecase
	deleteUsecase             DeletePestDiseaseRecordUsecase
	listUsecase               ListPestDiseaseRecordUsecase
	getByPlantingCycleUsecase GetPestDiseaseRecordsByPlantingCycleUsecase
}

func NewPestDiseaseRecordUsecase(pestDiseaseRecordRepo repository.PestDiseaseRecordRepository) PestDiseaseRecordUsecase {
	return &pestDiseaseRecordUsecase{
		createUsecase:             NewCreatePestDiseaseRecordUsecase(pestDiseaseRecordRepo),
		getUsecase:                NewGetPestDiseaseRecordUsecase(pestDiseaseRecordRepo),
		updateUsecase:             NewUpdatePestDiseaseRecordUsecase(pestDiseaseRecordRepo),
		deleteUsecase:             NewDeletePestDiseaseRecordUsecase(pestDiseaseRecordRepo),
		listUsecase:               NewListPestDiseaseRecordUsecase(pestDiseaseRecordRepo),
		getByPlantingCycleUsecase: NewGetPestDiseaseRecordsByPlantingCycleUsecase(pestDiseaseRecordRepo),
	}
}

func (u *pestDiseaseRecordUsecase) Create(ctx context.Context, req *CreatePestDiseaseRecordRequest) (*entity.PestDiseaseRecord, error) {
	return u.createUsecase.Execute(ctx, req)
}

func (u *pestDiseaseRecordUsecase) GetByID(ctx context.Context, id string) (*entity.PestDiseaseRecord, error) {
	return u.getUsecase.Execute(ctx, id)
}

func (u *pestDiseaseRecordUsecase) Update(ctx context.Context, req *UpdatePestDiseaseRecordRequest) (*entity.PestDiseaseRecord, error) {
	return u.updateUsecase.Execute(ctx, req)
}

func (u *pestDiseaseRecordUsecase) Delete(ctx context.Context, id string) error {
	return u.deleteUsecase.Execute(ctx, id)
}

func (u *pestDiseaseRecordUsecase) List(ctx context.Context, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error) {
	return u.listUsecase.Execute(ctx, pagination, filter)
}

func (u *pestDiseaseRecordUsecase) GetByPlantingCycleID(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error) {
	return u.getByPlantingCycleUsecase.Execute(ctx, plantingCycleID, pagination, filter)
}
