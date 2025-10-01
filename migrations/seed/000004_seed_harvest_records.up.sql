-- Seed data for harvest_records table
INSERT INTO harvest_records (
    id, planting_cycle_id, harvest_date, harvest_time, quantity_kg, 
    quality_grade, size_classification, market_price_per_kg, total_revenue,
    labor_hours, labor_cost, packaging_cost, storage_location, 
    storage_temperature, buyer_information, delivery_date, weather_at_harvest,
    plant_health_rating, notes, images, created_by
) VALUES 
(
    'hr-001', 'pc-001', '2024-01-15', '08:30:00', 150.500,
    'A+', 'L', 25000.00, 3762500.00,
    8.5, 425000.00, 75000.00, 'Kho lạnh A1',
    4.5, 'Công ty TNHH Thực phẩm Xanh', '2024-01-16', 'Nắng nhẹ, ít gió',
    5, 'Thu hoạch đúng thời điểm, chất lượng tốt', 
    '["img-001", "img-002", "img-003"]', 'user-001'
),
(
    'hr-002', 'pc-001', '2024-01-20', '07:45:00', 120.250,
    'A', 'M', 23000.00, 2765750.00,
    6.0, 300000.00, 60000.00, 'Kho lạnh A1',
    4.0, 'Siêu thị Metro', '2024-01-21', 'Nắng, nhiệt độ cao',
    4, 'Một số quả bị nứt do nhiệt độ cao', 
    '["img-004", "img-005"]', 'user-001'
),
(
    'hr-003', 'pc-002', '2024-02-10', '09:15:00', 200.750,
    'A+', 'XL', 28000.00, 5621000.00,
    12.0, 600000.00, 100000.00, 'Kho lạnh B2',
    3.5, 'Nhà hàng Hải Sản', '2024-02-11', 'Mát, độ ẩm cao',
    5, 'Thu hoạch xuất sắc, kích thước đồng đều', 
    '["img-006", "img-007", "img-008", "img-009"]', 'user-002'
),
(
    'hr-004', 'pc-002', '2024-02-15', '08:00:00', 85.300,
    'B', 'S', 20000.00, 1706000.00,
    4.5, 225000.00, 40000.00, 'Kho lạnh B2',
    4.0, 'Chợ đầu mối', '2024-02-16', 'Mưa nhẹ',
    3, 'Một số quả bị thối do độ ẩm cao', 
    '["img-010"]', 'user-002'
),
(
    'hr-005', 'pc-003', '2024-03-05', '10:30:00', 300.000,
    'A+', 'L', 30000.00, 9000000.00,
    15.0, 750000.00, 150000.00, 'Kho lạnh C3',
    2.0, 'Xuất khẩu Nhật Bản', '2024-03-06', 'Nắng đẹp',
    5, 'Chất lượng xuất khẩu, đạt tiêu chuẩn quốc tế', 
    '["img-011", "img-012", "img-013", "img-014", "img-015"]', 'user-003'
);
