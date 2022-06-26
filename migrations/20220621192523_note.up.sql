CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS note(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    message VARCHAR(255),
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp
);