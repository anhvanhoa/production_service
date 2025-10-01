package harvest_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"
	"time"
)

type CreateHarvestRecordUsecase interface {
	Execute(ctx context.Context, req *CreateHarvestRecordRequest) (*entity.HarvestRecord, error)
}

type CreateHarvestRecordRequest struct {
	PlantingCycleID    string
	HarvestDate        *time.Time
	HarvestTime        *time.Time
	QuantityKg         float64
	QualityGrade       string
	SizeClassification string
	MarketPricePerKg   float64
	LaborHours         float64
	LaborCost          float64
	PackagingCost      float64
	StorageLocation    string
	StorageTemperature float64
	BuyerInformation   string
	DeliveryDate       *time.Time
	WeatherAtHarvest   string
	PlantHealthRating  int
	Notes              string
	Images             string
	CreatedBy          string
}

type createHarvestRecordUsecase struct {
	harvestRecordRepo repository.HarvestRecordRepository
}

func NewCreateHarvestRecordUsecase(harvestRecordRepo repository.HarvestRecordRepository) CreateHarvestRecordUsecase {
	return &createHarvestRecordUsecase{
		harvestRecordRepo: harvestRecordRepo,
	}
}

func (u *createHarvestRecordUsecase) Execute(ctx context.Context, req *CreateHarvestRecordRequest) (*entity.HarvestRecord, error) {
	qualityGrade := entity.QualityGrade(req.QualityGrade)
	if !qualityGrade.IsValid() {
		return nil, ErrInvalidQualityGrade
	}

	sizeClassification := entity.SizeClassification(req.SizeClassification)
	if !sizeClassification.IsValid() {
		return nil, ErrInvalidSizeClassification
	}
	if !entity.IsValidPlantHealthRating(req.PlantHealthRating) {
		return nil, ErrInvalidPlantHealthRating
	}

	totalRevenue := req.QuantityKg * req.MarketPricePerKg
	harvestRecord := &entity.HarvestRecord{
		PlantingCycleID:    req.PlantingCycleID,
		HarvestDate:        req.HarvestDate,
		HarvestTime:        req.HarvestTime,
		QuantityKg:         req.QuantityKg,
		QualityGrade:       req.QualityGrade,
		SizeClassification: req.SizeClassification,
		MarketPricePerKg:   req.MarketPricePerKg,
		TotalRevenue:       totalRevenue,
		LaborHours:         req.LaborHours,
		LaborCost:          req.LaborCost,
		PackagingCost:      req.PackagingCost,
		StorageLocation:    req.StorageLocation,
		StorageTemperature: req.StorageTemperature,
		BuyerInformation:   req.BuyerInformation,
		DeliveryDate:       req.DeliveryDate,
		WeatherAtHarvest:   req.WeatherAtHarvest,
		PlantHealthRating:  req.PlantHealthRating,
		Notes:              req.Notes,
		Images:             req.Images,
		CreatedBy:          req.CreatedBy,
		CreatedAt:          time.Now(),
	}

	// Save to repository
	err := u.harvestRecordRepo.Create(ctx, harvestRecord)
	if err != nil {
		return nil, err
	}

	return harvestRecord, nil
}
