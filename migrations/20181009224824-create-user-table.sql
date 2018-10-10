
-- +migrate Up
CREATE TABLE IF NOT EXISTS user (
    id PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    fullname TEXT
);

-- +migrate Down
DROP TABLE IF EXISTS user;