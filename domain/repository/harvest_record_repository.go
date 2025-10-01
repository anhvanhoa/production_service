package repository

import (
	"context"
	"production_service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type HarvestRecordRepository interface {
	Create(ctx context.Context, harvestRecord *entity.HarvestRecord) error
	GetByID(ctx context.Context, id string) (*entity.HarvestRecord, error)
	GetByPlantingCycleID(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error)
	List(ctx context.Context, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error)
	Update(ctx context.Context, harvestRecord *entity.HarvestRecord) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	CountByPlantingCycleID(ctx context.Context, plantingCycleID string) (int64, error)
}
