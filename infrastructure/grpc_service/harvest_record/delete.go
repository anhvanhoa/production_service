package harvest_record_service

import (
	"context"

	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
)

func (s *HarvestRecordService) DeleteHarvestRecord(ctx context.Context, req *proto_harvest_record.DeleteHarvestRecordRequest) (*proto_harvest_record.DeleteHarvestRecordResponse, error) {
	err := s.harvestRecordUsecase.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_harvest_record.DeleteHarvestRecordResponse{}, nil
}
