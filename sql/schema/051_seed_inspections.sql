-- +goose Up
-- +goose StatementBegin
INSERT INTO inspections (id, hive_id, inspection_date, honey_stores, pests, diseases, queen_seen, eggs_present, larvae_present, capped_brood, notes, created_at, updated_at, overall_health) 
VALUES
('1', '1', CURRENT_TIMESTAMP, 'low', 'none', 'none', 0, 0, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'healthy'),
('2', '2', CURRENT_TIMESTAMP, 'medium', 'few', 'none', 1, 1, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'weak'),
('3', '3', CURRENT_TIMESTAMP, 'high', 'many', 'few', 0, 0, 1, 1, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'swarming'),
('4', '4', CURRENT_TIMESTAMP, 'low', 'none', 'none', 1, 1, 1, 1, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'healthy'),
('5', '5', CURRENT_TIMESTAMP, 'medium', 'few', 'many', 0, 0, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'weak');
-- +goose StatementEnd