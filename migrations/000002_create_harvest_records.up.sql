CREATE TABLE harvest_records (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(), -- ID duy nhất cho bản ghi thu hoạch
    planting_cycle_id VARCHAR(36) NOT NULL,      -- Liên kết tới chu kỳ trồng (planting_cycles)
    harvest_date DATE,                           -- Ngày thu hoạch
    harvest_time TIME,                           -- Thời gian thu hoạch
    quantity_kg DECIMAL(10,3),                   -- Khối lượng thu hoạch (kg)
    quality_grade VARCHAR(10) CHECK (quality_grade IN ('A+', 'A', 'B', 'C', 'Reject')), -- Phân loại chất lượng
    size_classification VARCHAR(50), -- Phân loại kích thước
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
    plant_health_rating INTEGER CHECK (plant_health_rating >= 1 AND plant_health_rating <= 5), -- Đánh giá sức khỏe cây trồng
    notes TEXT,                                  -- Ghi chú khác
    images JSONB,                                -- Danh sách ảnh/video minh họa
    created_by VARCHAR(36),                      -- Người tạo bản ghi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm tạo
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Thời điểm cập nhật
);

-- Tạo trigger để tự động cập nhật updated_at
CREATE TRIGGER trigger_harvest_records_updated_at
    BEFORE UPDATE ON harvest_records
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Tạo các index
CREATE INDEX idx_harvest_records_cycle ON harvest_records(planting_cycle_id);
CREATE INDEX idx_harvest_records_date ON harvest_records(harvest_date);
CREATE INDEX idx_harvest_records_grade ON harvest_records(quality_grade);

-- Thêm comment cho bảng
COMMENT ON TABLE harvest_records IS 'Bảng lưu trữ thông tin thu hoạch';
COMMENT ON COLUMN harvest_records.quality_grade IS 'Phân loại chất lượng: A+, A, B, C, Reject';
COMMENT ON COLUMN harvest_records.size_classification IS 'Phân loại kích thước: XL, L, M, S, XS';
COMMENT ON COLUMN harvest_records.plant_health_rating IS 'Đánh giá sức khỏe cây trồng từ 1-5';
COMMENT ON COLUMN harvest_records.images IS 'Array of media IDs';
