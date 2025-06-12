
CREATE SCHEMA IF NOT EXISTS news;

CREATE TABLE IF NOT EXISTS news.sessions (
    id VARCHAR(64) PRIMARY KEY,
    data BYTEA NOT NULL,
    expiry TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_sessions_expiry ON news.sessions (expiry);