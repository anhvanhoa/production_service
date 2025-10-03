# Production Service

Microservice quản lý bản ghi thu hoạch và sâu bệnh trong quy trình sản xuất nông nghiệp, được xây dựng bằng Go và tuân theo nguyên tắc Clean Architecture.

## 🏗️ Kiến trúc

Dự án này tuân theo **Clean Architecture** với sự phân tách rõ ràng các mối quan tâm:

```
├── domain/           # Tầng logic nghiệp vụ
│   ├── entity/       # Các thực thể nghiệp vụ cốt lõi
│   │   ├── harvest_record.go         # Entity bản ghi thu hoạch
│   │   └── pest_disease_record.go    # Entity bản ghi sâu bệnh
│   ├── repository/   # Giao diện truy cập dữ liệu
│   │   ├── harvest_record_repository.go
│   │   └── pest_disease_record_repository.go
│   └── usecase/      # Các trường hợp sử dụng nghiệp vụ
│       ├── harvest_record/         # Use cases bản ghi thu hoạch
│       └── pest_disease_record/    # Use cases bản ghi sâu bệnh
├── infrastructure/   # Các mối quan tâm bên ngoài
│   ├── grpc_service/ # Triển khai API gRPC
│   │   ├── harvest_record/         # gRPC handlers bản ghi thu hoạch
│   │   └── pest_disease_record/    # gRPC handlers bản ghi sâu bệnh
│   └── repo/         # Triển khai repository cơ sở dữ liệu
├── bootstrap/        # Khởi tạo ứng dụng
└── cmd/             # Điểm vào ứng dụng
```

## 🚀 Tính năng

### Quản lý Bản ghi Thu hoạch (Harvest Record)
- ✅ Tạo, đọc, cập nhật, xóa bản ghi thu hoạch
- ✅ Liệt kê bản ghi theo bộ lọc và phân trang
- ✅ Lấy bản ghi theo chu kỳ trồng
- ✅ Tính toán thông tin phân trang trả về

### Quản lý Bản ghi Sâu bệnh (Pest/Disease Record)
- ✅ Tạo, đọc, cập nhật, xóa bản ghi sâu bệnh
- ✅ Liệt kê bản ghi theo bộ lọc và phân trang
- ✅ Lấy bản ghi theo chu kỳ trồng

## 🛠️ Công nghệ sử dụng

- **Ngôn ngữ**: Go 1.24.6
- **Cơ sở dữ liệu**: PostgreSQL
- **API**: gRPC
- **Kiến trúc**: Clean Architecture
- **Dependencies**:
  - `github.com/go-pg/pg/v10` - PostgreSQL ORM
  - `google.golang.org/grpc` - Framework gRPC
  - `github.com/spf13/viper` - Quản lý cấu hình
  - `go.uber.org/zap` - Logging có cấu trúc

## 📋 Yêu cầu hệ thống

