package harvest_record_service

import (
	"context"

	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
)

func (s *HarvestRecordService) CreateHarvestRecord(ctx context.Context, req *proto_harvest_record.CreateHarvestRecordRequest) (*proto_harvest_record.CreateHarvestRecordResponse, error) {
	usecaseReq := s.convertToEntity(req)
	harvestRecord, err := s.harvestRecordUsecase.Create(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}
	return &proto_harvest_record.CreateHarvestRecordResponse{HarvestRecord: s.convertToProtoHarvestRecord(harvestRecord)}, nil
}
