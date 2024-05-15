-- -- +goose Up
-- -- how can i make an update to hives table when a new inspection is added or whne the last inspection is updated ps im using sqlite with turso
-- -- +goose StatementBegin
-- CREATE TRIGGER update_hive_status
-- AFTER INSERT ON inspections
-- BEGIN
--     UPDATE hives
--     SET last_inspection_date = new.inspection_date,
--         last_inspection_health = new.overall_health
--     WHERE hives.id = new.hive_id;
-- END;
-- -- +goose StatementEnd

-- -- +goose Down
-- -- +goose StatementBegin
-- DROP TRIGGER update_hive_status;
-- -- +goose StatementEnd


// TODO : LATEST see how to create a trigger in sqlite