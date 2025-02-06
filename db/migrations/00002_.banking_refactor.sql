-- +goose Up
ALTER TABLE accounts ADD COLUMN amount FLOAT;
-- +goose Down

ALTER TABLE accounts DROP COLUMN IF EXISTS amount;