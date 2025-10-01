package pest_disease_record_service

import (
	"context"

	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
)

func (s *PestDiseaseRecordService) CreatePestDiseaseRecord(ctx context.Context, req *proto_pest_disease_record.CreatePestDiseaseRecordRequest) (*proto_pest_disease_record.CreatePestDiseaseRecordResponse, error) {
	usecaseReq := s.convertToEntity(req)
	record, err := s.pestDiseaseRecordUsecase.Create(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}
	return &proto_pest_disease_record.CreatePestDiseaseRecordResponse{PestDiseaseRecord: s.convertToProtoPestDiseaseRecord(record)}, nil
}
