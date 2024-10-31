CREATE TABLE IF NOT EXISTS users
(
    id        SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    hashToken varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);


--ALTER TABLE users
--ADD COLUMN is_admin BOOLEAN NOT NULL DEFAULT FALSE;