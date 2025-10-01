-- Remove seed data for pest_disease_records table
DELETE FROM pest_disease_records WHERE id IN (
    'pdr-001', 'pdr-002', 'pdr-003', 'pdr-004', 'pdr-005'
);
