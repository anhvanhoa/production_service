package harvest_record_service

import (
	"context"

	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
)

func (s *HarvestRecordService) ListHarvestRecords(ctx context.Context, req *proto_harvest_record.ListHarvestRecordsRequest) (*proto_harvest_record.ListHarvestRecordsResponse, error) {
	pagination := s.convertToPagination(req.Pagination)
	filter := s.convertToFilter(req.Filter)

	harvestRecords, total, err := s.harvestRecordUsecase.List(ctx, pagination, filter)
	if err != nil {
		return nil, err
	}

	protoRecords := make([]*proto_harvest_record.HarvestRecord, len(harvestRecords))
	for i, record := range harvestRecords {
		protoRecords[i] = s.convertToProtoHarvestRecord(record)
	}

	return &proto_harvest_record.ListHarvestRecordsResponse{
		HarvestRecords: protoRecords,
		Pagination: &proto_common.PaginationResponse{
			Total:      int32(total),
			Page:       int32(pagination.Page),
			TotalPages: int32((total + int64(pagination.PageSize) - 1) / int64(pagination.PageSize)),
		},
	}, nil
}
