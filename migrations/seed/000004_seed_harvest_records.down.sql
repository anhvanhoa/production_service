-- Remove seed data for harvest_records table
DELETE FROM harvest_records WHERE id IN (
    'hr-001', 'hr-002', 'hr-003', 'hr-004', 'hr-005'
);
