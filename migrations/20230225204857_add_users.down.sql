ALTER TABLE resource DROP CONSTRAINT ulink;
ALTER TABLE resource ADD CONSTRAINT ulink UNIQUE (link);
ALTER TABLE resource DROP COLUMN author_id;
ALTER TABLE resource DROP COLUMN state;
DROP TYPE resource_state;
DROP TABLE session;
DROP TABLE profile;
DROP TABLE account;
DROP DOMAIN domain_email;
DROP EXTENSION citext;

