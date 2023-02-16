CREATE TABLE category (
  category_id BIGSERIAL NOT NULL PRIMARY KEY,
  title text NOT NULL,
  parent_id BIGINT DEFAULT NULL REFERENCES category(category_id),
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp DEFAULT NOW()
);

ALTER TABLE resource ADD COLUMN category_id BIGINT REFERENCES category(category_id);
