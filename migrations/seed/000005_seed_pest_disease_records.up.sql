-- Seed data for pest_disease_records table
INSERT INTO pest_disease_records (
    id, planting_cycle_id, type, name, scientific_name, severity,
    affected_area_percentage, affected_plant_count, detection_date, detection_method,
    symptoms, treatment_applied, treatment_date, treatment_cost, treatment_duration_days,
    recovery_status, effectiveness_rating, follow_up_date, prevention_measures,
    environmental_factors, images, notes, created_by
) VALUES 
(
    'pdr-001', 'pc-001', 'pest', 'Rệp sáp', 'Planococcus citri', 'medium',
    15.50, 25, '2024-01-10', 'visual',
    'Xuất hiện lớp sáp trắng trên lá và thân, lá vàng và rụng sớm',
    'Phun thuốc trừ sâu sinh học, cắt tỉa cành bị nhiễm', '2024-01-11', 150000.00, 7,
    'recovered', 4, '2024-01-18', 'Vệ sinh vườn thường xuyên, kiểm tra định kỳ',
    'Độ ẩm cao, nhiệt độ 25-30°C', '["img-pest-001", "img-pest-002"]', 
    'Phát hiện sớm và xử lý kịp thời', 'user-001'
),
(
    'pdr-002', 'pc-001', 'disease', 'Bệnh thối rễ', 'Phytophthora spp.', 'high',
    30.00, 45, '2024-01-25', 'lab_test',
    'Rễ bị thối đen, cây héo và chết, lá vàng từ gốc lên',
    'Tưới thuốc trừ nấm, cải thiện thoát nước', '2024-01-26', 300000.00, 14,
    'recovering', 3, '2024-02-09', 'Cải thiện hệ thống thoát nước, tránh tưới quá nhiều',
    'Đất ẩm ướt kéo dài, pH thấp', '["img-disease-001", "img-disease-002", "img-disease-003"]', 
    'Bệnh nghiêm trọng, cần theo dõi chặt chẽ', 'user-001'
),
(
    'pdr-003', 'pc-002', 'nutrient_deficiency', 'Thiếu Kali', 'Potassium deficiency', 'low',
    8.75, 12, '2024-02-05', 'visual',
    'Lá vàng ở mép, cháy đầu lá, quả nhỏ và chua',
    'Bón phân kali, tưới nước đều đặn', '2024-02-06', 80000.00, 10,
    'recovered', 5, '2024-02-16', 'Bón phân cân đối, kiểm tra pH đất định kỳ',
    'Đất cát, thoát nước nhanh', '["img-nutrient-001"]', 
    'Thiếu hụt dinh dưỡng nhẹ, dễ khắc phục', 'user-002'
),
(
    'pdr-004', 'pc-002', 'pest', 'Sâu đục quả', 'Helicoverpa armigera', 'critical',
    45.00, 80, '2024-02-20', 'trap',
    'Lỗ đục trên quả, sâu non màu xanh lá, quả bị hỏng',
    'Phun thuốc trừ sâu hóa học, thu gom quả bị hại', '2024-02-21', 500000.00, 5,
    'treating', 2, '2024-02-26', 'Đặt bẫy pheromone, vệ sinh vườn sau thu hoạch',
    'Nhiệt độ cao, độ ẩm thấp', '["img-pest-003", "img-pest-004", "img-pest-005"]', 
    'Dịch hại nghiêm trọng, ảnh hưởng lớn đến năng suất', 'user-002'
),
(
    'pdr-005', 'pc-003', 'environmental_stress', 'Stress nhiệt', 'Heat stress', 'medium',
    25.00, 35, '2024-03-01', 'sensor',
    'Lá héo, quả bị nứt, cây phát triển chậm',
    'Tăng cường tưới nước, che nắng, phun sương', '2024-03-02', 200000.00, 3,
    'recovering', 4, '2024-03-05', 'Lắp hệ thống tưới phun sương, che nắng',
    'Nhiệt độ cao 35-40°C, độ ẩm thấp 30%', '["img-stress-001", "img-stress-002"]', 
    'Thời tiết khắc nghiệt, cần biện pháp bảo vệ dài hạn', 'user-003'
);
