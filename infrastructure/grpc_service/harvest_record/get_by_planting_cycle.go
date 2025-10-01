package harvest_record_service

import (
	"context"

	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
)

func (s *HarvestRecordService) GetHarvestRecordsByPlantingCycle(ctx context.Context, req *proto_harvest_record.GetHarvestRecordsByPlantingCycleRequest) (*proto_harvest_record.GetHarvestRecordsByPlantingCycleResponse, error) {
	pagination := s.convertToPagination(req.Pagination)
	filter := s.convertToFilter(req.Filter)

	harvestRecords, total, err := s.harvestRecordUsecase.GetByPlantingCycleID(ctx, req.PlantingCycleId, pagination, filter)
	if err != nil {
		return nil, err
	}

	protoRecords := make([]*proto_harvest_record.HarvestRecord, len(harvestRecords))
	for i, record := range harvestRecords {
		protoRecords[i] = s.convertToProtoHarvestRecord(record)
	}

	return &proto_harvest_record.GetHarvestRecordsByPlantingCycleResponse{
		HarvestRecords: protoRecords,
		Total:          total,
	}, nil
}
