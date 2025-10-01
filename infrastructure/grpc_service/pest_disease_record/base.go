package pest_disease_record_service

import (
	"production_service/domain/usecase/pest_disease_record"
	"production_service/infrastructure/repo"

	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
)

type PestDiseaseRecordService struct {
	proto_pest_disease_record.UnsafePestDiseaseRecordServiceServer
	pestDiseaseRecordUsecase pest_disease_record.PestDiseaseRecordUsecase
}

func NewPestDiseaseRecordService(repos *repo.RepositoryFactory) proto_pest_disease_record.PestDiseaseRecordServiceServer {
	pestDiseaseRecordUsecase := pest_disease_record.NewPestDiseaseRecordUsecase(repos.NewPestDiseaseRecordRepository())
	return &PestDiseaseRecordService{pestDiseaseRecordUsecase: pestDiseaseRecordUsecase}
}
