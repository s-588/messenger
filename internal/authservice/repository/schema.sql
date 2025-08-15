CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users(
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(60) NOT NULL,
    registered_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE refresh_tokens (
    token_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    token_hash TEXT NOT NULL,
    issued_at TIMESTAMP NOT NULL DEFAULT now(),
    expires_at TIMESTAMP NOT NULL,
    device_info TEXT,
    ip_address TEXT
);

CREATE INDEX refresh_tokens_token_hash_idx ON refresh_tokens(token_hash);
CREATE INDEX refresh_tokens_user_id_idx ON refresh_tokens(user_id);

