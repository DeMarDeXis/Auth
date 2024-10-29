CREATE TABLE IF NOT EXISTS users
(
    id        SERIAL PRIMARY KEY,
    email     varchar(255) NOT NULL UNIQUE,
    pass_hash bytea NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_email ON users (email);

CREATE TABLE IF NOT EXISTS apps
(
    id     INTEGER PRIMARY KEY,
    name   TEXT NOT NULL UNIQUE,
    secret TEXT NOT NULL UNIQUE
);

--ALTER TABLE users
--ADD COLUMN is_admin BOOLEAN NOT NULL DEFAULT FALSE;