package pest_disease_record

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrInvalidRecordType              = oops.New("Loại vấn đề không hợp lệ")
	ErrInvalidSeverity                = oops.New("Mức độ nghiêm trọng không hợp lệ")
	ErrInvalidDetectionMethod         = oops.New("Phương pháp phát hiện không hợp lệ")
	ErrInvalidRecoveryStatus          = oops.New("Tình trạng phục hồi không hợp lệ")
	ErrInvalidEffectivenessRating     = oops.New("Đánh giá hiệu quả biện pháp xử lý không hợp lệ")
	ErrPestDiseaseRecordNotFound      = oops.New("Bản ghi sâu bệnh không tồn tại")
	ErrPestDiseaseRecordAlreadyExists = oops.New("Bản ghi sâu bệnh đã tồn tại")
	ErrPestDiseaseRecordNotCreated    = oops.New("Bản ghi sâu bệnh không được tạo")
	ErrPestDiseaseRecordNotUpdated    = oops.New("Bản ghi sâu bệnh không được cập nhật")
	ErrPestDiseaseRecordNotDeleted    = oops.New("Bản ghi sâu bệnh không được xóa")
)
