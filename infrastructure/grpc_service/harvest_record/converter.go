package harvest_record_service

import (
	"production_service/domain/entity"
	"production_service/domain/usecase/harvest_record"

	"github.com/anhvanhoa/service-core/common"
	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *HarvestRecordService) convertToEntity(harvestRecord *proto_harvest_record.CreateHarvestRecordRequest) *harvest_record.CreateHarvestRecordRequest {
	hr := &harvest_record.CreateHarvestRecordRequest{
		PlantingCycleID:    harvestRecord.PlantingCycleId,
		QuantityKg:         harvestRecord.QuantityKg,
		QualityGrade:       harvestRecord.QualityGrade,
		SizeClassification: harvestRecord.SizeClassification,
		MarketPricePerKg:   harvestRecord.MarketPricePerKg,
		LaborHours:         harvestRecord.LaborHours,
		LaborCost:          harvestRecord.LaborCost,
		PackagingCost:      harvestRecord.PackagingCost,
		StorageLocation:    harvestRecord.StorageLocation,
		StorageTemperature: harvestRecord.StorageTemperature,
		BuyerInformation:   harvestRecord.BuyerInformation,
		WeatherAtHarvest:   harvestRecord.WeatherAtHarvest,
		PlantHealthRating:  int(harvestRecord.PlantHealthRating),
		Notes:              harvestRecord.Notes,
		Images:             harvestRecord.Images,
		CreatedBy:          harvestRecord.CreatedBy,
	}

	if harvestRecord.HarvestDate != nil {
		harvestDate := harvestRecord.HarvestDate.AsTime()
		hr.HarvestDate = &harvestDate
	}
	if harvestRecord.HarvestTime != nil {
		harvestTime := harvestRecord.HarvestTime.AsTime()
		hr.HarvestTime = &harvestTime
	}
	if harvestRecord.DeliveryDate != nil {
		deliveryDate := harvestRecord.DeliveryDate.AsTime()
		hr.DeliveryDate = &deliveryDate
	}
	return hr
}

func (s *HarvestRecordService) convertToUpdateEntity(harvestRecord *proto_harvest_record.UpdateHarvestRecordRequest) *harvest_record.UpdateHarvestRecordRequest {
	hr := &harvest_record.UpdateHarvestRecordRequest{
		ID:                 harvestRecord.Id,
		QuantityKg:         harvestRecord.QuantityKg,
		QualityGrade:       harvestRecord.QualityGrade,
		SizeClassification: harvestRecord.SizeClassification,
		MarketPricePerKg:   harvestRecord.MarketPricePerKg,
		LaborHours:         harvestRecord.LaborHours,
		LaborCost:          harvestRecord.LaborCost,
		PackagingCost:      harvestRecord.PackagingCost,
		StorageLocation:    harvestRecord.StorageLocation,
		StorageTemperature: harvestRecord.StorageTemperature,
		BuyerInformation:   harvestRecord.BuyerInformation,
		WeatherAtHarvest:   harvestRecord.WeatherAtHarvest,
		PlantHealthRating:  int(harvestRecord.PlantHealthRating),
		Notes:              harvestRecord.Notes,
		Images:             harvestRecord.Images,
	}

	if harvestRecord.HarvestDate != nil {
		harvestDate := harvestRecord.HarvestDate.AsTime()
		hr.HarvestDate = &harvestDate
	}
	if harvestRecord.HarvestTime != nil {
		harvestTime := harvestRecord.HarvestTime.AsTime()
		hr.HarvestTime = &harvestTime
	}
	if harvestRecord.DeliveryDate != nil {
		deliveryDate := harvestRecord.DeliveryDate.AsTime()
		hr.DeliveryDate = &deliveryDate
	}
	return hr
}

func (s *HarvestRecordService) convertToProtoHarvestRecord(harvestRecord *entity.HarvestRecord) *proto_harvest_record.HarvestRecord {
	protoRecord := &proto_harvest_record.HarvestRecord{
		Id:                 harvestRecord.ID,
		PlantingCycleId:    harvestRecord.PlantingCycleID,
		QuantityKg:         harvestRecord.QuantityKg,
		QualityGrade:       harvestRecord.QualityGrade,
		SizeClassification: harvestRecord.SizeClassification,
		MarketPricePerKg:   harvestRecord.MarketPricePerKg,
		TotalRevenue:       harvestRecord.TotalRevenue,
		LaborHours:         harvestRecord.LaborHours,
		LaborCost:          harvestRecord.LaborCost,
		PackagingCost:      harvestRecord.PackagingCost,
		StorageLocation:    harvestRecord.StorageLocation,
		StorageTemperature: harvestRecord.StorageTemperature,
		BuyerInformation:   harvestRecord.BuyerInformation,
		WeatherAtHarvest:   harvestRecord.WeatherAtHarvest,
		PlantHealthRating:  int32(harvestRecord.PlantHealthRating),
		Notes:              harvestRecord.Notes,
		Images:             harvestRecord.Images,
		CreatedBy:          harvestRecord.CreatedBy,
		CreatedAt:          timestamppb.New(harvestRecord.CreatedAt),
	}
	if harvestRecord.HarvestDate != nil {
		protoRecord.HarvestDate = timestamppb.New(*harvestRecord.HarvestDate)
	}
	if harvestRecord.HarvestTime != nil {
		protoRecord.HarvestTime = timestamppb.New(*harvestRecord.HarvestTime)
	}
	if harvestRecord.DeliveryDate != nil {
		protoRecord.DeliveryDate = timestamppb.New(*harvestRecord.DeliveryDate)
	}
	if harvestRecord.UpdatedAt != nil {
		protoRecord.UpdatedAt = timestamppb.New(*harvestRecord.UpdatedAt)
	}
	return protoRecord
}

func (s *HarvestRecordService) convertToPagination(pagination *proto_common.PaginationRequest) common.Pagination {
	if pagination == nil {
		return common.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	return common.Pagination{
		Page:      int(pagination.Page),
		PageSize:  int(pagination.PageSize),
		SortBy:    pagination.SortBy,
		SortOrder: pagination.SortOrder,
	}
}

func (s *HarvestRecordService) convertToFilter(filter *proto_harvest_record.HarvestRecordFilter) entity.FilterHarvestRecord {
	if filter == nil {
		return entity.FilterHarvestRecord{}
	}

	entityFilter := entity.FilterHarvestRecord{
		QualityGrade:       filter.QualityGrade,
		SizeClassification: filter.SizeClassification,
		MarketPricePerKg:   filter.MarketPricePerKg,
		TotalRevenue:       filter.TotalRevenue,
		PlantHealthRating:  int(filter.PlantHealthRating),
		Notes:              filter.Notes,
		Images:             filter.Images,
		CreatedBy:          filter.CreatedBy,
	}

	if filter.HarvestDate != nil {
		harvestDate := filter.HarvestDate.AsTime()
		entityFilter.HarvestDate = &harvestDate
	}
	if filter.CreatedAt != nil {
		entityFilter.CreatedAt = filter.CreatedAt.AsTime()
	}

	return entityFilter
}
