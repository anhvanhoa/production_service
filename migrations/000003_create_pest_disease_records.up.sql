CREATE TABLE pest_disease_records (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(), -- ID duy nhất cho bản ghi sâu bệnh
    planting_cycle_id VARCHAR(36) NOT NULL,      -- Liên kết tới chu kỳ trồng
    type VARCHAR(50) CHECK (type IN ('pest', 'disease', 'nutrient_deficiency', 'environmental_stress')), -- Loại vấn đề
    name VARCHAR(255),                           -- Tên sâu bệnh / vấn đề
    scientific_name VARCHAR(255),                -- Tên khoa học
    severity VARCHAR(50) CHECK (severity IN ('low', 'medium', 'high', 'critical')), -- Mức độ nghiêm trọng
    affected_area_percentage DECIMAL(5,2),       -- % diện tích bị ảnh hưởng
    affected_plant_count INTEGER,                -- Số cây bị ảnh hưởng
    detection_date DATE,                         -- Ngày phát hiện
    detection_method VARCHAR(100) CHECK (detection_method IN ('visual', 'trap', 'sensor', 'lab_test')), -- Phương pháp phát hiện
    symptoms TEXT,                               -- Triệu chứng
    treatment_applied VARCHAR(500),              -- Biện pháp xử lý đã áp dụng
    treatment_date DATE,                         -- Ngày xử lý
    treatment_cost DECIMAL(10,2),                -- Chi phí xử lý
    treatment_duration_days INTEGER,             -- Số ngày điều trị
    recovery_status VARCHAR(50) CHECK (recovery_status IN ('treating', 'recovering', 'recovered', 'failed', 'spreading')), -- Tình trạng phục hồi
    effectiveness_rating INTEGER CHECK (effectiveness_rating >= 1 AND effectiveness_rating <= 5), -- Đánh giá hiệu quả biện pháp xử lý
    follow_up_date DATE,                         -- Ngày theo dõi tiếp theo
    prevention_measures TEXT,                    -- Biện pháp phòng ngừa
    environmental_factors TEXT,                  -- Yếu tố môi trường liên quan
    images JSONB,                                -- Hình ảnh/video minh chứng
    notes TEXT,                                  -- Ghi chú khác
    created_by VARCHAR(36),                      -- Người tạo bản ghi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm tạo
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Thời điểm cập nhật
);

-- Tạo trigger để tự động cập nhật updated_at
CREATE TRIGGER trigger_pest_disease_records_updated_at
    BEFORE UPDATE ON pest_disease_records
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Tạo các index
CREATE INDEX idx_pest_disease_cycle ON pest_disease_records(planting_cycle_id);
CREATE INDEX idx_pest_disease_type_severity ON pest_disease_records(type, severity);
CREATE INDEX idx_pest_disease_detection_date ON pest_disease_records(detection_date);
CREATE INDEX idx_pest_disease_status ON pest_disease_records(recovery_status);

-- Thêm comment cho bảng
COMMENT ON TABLE pest_disease_records IS 'Bảng lưu trữ thông tin sâu bệnh và các vấn đề khác';
COMMENT ON COLUMN pest_disease_records.type IS 'Loại vấn đề: pest, disease, nutrient_deficiency, environmental_stress';
COMMENT ON COLUMN pest_disease_records.severity IS 'Mức độ nghiêm trọng: low, medium, high, critical';
COMMENT ON COLUMN pest_disease_records.detection_method IS 'Phương pháp phát hiện: visual, trap, sensor, lab_test';
COMMENT ON COLUMN pest_disease_records.recovery_status IS 'Tình trạng phục hồi: treating, recovering, recovered, failed, spreading';
COMMENT ON COLUMN pest_disease_records.effectiveness_rating IS 'Đánh giá hiệu quả biện pháp xử lý từ 1-5';
COMMENT ON COLUMN pest_disease_records.images IS 'Array of media IDs';
