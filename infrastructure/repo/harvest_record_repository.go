package repo

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type harvestRecordRepository struct {
	db    *pg.DB
	utils utils.Helper
}

func NewHarvestRecordRepository(db *pg.DB, utils utils.Helper) repository.HarvestRecordRepository {
	return &harvestRecordRepository{
		db:    db,
		utils: utils,
	}
}

func (r *harvestRecordRepository) Create(ctx context.Context, harvestRecord *entity.HarvestRecord) error {
	_, err := r.db.Model(harvestRecord).Context(ctx).Insert()
	return err
}

func (r *harvestRecordRepository) GetByID(ctx context.Context, id string) (*entity.HarvestRecord, error) {
	harvestRecord := &entity.HarvestRecord{}
	err := r.db.Model(harvestRecord).
		Where("id = ?", id).
		Context(ctx).
		Select()
	if err != nil {
		return nil, err
	}
	return harvestRecord, nil
}

func (r *harvestRecordRepository) GetByPlantingCycleID(ctx context.Context, plantingCycleID string, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error) {
	var harvestRecords []*entity.HarvestRecord

	query := r.db.Model(&harvestRecords).
		Where("planting_cycle_id = ?", plantingCycleID).
		Context(ctx)

	query = r.ApplyFilters(query, filter)

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	if pagination.PageSize > 0 {
		query = query.Limit(pagination.PageSize)
	}
	if pagination.Page > 0 {
		offset := r.utils.CalculateOffset(pagination.Page, pagination.PageSize)
		query = query.Offset(offset)
	}

	query = query.Order("created_at DESC")

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return harvestRecords, int64(total), nil
}

func (r *harvestRecordRepository) List(ctx context.Context, pagination common.Pagination, filter entity.FilterHarvestRecord) ([]*entity.HarvestRecord, int64, error) {
	var harvestRecords []*entity.HarvestRecord

	query := r.db.Model(&harvestRecords).Context(ctx)

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
	return harvestRecords, int64(total), nil
}

func (r *harvestRecordRepository) Update(ctx context.Context, harvestRecord *entity.HarvestRecord) error {
	_, err := r.db.Model(harvestRecord).
		Where("id = ?", harvestRecord.ID).
		Context(ctx).
		Update()
	return err
}

func (r *harvestRecordRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model(&entity.HarvestRecord{}).
		Where("id = ?", id).
		Context(ctx).
		Delete()
	return err
}

func (r *harvestRecordRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.db.Model(&entity.HarvestRecord{}).Context(ctx).Count()
	return int64(count), err
}

func (r *harvestRecordRepository) CountByPlantingCycleID(ctx context.Context, plantingCycleID string) (int64, error) {
	count, err := r.db.Model(&entity.HarvestRecord{}).
		Where("planting_cycle_id = ?", plantingCycleID).
		Context(ctx).
		Count()
	return int64(count), err
}

func (r *harvestRecordRepository) ApplyFilters(query *pg.Query, filter entity.FilterHarvestRecord) *pg.Query {
	if filter.HarvestDate != nil {
		query = query.Where("harvest_date = ?", filter.HarvestDate)
	}
	if filter.QualityGrade != "" {
		query = query.Where("quality_grade = ?", filter.QualityGrade)
	}
	if filter.SizeClassification != "" {
		query = query.Where("size_classification = ?", filter.SizeClassification)
	}
	if filter.MarketPricePerKg > 0 {
		query = query.Where("market_price_per_kg = ?", filter.MarketPricePerKg)
	}
	if filter.TotalRevenue > 0 {
		query = query.Where("total_revenue = ?", filter.TotalRevenue)
	}
	if filter.PlantHealthRating > 0 {
		query = query.Where("plant_health_rating = ?", filter.PlantHealthRating)
	}
	if filter.Notes != "" {
		query = query.Where("notes ILIKE ?", "%"+filter.Notes+"%")
	}
	if filter.Images != "" {
		query = query.Where("images ILIKE ?", "%"+filter.Images+"%")
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if !filter.CreatedAt.IsZero() {
		query = query.Where("created_at >= ?", filter.CreatedAt)
	}
	return query
}
