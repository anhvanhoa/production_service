package harvest_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type HarvestRecordUsecase interface {
	Create(ctx context.Context, req *CreateHarvestRecordRequest) (*entity.HarvestRecord, error)
	GetByID(ctx context.Context, id string) (*entity.HarvestRecord, error)
	Update(ctx context.Context, req *UpdateHarvestRecordRequest) (*entity.HarvestRecord, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error)
	GetByPlantingCycleID(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error)
}

type harvestRecordUsecase struct {
	createUsecase             CreateHarvestRecordUsecase
	getUsecase                GetHarvestRecordUsecase
	updateUsecase             UpdateHarvestRecordUsecase
	deleteUsecase             DeleteHarvestRecordUsecase
	listUsecase               ListHarvestRecordUsecase
	getByPlantingCycleUsecase GetHarvestRecordsByPlantingCycleUsecase
}

func NewHarvestRecordUsecase(harvestRecordRepo repository.HarvestRecordRepository) HarvestRecordUsecase {
	return &harvestRecordUsecase{
		createUsecase:             NewCreateHarvestRecordUsecase(harvestRecordRepo),
		getUsecase:                NewGetHarvestRecordUsecase(harvestRecordRepo),
		updateUsecase:             NewUpdateHarvestRecordUsecase(harvestRecordRepo),
		deleteUsecase:             NewDeleteHarvestRecordUsecase(harvestRecordRepo),
		listUsecase:               NewListHarvestRecordUsecase(harvestRecordRepo),
		getByPlantingCycleUsecase: NewGetHarvestRecordsByPlantingCycleUsecase(harvestRecordRepo),
	}
}

func (u *harvestRecordUsecase) Create(ctx context.Context, req *CreateHarvestRecordRequest) (*entity.HarvestRecord, error) {
	return u.createUsecase.Execute(ctx, req)
}

func (u *harvestRecordUsecase) GetByID(ctx context.Context, id string) (*entity.HarvestRecord, error) {
	return u.getUsecase.Execute(ctx, id)
}

func (u *harvestRecordUsecase) Update(ctx context.Context, req *UpdateHarvestRecordRequest) (*entity.HarvestRecord, error) {
	return u.updateUsecase.Execute(ctx, req)
}

func (u *harvestRecordUsecase) Delete(ctx context.Context, id string) error {
	return u.deleteUsecase.Execute(ctx, id)
}

func (u *harvestRecordUsecase) List(ctx context.Context, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error) {
	return u.listUsecase.Execute(ctx, pagination, filter)
}

func (u *harvestRecordUsecase) GetByPlantingCycleID(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error) {
	return u.getByPlantingCycleUsecase.Execute(ctx, plantingCycleID, pagination, filter)
}
