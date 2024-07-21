-- +goose Up
-- +goose StatementBegin
INSERT INTO user_farms (user_id, farm_id)
VALUES ('5', '1'),
    ('5', '2'),
    ('5', '3'),
    ('10', '4'),
    ('10', '5'),
    ('15', '6'),
    ('15', '7'),
    ('15', '8'),
    ('20', '9'),
    ('20', '10');
-- +goose StatementEnd