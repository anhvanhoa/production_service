package pest_disease_record_service

import (
	"context"

	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
)

func (s *PestDiseaseRecordService) DeletePestDiseaseRecord(ctx context.Context, req *proto_pest_disease_record.DeletePestDiseaseRecordRequest) (*proto_pest_disease_record.DeletePestDiseaseRecordResponse, error) {
	err := s.pestDiseaseRecordUsecase.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_pest_disease_record.DeletePestDiseaseRecordResponse{}, nil
}
