package pest_disease_record_service

import (
	"context"

	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
)

func (s *PestDiseaseRecordService) GetPestDiseaseRecord(ctx context.Context, req *proto_pest_disease_record.GetPestDiseaseRecordRequest) (*proto_pest_disease_record.GetPestDiseaseRecordResponse, error) {
	record, err := s.pestDiseaseRecordUsecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_pest_disease_record.GetPestDiseaseRecordResponse{PestDiseaseRecord: s.convertToProtoPestDiseaseRecord(record)}, nil
}
