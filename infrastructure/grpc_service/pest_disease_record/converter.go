package pest_disease_record_service

import (
	"production_service/domain/entity"
	"production_service/domain/usecase/pest_disease_record"

	"github.com/anhvanhoa/service-core/common"
	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *PestDiseaseRecordService) convertToEntity(record *proto_pest_disease_record.CreatePestDiseaseRecordRequest) *pest_disease_record.CreatePestDiseaseRecordRequest {
	req := &pest_disease_record.CreatePestDiseaseRecordRequest{
		PlantingCycleID:        record.PlantingCycleId,
		Type:                   record.Type,
		Name:                   record.Name,
		ScientificName:         record.ScientificName,
		Severity:               record.Severity,
		AffectedAreaPercentage: record.AffectedAreaPercentage,
		AffectedPlantCount:     int(record.AffectedPlantCount),
		DetectionMethod:        record.DetectionMethod,
		Symptoms:               record.Symptoms,
		TreatmentApplied:       record.TreatmentApplied,
		TreatmentCost:          record.TreatmentCost,
		TreatmentDurationDays:  int(record.TreatmentDurationDays),
		RecoveryStatus:         record.RecoveryStatus,
		EffectivenessRating:    int(record.EffectivenessRating),
		PreventionMeasures:     record.PreventionMeasures,
		EnvironmentalFactors:   record.EnvironmentalFactors,
		Images:                 record.Images,
		Notes:                  record.Notes,
		CreatedBy:              record.CreatedBy,
	}

	if record.DetectionDate != nil {
		detectionDate := record.DetectionDate.AsTime()
		req.DetectionDate = &detectionDate
	}
	if record.TreatmentDate != nil {
		treatmentDate := record.TreatmentDate.AsTime()
		req.TreatmentDate = &treatmentDate
	}
	if record.FollowUpDate != nil {
		followUpDate := record.FollowUpDate.AsTime()
		req.FollowUpDate = &followUpDate
	}

	return req
}

func (s *PestDiseaseRecordService) convertToUpdateEntity(record *proto_pest_disease_record.UpdatePestDiseaseRecordRequest) *pest_disease_record.UpdatePestDiseaseRecordRequest {
	req := &pest_disease_record.UpdatePestDiseaseRecordRequest{
		ID:                     record.Id,
		Type:                   record.Type,
		Name:                   record.Name,
		ScientificName:         record.ScientificName,
		Severity:               record.Severity,
		AffectedAreaPercentage: record.AffectedAreaPercentage,
		AffectedPlantCount:     int(record.AffectedPlantCount),
		DetectionMethod:        record.DetectionMethod,
		Symptoms:               record.Symptoms,
		TreatmentApplied:       record.TreatmentApplied,
		TreatmentCost:          record.TreatmentCost,
		TreatmentDurationDays:  int(record.TreatmentDurationDays),
		RecoveryStatus:         record.RecoveryStatus,
		EffectivenessRating:    int(record.EffectivenessRating),
		PreventionMeasures:     record.PreventionMeasures,
		EnvironmentalFactors:   record.EnvironmentalFactors,
		Images:                 record.Images,
		Notes:                  record.Notes,
	}

	if record.DetectionDate != nil {
		detectionDate := record.DetectionDate.AsTime()
		req.DetectionDate = &detectionDate
	}
	if record.TreatmentDate != nil {
		treatmentDate := record.TreatmentDate.AsTime()
		req.TreatmentDate = &treatmentDate
	}
	if record.FollowUpDate != nil {
		followUpDate := record.FollowUpDate.AsTime()
		req.FollowUpDate = &followUpDate
	}

	return req
}

func (s *PestDiseaseRecordService) convertToProtoPestDiseaseRecord(record *entity.PestDiseaseRecord) *proto_pest_disease_record.PestDiseaseRecord {
	protoRecord := &proto_pest_disease_record.PestDiseaseRecord{
		Id:                     record.ID,
		PlantingCycleId:        record.PlantingCycleID,
		Type:                   record.Type,
		Name:                   record.Name,
		ScientificName:         record.ScientificName,
		Severity:               record.Severity,
		AffectedAreaPercentage: record.AffectedAreaPercentage,
		AffectedPlantCount:     int32(record.AffectedPlantCount),
		DetectionMethod:        record.DetectionMethod,
		Symptoms:               record.Symptoms,
		TreatmentApplied:       record.TreatmentApplied,
		TreatmentCost:          record.TreatmentCost,
		TreatmentDurationDays:  int32(record.TreatmentDurationDays),
		RecoveryStatus:         record.RecoveryStatus,
		EffectivenessRating:    int32(record.EffectivenessRating),
		PreventionMeasures:     record.PreventionMeasures,
		EnvironmentalFactors:   record.EnvironmentalFactors,
		Images:                 record.Images,
		Notes:                  record.Notes,
		CreatedBy:              record.CreatedBy,
		CreatedAt:              timestamppb.New(record.CreatedAt),
	}

	if record.DetectionDate != nil {
		protoRecord.DetectionDate = timestamppb.New(*record.DetectionDate)
	}
	if record.TreatmentDate != nil {
		protoRecord.TreatmentDate = timestamppb.New(*record.TreatmentDate)
	}
	if record.FollowUpDate != nil {
		protoRecord.FollowUpDate = timestamppb.New(*record.FollowUpDate)
	}
	if record.UpdatedAt != nil {
		protoRecord.UpdatedAt = timestamppb.New(*record.UpdatedAt)
	}

	return protoRecord
}

func (s *PestDiseaseRecordService) convertToPagination(pagination *proto_pest_disease_record.Pagination) common.Pagination {
	if pagination == nil {
		return common.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	return common.Pagination{
		Page:      int(pagination.Page),
		PageSize:  int(pagination.PageSize),
		SortBy:    pagination.SortBy,
		SortOrder: pagination.SortOrder,
	}
}

func (s *PestDiseaseRecordService) convertToFilter(filter *proto_pest_disease_record.PestDiseaseRecordFilter) entity.FilterPestDiseaseRecord {
	if filter == nil {
		return entity.FilterPestDiseaseRecord{}
	}

	entityFilter := entity.FilterPestDiseaseRecord{
		DetectionMethod:       filter.DetectionMethod,
		TreatmentDurationDays: int(filter.TreatmentDurationDays),
	}

	if filter.DetectionDate != nil {
		detectionDate := filter.DetectionDate.AsTime()
		entityFilter.DetectionDate = &detectionDate
	}
	if filter.TreatmentDate != nil {
		treatmentDate := filter.TreatmentDate.AsTime()
		entityFilter.TreatmentDate = &treatmentDate
	}
	if filter.FollowUpDate != nil {
		followUpDate := filter.FollowUpDate.AsTime()
		entityFilter.FollowUpDate = &followUpDate
	}

	return entityFilter
}