- Go 1.24.6 trở lên
- PostgreSQL 12 trở lên
- [golang-migrate](https://github.com/golang-migrate/migrate) để quản lý migration cơ sở dữ liệu

## 🚀 Hướng dẫn nhanh

### 1. Clone repository
```bash
git clone <repository-url>
cd production_service
```

### 2. Cài đặt dependencies
```bash
go mod download
```

### 3. Thiết lập cơ sở dữ liệu
```bash
# Tạo cơ sở dữ liệu
make create-db

# Chạy migrations
make up
```

### 4. Cấu hình ứng dụng
Sao chép và chỉnh sửa file cấu hình:
```bash
cp dev.config.yml config.yml
```

Cập nhật chuỗi kết nối cơ sở dữ liệu trong `config.yml`:
```yaml
node_env: "development"
url_db: "postgres://postgres:123456@localhost:5432/production_service_db?sslmode=disable"
name_service: "ProductionService"
port_grpc: 50056
host_grpc: "localhost"
interval_check: "20s"
timeout_check: "15s"
```

### 5. Chạy ứng dụng
```bash
# Build và chạy service chính
make run

# Hoặc chạy client để test
make client
```

## 🗄️ Quản lý Cơ sở dữ liệu

Dự án sử dụng `golang-migrate` để quản lý schema cơ sở dữ liệu:

```bash
# Chạy tất cả migrations đang chờ
make up

# Rollback migration cuối cùng
make down

# Reset cơ sở dữ liệu hoàn toàn
make reset

# Tạo migration mới
make create name=migration_name

# Force migration đến phiên bản cụ thể
make force v=1
```

## 🌱 Dữ liệu mẫu

Dự án bao gồm dữ liệu mẫu để phát triển và kiểm thử:

```bash
# Chèn dữ liệu mẫu (tham số bắt buộc: up hoặc down)
go run script/seed/main.go up
# hoặc
go run script/seed/main.go down
```

### Dữ liệu mẫu bao gồm:

**Bản ghi Thu hoạch và Sâu bệnh mẫu:**
- Nhiều bản ghi thu hoạch với thông tin chất lượng, giá, doanh thu, chi phí, điều kiện bảo quản...
- Nhiều bản ghi sâu bệnh với thông tin loại, mức độ, triệu chứng, điều trị, hiệu quả...

## 📁 Cấu trúc Dự án

```
production_service/
├── bootstrap/               # Khởi tạo ứng dụng
│   ├── app.go               # Khởi tạo app
│   └── env.go               # Cấu hình môi trường
├── cmd/                     # Điểm vào ứng dụng
│   ├── main.go              # Điểm vào service chính
│   └── client/              # gRPC client để test
├── domain/                  # Logic nghiệp vụ (Clean Architecture)
│   ├── entity/              # Các thực thể nghiệp vụ cốt lõi
│   │   ├── harvest_record.go        # Entity bản ghi thu hoạch
│   │   └── pest_disease_record.go   # Entity bản ghi sâu bệnh
│   ├── repository/          # Giao diện truy cập dữ liệu
│   │   ├── harvest_record_repository.go
│   │   └── pest_disease_record_repository.go
│   └── usecase/             # Các trường hợp sử dụng nghiệp vụ
│       ├── harvest_record/         # Use cases bản ghi thu hoạch
│       └── pest_disease_record/    # Use cases bản ghi sâu bệnh
├── infrastructure/          # Các mối quan tâm bên ngoài
│   ├── grpc_service/        # Triển khai API gRPC
│   │   ├── harvest_record/         # gRPC handlers bản ghi thu hoạch
│   │   ├── pest_disease_record/    # gRPC handlers bản ghi sâu bệnh
│   │   └── sesrver.go              # Thiết lập gRPC server
│   └── repo/                # Triển khai cơ sở dữ liệu
│       ├── harvest_record_repository.go
│       ├── pest_disease_record_repository.go
│       └── repository_factory.go
├── migrations/              # Database migrations
│   ├── 000000_common.up.sql
│   ├── 000002_create_harvest_records.up.sql
│   ├── 000003_create_pest_disease_records.up.sql
│   └── seed/                # Dữ liệu mẫu
├── script/seed/             # Script chèn dữ liệu mẫu
├── doc/                     # Tài liệu
└── logs/                    # Log ứng dụng
```

## 🔧 Các lệnh có sẵn

```bash
# Thao tác cơ sở dữ liệu
make up              # Chạy migrations
make down            # Rollback migration
make reset           # Reset cơ sở dữ liệu
make create-db       # Tạo cơ sở dữ liệu
make drop-db         # Xóa cơ sở dữ liệu

# Ứng dụng
make build           # Build ứng dụng
make run             # Chạy service chính
make client          # Chạy client test
make test            # Chạy tests client

# Trợ giúp
make help            # Hiển thị tất cả lệnh có sẵn
```

## 📊 Mô hình Dữ liệu

### Bản ghi Thu hoạch (Harvest Record)
- **ID**
- **PlantingCycleID**
- **HarvestDate/HarvestTime**
- **QuantityKg, QualityGrade, SizeClassification**
- **MarketPricePerKg, TotalRevenue**
- **LaborHours, LaborCost, PackagingCost**
- **StorageLocation, StorageTemperature**
- **BuyerInformation, DeliveryDate**
- **WeatherAtHarvest, PlantHealthRating**
- **Notes, Images, CreatedBy, CreatedAt, UpdatedAt**

### Bản ghi Sâu bệnh (Pest/Disease Record)
- **ID**
- **PlantingCycleID**
- **Type, Name, ScientificName**
- **Severity, AffectedAreaPercentage, AffectedPlantCount**
- **DetectionDate, DetectionMethod**
- **Symptoms**
- **TreatmentApplied, TreatmentDate, TreatmentCost, TreatmentDurationDays**
- **RecoveryStatus, EffectivenessRating, FollowUpDate**
- **PreventionMeasures, EnvironmentalFactors**
- **Images, Notes, CreatedBy, CreatedAt, UpdatedAt**

## 🔌 API Endpoints

Service cung cấp các endpoint gRPC:

### Harvest Record Service
- `CreateHarvestRecord`
- `GetHarvestRecord`
- `UpdateHarvestRecord`
- `DeleteHarvestRecord`
- `ListHarvestRecords`
- `GetHarvestRecordsByPlantingCycle`

### Pest Disease Record Service
- `CreatePestDiseaseRecord`
- `GetPestDiseaseRecord`
- `UpdatePestDiseaseRecord`
- `DeletePestDiseaseRecord`
- `ListPestDiseaseRecords`
- `GetPestDiseaseRecordsByPlantingCycle`

## 🧪 Testing

Chạy client test để tương tác với service:

```bash
make client
```

Điều này sẽ khởi động một client tương tác nơi bạn có thể test tất cả các endpoint gRPC.

## 📝 Cấu hình

Ứng dụng sử dụng Viper để quản lý cấu hình. Các tùy chọn cấu hình chính:

- `node_env`: Môi trường (development, production)
- `url_db`: Chuỗi kết nối PostgreSQL
- `name_service`: Tên service cho discovery
- `port_grpc`: Cổng gRPC server
- `host_grpc`: Host gRPC server
- `interval_check`: Khoảng thời gian kiểm tra sức khỏe
- `timeout_check`: Timeout kiểm tra sức khỏe

## 🚀 Triển khai

1. **Build ứng dụng**:
   ```bash
   make build
   ```

2. **Thiết lập cơ sở dữ liệu production**:
   ```bash
   make create-db
   make up
   ```

3. **Chạy service**:
   ```bash
   ./bin/app
   ```

## 🤝 Đóng góp

1. Fork repository
2. Tạo feature branch
3. Thực hiện thay đổi
4. Thêm tests nếu cần thiết
5. Submit pull request

## 📄 Giấy phép

Dự án này được cấp phép theo MIT License.

## 🆘 Hỗ trợ

Để được hỗ trợ và đặt câu hỏi, vui lòng tạo issue trong repository.

---

**Lưu ý**: Service này được thiết kế để quản lý bản ghi thu hoạch và sâu bệnh trong quy trình sản xuất nông nghiệp, tuân theo các nguyên tắc kiến trúc microservice để có thể mở rộng và bảo trì dễ dàng.
