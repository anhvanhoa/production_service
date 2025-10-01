package harvest_record_service

import (
	"context"

	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
)

func (s *HarvestRecordService) GetHarvestRecord(ctx context.Context, req *proto_harvest_record.GetHarvestRecordRequest) (*proto_harvest_record.GetHarvestRecordResponse, error) {
	harvestRecord, err := s.harvestRecordUsecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_harvest_record.GetHarvestRecordResponse{HarvestRecord: s.convertToProtoHarvestRecord(harvestRecord)}, nil
}
