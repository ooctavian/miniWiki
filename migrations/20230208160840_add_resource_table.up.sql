CREATE TABLE resource(
    resource_id BIGSERIAL PRIMARY KEY,
    title TEXT,
    description TEXT,
    link TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)