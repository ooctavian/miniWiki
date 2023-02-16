CREATE TABLE resource(
    resource_id BIGSERIAL PRIMARY KEY,
    title TEXT,
    description TEXT,
    link TEXT,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW()
)