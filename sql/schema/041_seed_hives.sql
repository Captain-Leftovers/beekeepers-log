-- +goose Up
-- +goose StatementBegin
INSERT INTO hives (id, farm_id, name, isEmpty, notes, created_at, updated_at) 
VALUES
('1', '1', 'Hive 1 for Farm 1', 1, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('2', '1', 'Hive 2 for Farm 1', 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('3', '2', 'Hive 1 for Farm 2', 1, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('4', '2', 'Hive 2 for Farm 2', 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('5', '2', 'Hive 3 for Farm 2', 1, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('6', '3', 'Hive 1 for Farm 3', 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('7', '3', 'Hive 2 for Farm 3', 1, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('8', '3', 'Hive 3 for Farm 3', 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('9', '3', 'Hive 4 for Farm 3', 1, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('10', '4', 'Hive 1 for Farm 4', 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- +goose StatementEnd