package pest_disease_record_service

import (
	"context"

	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
)

func (s *PestDiseaseRecordService) UpdatePestDiseaseRecord(ctx context.Context, req *proto_pest_disease_record.UpdatePestDiseaseRecordRequest) (*proto_pest_disease_record.UpdatePestDiseaseRecordResponse, error) {
	usecaseReq := s.convertToUpdateEntity(req)
	record, err := s.pestDiseaseRecordUsecase.Update(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}
	return &proto_pest_disease_record.UpdatePestDiseaseRecordResponse{PestDiseaseRecord: s.convertToProtoPestDiseaseRecord(record)}, nil
}
