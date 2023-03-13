CREATE EXTENSION citext;
CREATE DOMAIN domain_email AS citext
    CHECK(VALUE ~ '^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$');

CREATE TABLE account(
    account_id BIGSERIAL PRIMARY KEY,
    email domain_email UNIQUE,
    password TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX account_lower_email
    ON account(lower(email));

CREATE TABLE profile(
    account_id BIGINT REFERENCES account(account_id),
    name TEXT NOT NULL,
    alias TEXT DEFAULT NULL,
    picture_url TEXT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE session(
    session_id VARCHAR PRIMARY KEY,
    account_id BIGINT REFERENCES account(account_id),
    ip_address VARCHAR(45) DEFAULT NULL,
    user_agent TEXT DEFAULT NULL,
    expire_at TIMESTAMP DEFAULT NOW() + (60 ||' minutes')::interval
);

CREATE TYPE resource_state AS ENUM ('PUBLIC', 'PRIVATE');
ALTER TABLE resource ADD COLUMN state resource_state DEFAULT 'PUBLIC';
ALTER TABLE resource ADD COLUMN author_id BIGINT REFERENCES account(account_id);
ALTER TABLE category ADD COLUMN author_id BIGINT REFERENCES account(account_id);

ALTER TABLE resource DROP CONSTRAINT ulink;
ALTER TABLE category DROP CONSTRAINT utitle;
ALTER TABLE resource ADD CONSTRAINT ulink UNIQUE (link,author_id);
ALTER TABLE category ADD CONSTRAINT utitle UNIQUE (title,author_id);

