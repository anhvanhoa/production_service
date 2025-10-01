package entity

import (
	"time"
)

type PestDiseaseRecord struct {
	tableName              struct{} `pg:"pest_disease_records"`
	ID                     string
	PlantingCycleID        string
	Type                   string
	Name                   string
	ScientificName         string
	Severity               string
	AffectedAreaPercentage float64
	AffectedPlantCount     int
	DetectionDate          *time.Time
	DetectionMethod        string
	Symptoms               string
	TreatmentApplied       string
	TreatmentDate          *time.Time
	TreatmentCost          float64
	TreatmentDurationDays  int
	RecoveryStatus         string
	EffectivenessRating    int
	FollowUpDate           *time.Time
	PreventionMeasures     string
	EnvironmentalFactors   string
	Images                 string
	Notes                  string
	CreatedBy              string
	CreatedAt              time.Time
	UpdatedAt              *time.Time
}

func (p *PestDiseaseRecord) TableName() any {
	return p.tableName
}

type FilterPestDiseaseRecord struct {
	DetectionDate         *time.Time
	DetectionMethod       string
	TreatmentDate         *time.Time
	TreatmentDurationDays int
	FollowUpDate          *time.Time
}

// RecordType represents the type of record enum
type RecordType string

const (
	RecordTypePest                RecordType = "pest"
	RecordTypeDisease             RecordType = "disease"
	RecordTypeNutrientDeficiency  RecordType = "nutrient_deficiency"
	RecordTypeEnvironmentalStress RecordType = "environmental_stress"
)

// Severity represents the severity level enum
type Severity string

const (
	SeverityLow      Severity = "low"
	SeverityMedium   Severity = "medium"
	SeverityHigh     Severity = "high"
	SeverityCritical Severity = "critical"
)

// DetectionMethod represents the detection method enum
type DetectionMethod string

const (
	DetectionMethodVisual  DetectionMethod = "visual"
	DetectionMethodTrap    DetectionMethod = "trap"
	DetectionMethodSensor  DetectionMethod = "sensor"
	DetectionMethodLabTest DetectionMethod = "lab_test"
)

// RecoveryStatus represents the recovery status enum
type RecoveryStatus string

const (
	RecoveryStatusTreating   RecoveryStatus = "treating"
	RecoveryStatusRecovering RecoveryStatus = "recovering"
	RecoveryStatusRecovered  RecoveryStatus = "recovered"
	RecoveryStatusFailed     RecoveryStatus = "failed"
	RecoveryStatusSpreading  RecoveryStatus = "spreading"
)

// IsValidRecordType checks if the record type is valid
func (rt RecordType) IsValid() bool {
	switch rt {
	case RecordTypePest, RecordTypeDisease, RecordTypeNutrientDeficiency, RecordTypeEnvironmentalStress:
		return true
	default:
		return false
	}
}

// IsValidSeverity checks if the severity is valid
func (s Severity) IsValid() bool {
	switch s {
	case SeverityLow, SeverityMedium, SeverityHigh, SeverityCritical:
		return true
	default:
		return false
	}
}

// IsValidDetectionMethod checks if the detection method is valid
func (dm DetectionMethod) IsValid() bool {
	switch dm {
	case DetectionMethodVisual, DetectionMethodTrap, DetectionMethodSensor, DetectionMethodLabTest:
		return true
	default:
		return false
	}
}

// IsValidRecoveryStatus checks if the recovery status is valid
func (rs RecoveryStatus) IsValid() bool {
	switch rs {
	case RecoveryStatusTreating, RecoveryStatusRecovering, RecoveryStatusRecovered, RecoveryStatusFailed, RecoveryStatusSpreading:
		return true
	default:
		return false
	}
}

// IsValidEffectivenessRating checks if the effectiveness rating is valid (1-5)
func IsValidEffectivenessRating(rating int) bool {
	return rating >= 1 && rating <= 5
}
