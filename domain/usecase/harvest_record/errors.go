package harvest_record

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrInvalidQualityGrade        = oops.New("Số lượng thu hoạch không hợp lệ")
	ErrInvalidSizeClassification  = oops.New("Phân loại kích thước không hợp lệ")
	ErrInvalidPlantHealthRating   = oops.New("Đánh giá sức khỏe cây trồng không hợp lệ")
	ErrHarvestRecordNotFound      = oops.New("Bản ghi thu hoạch không tồn tại")
	ErrHarvestRecordAlreadyExists = oops.New("Bản ghi thu hoạch đã tồn tại")
	ErrHarvestRecordNotCreated    = oops.New("Bản ghi thu hoạch không được tạo")
	ErrHarvestRecordNotUpdated    = oops.New("Bản ghi thu hoạch không được cập nhật")
	ErrHarvestRecordNotDeleted    = oops.New("Bản ghi thu hoạch không được xóa")
)
