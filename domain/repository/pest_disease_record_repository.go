package repository

import (
	"context"
	"production_service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type PestDiseaseRecordRepository interface {
	Create(ctx context.Context, pestDiseaseRecord *entity.PestDiseaseRecord) error
	GetByID(ctx context.Context, id string) (*entity.PestDiseaseRecord, error)
	List(ctx context.Context, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error)
	GetByPlantingCycleID(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error)
	Update(ctx context.Context, pestDiseaseRecord *entity.PestDiseaseRecord) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	CountByPlantingCycleID(ctx context.Context, plantingCycleID string) (int64, error)
}
