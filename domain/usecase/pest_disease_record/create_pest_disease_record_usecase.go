package pest_disease_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"
	"time"
)

type CreatePestDiseaseRecordUsecase interface {
	Execute(ctx context.Context, req *CreatePestDiseaseRecordRequest) (*entity.PestDiseaseRecord, error)
}

type CreatePestDiseaseRecordRequest struct {
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
}

type createPestDiseaseRecordUsecase struct {
	pestDiseaseRecordRepo repository.PestDiseaseRecordRepository
}

func NewCreatePestDiseaseRecordUsecase(pestDiseaseRecordRepo repository.PestDiseaseRecordRepository) CreatePestDiseaseRecordUsecase {
	return &createPestDiseaseRecordUsecase{
		pestDiseaseRecordRepo: pestDiseaseRecordRepo,
	}
}

func (u *createPestDiseaseRecordUsecase) Execute(ctx context.Context, req *CreatePestDiseaseRecordRequest) (*entity.PestDiseaseRecord, error) {
	recordType := entity.RecordType(req.Type)
	if !recordType.IsValid() {
		return nil, ErrInvalidRecordType
	}

	severity := entity.Severity(req.Severity)
	if !severity.IsValid() {
		return nil, ErrInvalidSeverity
	}

	detectionMethod := entity.DetectionMethod(req.DetectionMethod)
	if !detectionMethod.IsValid() {
		return nil, ErrInvalidDetectionMethod
	}

	if req.RecoveryStatus != "" {
		recoveryStatus := entity.RecoveryStatus(req.RecoveryStatus)
		if !recoveryStatus.IsValid() {
			return nil, ErrInvalidRecoveryStatus
		}
	}

	if req.EffectivenessRating > 0 {
		if !entity.IsValidEffectivenessRating(req.EffectivenessRating) {
			return nil, ErrInvalidEffectivenessRating
		}
	}

	pestDiseaseRecord := &entity.PestDiseaseRecord{
		PlantingCycleID:        req.PlantingCycleID,
		Type:                   req.Type,
		Name:                   req.Name,
		ScientificName:         req.ScientificName,
		Severity:               req.Severity,
		AffectedAreaPercentage: req.AffectedAreaPercentage,
		AffectedPlantCount:     req.AffectedPlantCount,
		DetectionDate:          req.DetectionDate,
		DetectionMethod:        req.DetectionMethod,
		Symptoms:               req.Symptoms,
		TreatmentApplied:       req.TreatmentApplied,
		TreatmentDate:          req.TreatmentDate,
		TreatmentCost:          req.TreatmentCost,
		TreatmentDurationDays:  req.TreatmentDurationDays,
		RecoveryStatus:         req.RecoveryStatus,
		EffectivenessRating:    req.EffectivenessRating,
		FollowUpDate:           req.FollowUpDate,
		PreventionMeasures:     req.PreventionMeasures,
		EnvironmentalFactors:   req.EnvironmentalFactors,
		Images:                 req.Images,
		Notes:                  req.Notes,
		CreatedBy:              req.CreatedBy,
		CreatedAt:              time.Now(),
	}

	err := u.pestDiseaseRecordRepo.Create(ctx, pestDiseaseRecord)
	if err != nil {
		return nil, err
	}

	return pestDiseaseRecord, nil
}
