package harvest_record_service

import (
	"context"

	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
)

func (s *HarvestRecordService) UpdateHarvestRecord(ctx context.Context, req *proto_harvest_record.UpdateHarvestRecordRequest) (*proto_harvest_record.UpdateHarvestRecordResponse, error) {
	usecaseReq := s.convertToUpdateEntity(req)
	harvestRecord, err := s.harvestRecordUsecase.Update(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}
	return &proto_harvest_record.UpdateHarvestRecordResponse{HarvestRecord: s.convertToProtoHarvestRecord(harvestRecord)}, nil
}
