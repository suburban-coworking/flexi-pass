
-- +migrate Up
CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY,
    fullname TEXT,
    email TEXT NOT NULL UNIQUE
);

-- +migrate Down
DROP TABLE IF EXISTS user;