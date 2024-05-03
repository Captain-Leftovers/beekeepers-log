
-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, username, email, password, role, created_at, updated_at) 
VALUES
('1', 'dummy1', 'dummy1@example.com', 'password1', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('2', 'dummy2', 'dummy2@example.com', 'password2', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('3', 'dummy3', 'dummy3@example.com', 'password3', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('4', 'dummy4', 'dummy4@example.com', 'password4', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('5', 'dummy5', 'dummy5@example.com', 'password5', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('6', 'dummy6', 'dummy6@example.com', 'password6', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('7', 'dummy7', 'dummy7@example.com', 'password7', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('8', 'dummy8', 'dummy8@example.com', 'password8', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('9', 'dummy9', 'dummy9@example.com', 'password9', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('10', 'dummy10', 'dummy10@example.com', 'password10', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('11', 'dummy11', 'dummy11@example.com', 'password11', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('12', 'dummy12', 'dummy12@example.com', 'password12', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('13', 'dummy13', 'dummy13@example.com', 'password13', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('14', 'dummy14', 'dummy14@example.com', 'password14', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('15', 'dummy15', 'dummy15@example.com', 'password15', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('16', 'dummy16', 'dummy16@example.com', 'password16', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('17', 'dummy17', 'dummy17@example.com', 'password17', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('18', 'dummy18', 'dummy18@example.com', 'password18', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('19', 'dummy19', 'dummy19@example.com', 'password19', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('20', 'dummy20', 'dummy20@example.com', 'password20', 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE username LIKE 'dummy%';
-- +goose StatementEnd