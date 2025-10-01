package repo

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type pestDiseaseRecordRepository struct {
	db    *pg.DB
	utils utils.Helper
}

func NewPestDiseaseRecordRepository(db *pg.DB, utils utils.Helper) repository.PestDiseaseRecordRepository {
	return &pestDiseaseRecordRepository{
		db:    db,
		utils: utils,
	}
}

func (r *pestDiseaseRecordRepository) Create(ctx context.Context, pestDiseaseRecord *entity.PestDiseaseRecord) error {
	_, err := r.db.Model(pestDiseaseRecord).Context(ctx).Insert()
	return err
}

func (r *pestDiseaseRecordRepository) GetByID(ctx context.Context, id string) (*entity.PestDiseaseRecord, error) {
	pestDiseaseRecord := &entity.PestDiseaseRecord{}
	err := r.db.Model(pestDiseaseRecord).
		Where("id = ?", id).
		Context(ctx).
		Select()

	if err != nil {
		return nil, err
	}

	return pestDiseaseRecord, nil
}

func (r *pestDiseaseRecordRepository) GetByPlantingCycleID(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error) {
	var pestDiseaseRecords []*entity.PestDiseaseRecord

	query := r.db.Model(&pestDiseaseRecords).
		Where("planting_cycle_id = ?", plantingCycleID).
		Context(ctx)

	query = r.ApplyFilters(query, filter)
	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	query = query.Limit(pagination.PageSize)
	if pagination.Page > 0 {
		offset := r.utils.CalculateOffset(pagination.Page, pagination.PageSize)
		query = query.Offset(offset)
	}

	query = query.Order(pagination.SortBy + " " + pagination.SortOrder)
	err = query.Select()
	if err != nil {
		return nil, 0, err
	}
	return pestDiseaseRecords, int64(total), nil
}

func (r *pestDiseaseRecordRepository) List(ctx context.Context, pagination common.Pagination, filter entity.FilterPestDiseaseRecord) ([]*entity.PestDiseaseRecord, int64, error) {
	var pestDiseaseRecords []*entity.PestDiseaseRecord
	query := r.db.Model(&pestDiseaseRecords).Context(ctx)
	query = r.ApplyFilters(query, filter)
	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	query = query.Limit(pagination.PageSize)
	if pagination.Page > 0 {
		offset := r.utils.CalculateOffset(pagination.Page, pagination.PageSize)
		query = query.Offset(offset)
	}

	query = query.Order(pagination.SortBy + " " + pagination.SortOrder)
	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return pestDiseaseRecords, int64(total), nil
}

func (r *pestDiseaseRecordRepository) Update(ctx context.Context, pestDiseaseRecord *entity.PestDiseaseRecord) error {
	_, err := r.db.Model(pestDiseaseRecord).
		Where("id = ?", pestDiseaseRecord.ID).
		Context(ctx).
		Update()
	return err
}

func (r *pestDiseaseRecordRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model(&entity.PestDiseaseRecord{}).
		Where("id = ?", id).
		Context(ctx).
		Delete()
	return err
}

func (r *pestDiseaseRecordRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.db.Model(&entity.PestDiseaseRecord{}).Context(ctx).Count()
	return int64(count), err
}

func (r *pestDiseaseRecordRepository) CountByPlantingCycleID(ctx context.Context, plantingCycleID string) (int64, error) {
	count, err := r.db.Model(&entity.PestDiseaseRecord{}).
		Where("planting_cycle_id = ?", plantingCycleID).
		Context(ctx).
		Count()
	return int64(count), err
}

func (r *pestDiseaseRecordRepository) ApplyFilters(query *pg.Query, filter entity.FilterPestDiseaseRecord) *pg.Query {
	if filter.DetectionDate != nil {
		query = query.Where("detection_date = ?", filter.DetectionDate)
	}
	if filter.DetectionMethod != "" {
		query = query.Where("detection_method = ?", filter.DetectionMethod)
	}
	if filter.TreatmentDate != nil {
		query = query.Where("treatment_date = ?", filter.TreatmentDate)
	}
	if filter.TreatmentDurationDays > 0 {
		query = query.Where("treatment_duration_days = ?", filter.TreatmentDurationDays)
	}
	if filter.FollowUpDate != nil {
		query = query.Where("follow_up_date = ?", filter.FollowUpDate)
	}
	return query
}
