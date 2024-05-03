-- +goose Up
-- +goose StatementBegin
INSERT INTO farms (id, name, created_at, updated_at) 
VALUES
('11', 'dummy Farm 1 for lalilulelo', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('12', 'dummy Farm 2 for lalilulelo', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO user_farms (user_id, farm_id) 
VALUES
('8ca1b97c-79d9-4c66-8f6b-634ae2cf7645', '11'),
('8ca1b97c-79d9-4c66-8f6b-634ae2cf7645', '12');

INSERT INTO hives (id, farm_id, name, isEmpty, notes, created_at, updated_at) 
VALUES
('11', '11', 'Hive 1 for Farm 1', 1, 'my hive 1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('12', '11', 'Hive 2 for Farm 1', 0, 'my hive 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('13', '12', 'Hive 1 for Farm 2', 1, 'my hive text 3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('14', '12', 'Hive 2 for Farm 2', 0, 'hive 3 text here', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO inspections (id, hive_id, inspection_date, honey_stores, pests, diseases, queen_seen, eggs_present, larvae_present, capped_brood, notes, created_at, updated_at, overall_health) 
VALUES
('6', '11', CURRENT_TIMESTAMP, 'low', 'none', 'none', 1, 0, 1, 1, 'see the hive next time its good', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'healthy'),
('7', '12', CURRENT_TIMESTAMP, 'medium', 'few', 'none', 0, 1, 0, 0, 'see this hive next time', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'weak');
-- +goose StatementEnd