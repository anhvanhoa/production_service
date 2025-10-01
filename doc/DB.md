CREATE TABLE harvest_records (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()), -- ID duy nhất cho bản ghi thu hoạch
    planting_cycle_id VARCHAR(36) NOT NULL,      -- Liên kết tới chu kỳ trồng (planting_cycles)
    harvest_date DATE,                           -- Ngày thu hoạch
    harvest_time TIME,                           -- Thời gian thu hoạch
    quantity_kg DECIMAL(10,3),                   -- Khối lượng thu hoạch (kg)
    quality_grade VARCHAR(10) COMMENT 'A+, A, B, C, Reject', -- Phân loại chất lượng
    size_classification VARCHAR(50) COMMENT 'XL, L, M, S, XS', -- Phân loại kích thước
    market_price_per_kg DECIMAL(10,2),           -- Giá bán thị trường trên mỗi kg
    total_revenue DECIMAL(12,2),                 -- Doanh thu = quantity_kg * market_price_per_kg
    labor_hours DECIMAL(5,2),                    -- Số giờ lao động để thu hoạch
    labor_cost DECIMAL(10,2),                    -- Chi phí nhân công
    packaging_cost DECIMAL(10,2),                -- Chi phí đóng gói
    storage_location VARCHAR(255),               -- Vị trí lưu trữ sau thu hoạch
    storage_temperature DECIMAL(5,2),            -- Nhiệt độ bảo quản (°C)
    buyer_information TEXT,                      -- Thông tin người mua
    delivery_date DATE,                          -- Ngày giao hàng
    weather_at_harvest VARCHAR(200),             -- Thời tiết khi thu hoạch
    plant_health_rating INTEGER COMMENT '1-5 rating', -- Đánh giá sức khỏe cây trồng
    notes TEXT,                                  -- Ghi chú khác
    images JSON COMMENT 'Array of media IDs',    -- Danh sách ảnh/video minh họa
    created_by VARCHAR(36),                      -- Người tạo bản ghi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm tạo
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Thời điểm cập nhật
    
    FOREIGN KEY (planting_cycle_id) REFERENCES planting_cycles(id) ON DELETE CASCADE, -- Khóa ngoại tới planting_cycles
    FOREIGN KEY (created_by) REFERENCES users(id), -- Khóa ngoại tới users
    INDEX idx_harvest_records_cycle (planting_cycle_id), -- Index để truy vấn theo chu kỳ trồng
    INDEX idx_harvest_records_date (harvest_date),      -- Index để truy vấn theo ngày thu hoạch
    INDEX idx_harvest_records_grade (quality_grade)     -- Index để truy vấn theo chất lượng
);


CREATE TABLE pest_disease_records (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()), -- ID duy nhất cho bản ghi sâu bệnh
    planting_cycle_id VARCHAR(36) NOT NULL,      -- Liên kết tới chu kỳ trồng
    type VARCHAR(50) COMMENT 'pest, disease, nutrient_deficiency, environmental_stress', -- Loại vấn đề
    name VARCHAR(255),                           -- Tên sâu bệnh / vấn đề
    scientific_name VARCHAR(255),                -- Tên khoa học
    severity VARCHAR(50) COMMENT 'low, medium, high, critical', -- Mức độ nghiêm trọng
    affected_area_percentage DECIMAL(5,2),       -- % diện tích bị ảnh hưởng
    affected_plant_count INTEGER,                -- Số cây bị ảnh hưởng
    detection_date DATE,                         -- Ngày phát hiện
    detection_method VARCHAR(100) COMMENT 'visual, trap, sensor, lab_test', -- Phương pháp phát hiện
    symptoms TEXT,                               -- Triệu chứng
    treatment_applied VARCHAR(500),              -- Biện pháp xử lý đã áp dụng
    treatment_date DATE,                         -- Ngày xử lý
    treatment_cost DECIMAL(10,2),                -- Chi phí xử lý
    treatment_duration_days INTEGER,             -- Số ngày điều trị
    recovery_status VARCHAR(50) COMMENT 'treating, recovering, recovered, failed, spreading', -- Tình trạng phục hồi
    effectiveness_rating INTEGER COMMENT '1-5 rating của treatment', -- Đánh giá hiệu quả biện pháp xử lý
    follow_up_date DATE,                         -- Ngày theo dõi tiếp theo
    prevention_measures TEXT,                    -- Biện pháp phòng ngừa
    environmental_factors TEXT,                  -- Yếu tố môi trường liên quan
    images JSON COMMENT 'Array of media IDs',    -- Hình ảnh/video minh chứng
    notes TEXT,                                  -- Ghi chú khác
    created_by VARCHAR(36),                      -- Người tạo bản ghi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm tạo
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Thời điểm cập nhật
    
    FOREIGN KEY (planting_cycle_id) REFERENCES planting_cycles(id) ON DELETE CASCADE, -- Khóa ngoại tới planting_cycles
    FOREIGN KEY (created_by) REFERENCES users(id), -- Khóa ngoại tới users
    INDEX idx_pest_disease_cycle (planting_cycle_id),         -- Index theo chu kỳ trồng
    INDEX idx_pest_disease_type_severity (type, severity),    -- Index theo loại và mức độ
    INDEX idx_pest_disease_detection_date (detection_date),   -- Index theo ngày phát hiện
    INDEX idx_pest_disease_status (recovery_status)           -- Index theo tình trạng phục hồi
);