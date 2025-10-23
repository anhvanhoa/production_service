package repo

import (
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type RepositoryFactory struct {
	db    *pg.DB
	utils utils.Helper
}

func NewRepositoryFactory(db *pg.DB, utils utils.Helper) *RepositoryFactory {
	return &RepositoryFactory{
		db:    db,
		utils: utils,
	}
}

func (f *RepositoryFactory) NewHarvestRecordRepository() repository.HarvestRecordRepository {
	return NewHarvestRecordRepository(f.db, f.utils)
}

func (f *RepositoryFactory) NewPestDiseaseRecordRepository() repository.PestDiseaseRecordRepository {
	return NewPestDiseaseRecordRepository(f.db, f.utils)
}
