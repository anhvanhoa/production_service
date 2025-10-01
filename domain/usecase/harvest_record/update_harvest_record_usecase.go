package harvest_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"
	"time"
)

type UpdateHarvestRecordUsecase interface {
	Execute(ctx context.Context, req *UpdateHarvestRecordRequest) (*entity.HarvestRecord, error)
}

type UpdateHarvestRecordRequest struct {
	ID                 string
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
}

type updateHarvestRecordUsecase struct {
	harvestRecordRepo repository.HarvestRecordRepository
}

func NewUpdateHarvestRecordUsecase(harvestRecordRepo repository.HarvestRecordRepository) UpdateHarvestRecordUsecase {
	return &updateHarvestRecordUsecase{
		harvestRecordRepo: harvestRecordRepo,
	}
}

func (u *updateHarvestRecordUsecase) Execute(ctx context.Context, req *UpdateHarvestRecordRequest) (*entity.HarvestRecord, error) {
	existingRecord, err := u.harvestRecordRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if existingRecord == nil {
		return nil, ErrHarvestRecordNotFound
	}

	if req.QualityGrade != "" {
		qualityGrade := entity.QualityGrade(req.QualityGrade)
		if !qualityGrade.IsValid() {
			return nil, ErrInvalidQualityGrade
		}
		existingRecord.QualityGrade = req.QualityGrade
	}

	if req.SizeClassification != "" {
		sizeClassification := entity.SizeClassification(req.SizeClassification)
		if !sizeClassification.IsValid() {
			return nil, ErrInvalidSizeClassification
		}
		existingRecord.SizeClassification = req.SizeClassification
	}

	if req.PlantHealthRating > 0 {
		if !entity.IsValidPlantHealthRating(req.PlantHealthRating) {
			return nil, ErrInvalidPlantHealthRating
		}
		existingRecord.PlantHealthRating = req.PlantHealthRating
	}

	existingRecord.TotalRevenue = existingRecord.QuantityKg * existingRecord.MarketPricePerKg

	now := time.Now()
	existingRecord.UpdatedAt = &now

	err = u.harvestRecordRepo.Update(ctx, existingRecord)
	if err != nil {
		return nil, err
	}

	return existingRecord, nil
}
