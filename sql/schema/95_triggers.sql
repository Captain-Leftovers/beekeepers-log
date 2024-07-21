-- +goose Up
CREATE TABLE testing(
    id INT PRIMARY KEY,
    description TEXT NOT NULL DEFAULT 'none'
);
CREATE TRIGGER testTrigger
AFTER
INSERT ON test_table FOR EACH ROW 
BEGIN
   UPDATE testing SET description = 'new description' WHERE id = NEW.id
END;
-- +goose Down
DROP TRIGGER testTrigger;
DROP TABLE testing;