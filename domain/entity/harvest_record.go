package entity

import (
	"time"
)

type HarvestRecord struct {
	tableName          struct{} `pg:"harvest_records"`
	ID                 string
	PlantingCycleID    string
	HarvestDate        *time.Time
	HarvestTime        *time.Time
	QuantityKg         float64
	QualityGrade       string
	SizeClassification string
	MarketPricePerKg   float64
	TotalRevenue       float64
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
	CreatedAt          time.Time
	UpdatedAt          *time.Time
}

func (h *HarvestRecord) TableName() any {
	return h.tableName
}

type FilterHarvestRecord struct {
	HarvestDate        *time.Time
	QualityGrade       string
	SizeClassification string
	MarketPricePerKg   float64
	TotalRevenue       float64
	PlantHealthRating  int
	Notes              string
	Images             string
	CreatedBy          string
	CreatedAt          time.Time
}

type QualityGrade string

const (
	QualityGradeAPlus  QualityGrade = "A+"
	QualityGradeA      QualityGrade = "A"
	QualityGradeB      QualityGrade = "B"
	QualityGradeC      QualityGrade = "C"
	QualityGradeReject QualityGrade = "Reject"
)

// SizeClassification represents the size classification enum
type SizeClassification string

const (
	SizeXL SizeClassification = "XL"
	SizeL  SizeClassification = "L"
	SizeM  SizeClassification = "M"
	SizeS  SizeClassification = "S"
	SizeXS SizeClassification = "XS"
)

// IsValidQualityGrade checks if the quality grade is valid
func (qg QualityGrade) IsValid() bool {
	switch qg {
	case QualityGradeAPlus, QualityGradeA, QualityGradeB, QualityGradeC, QualityGradeReject:
		return true
	default:
		return false
	}
}

// IsValidSizeClassification checks if the size classification is valid
func (sc SizeClassification) IsValid() bool {
	switch sc {
	case SizeXL, SizeL, SizeM, SizeS, SizeXS:
		return true
	default:
		return false
	}
}

// IsValidPlantHealthRating checks if the plant health rating is valid (1-5)
func IsValidPlantHealthRating(rating int) bool {
	return rating >= 1 && rating <= 5
}
