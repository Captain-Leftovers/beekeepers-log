-- +goose Up
-- +goose StatementBegin
INSERT INTO farms (id, name, created_at, updated_at)
VALUES (
        '1',
        'dummy Farm 1 for User 5',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        '2',
        'dummy Farm 2 for User 5',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        '3',
        'dummy Farm 3 for User 5',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        '4',
        'dummy Farm 1 for User 10',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        '5',
        'dummy Farm 2 for User 10',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        '6',
        'dummy Farm 1 for User 15',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        '7',
        'dummy Farm 2 for User 15',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        '8',
        'dummy Farm 3 for User 15',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        '9',
        'dummy Farm 1 for User 20',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        '10',
        'dummy Farm 2 for User 20',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM farms
WHERE name LIKE 'dummy%';
-- +goose StatementEnd