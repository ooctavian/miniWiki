ALTER TABLE account DROP COLUMN name;
ALTER TABLE account DROP COLUMN alias;
ALTER TABLE account DROP COLUMN picture_url;
CREATE TABLE profile(
    account_id BIGINT REFERENCES account(account_id),
    name TEXT NOT NULL,
    alias TEXT DEFAULT NULL,
    picture_url TEXT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
ALTER TABLE resource RENAME COLUMN picture_url TO image;