package pest_disease_record

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"
	"time"
)

type UpdatePestDiseaseRecordUsecase interface {
	Execute(ctx context.Context, req *UpdatePestDiseaseRecordRequest) (*entity.PestDiseaseRecord, error)
}

type UpdatePestDiseaseRecordRequest struct {
	ID                     string
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
}

type updatePestDiseaseRecordUsecase struct {
	pestDiseaseRecordRepo repository.PestDiseaseRecordRepository
}

func NewUpdatePestDiseaseRecordUsecase(pestDiseaseRecordRepo repository.PestDiseaseRecordRepository) UpdatePestDiseaseRecordUsecase {
	return &updatePestDiseaseRecordUsecase{
		pestDiseaseRecordRepo: pestDiseaseRecordRepo,
	}
}

func (u *updatePestDiseaseRecordUsecase) Execute(ctx context.Context, req *UpdatePestDiseaseRecordRequest) (*entity.PestDiseaseRecord, error) {
	existingRecord, err := u.pestDiseaseRecordRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if existingRecord == nil {
		return nil, ErrPestDiseaseRecordNotFound
	}

	if req.Type != "" {
		recordType := entity.RecordType(req.Type)
		if !recordType.IsValid() {
			return nil, ErrInvalidRecordType
		}
		existingRecord.Type = req.Type
	}

	if req.Severity != "" {
		severity := entity.Severity(req.Severity)
		if !severity.IsValid() {
			return nil, ErrInvalidSeverity
		}
		existingRecord.Severity = req.Severity
	}

	if req.DetectionMethod != "" {
		detectionMethod := entity.DetectionMethod(req.DetectionMethod)
		if !detectionMethod.IsValid() {
			return nil, ErrInvalidDetectionMethod
		}
		existingRecord.DetectionMethod = req.DetectionMethod
	}

	if req.RecoveryStatus != "" {
		recoveryStatus := entity.RecoveryStatus(req.RecoveryStatus)
		if !recoveryStatus.IsValid() {
			return nil, ErrInvalidRecoveryStatus
		}
		existingRecord.RecoveryStatus = req.RecoveryStatus
	}

	if req.EffectivenessRating > 0 {
		if !entity.IsValidEffectivenessRating(req.EffectivenessRating) {
			return nil, ErrInvalidEffectivenessRating
		}
		existingRecord.EffectivenessRating = req.EffectivenessRating
	}

	now := time.Now()
	existingRecord.UpdatedAt = &now

	err = u.pestDiseaseRecordRepo.Update(ctx, existingRecord)
	if err != nil {
		return nil, err
	}

	return existingRecord, nil
}
