package harvest_record_service

import (
	"production_service/domain/usecase/harvest_record"
	"production_service/infrastructure/repo"

	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
)

type HarvestRecordService struct {
	harvestRecordUsecase harvest_record.HarvestRecordUsecase
	proto_harvest_record.UnsafeHarvestRecordServiceServer
}

func NewHarvestRecordService(repos *repo.RepositoryFactory) proto_harvest_record.HarvestRecordServiceServer {
	harvestRecordUsecase := harvest_record.NewHarvestRecordUsecase(repos.NewHarvestRecordRepository())
	return &HarvestRecordService{harvestRecordUsecase: harvestRecordUsecase}
}
