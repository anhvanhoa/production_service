package pest_disease_record_service

import (
	"context"

	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
)

func (s *PestDiseaseRecordService) ListPestDiseaseRecords(ctx context.Context, req *proto_pest_disease_record.ListPestDiseaseRecordsRequest) (*proto_pest_disease_record.ListPestDiseaseRecordsResponse, error) {
	pagination := s.convertToPagination(req.Pagination)
	filter := s.convertToFilter(req.Filter)

	records, total, err := s.pestDiseaseRecordUsecase.List(ctx, pagination, filter)
	if err != nil {
		return nil, err
	}

	protoRecords := make([]*proto_pest_disease_record.PestDiseaseRecord, len(records))
	for i, record := range records {
		protoRecords[i] = s.convertToProtoPestDiseaseRecord(record)
	}

	return &proto_pest_disease_record.ListPestDiseaseRecordsResponse{
		PestDiseaseRecords: protoRecords,
		Total:              total,
	}, nil
}
