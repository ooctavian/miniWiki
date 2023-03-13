CREATE TABLE category (
  category_id BIGSERIAL NOT NULL PRIMARY KEY,
  title TEXT NOT NULL,
  parent_id BIGINT DEFAULT NULL REFERENCES category(category_id),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

ALTER TABLE resource ADD COLUMN category_id BIGINT REFERENCES category(category_id);
