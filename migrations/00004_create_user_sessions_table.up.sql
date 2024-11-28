-- Migration for user_sessions table
CREATE TABLE user_sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,  -- Foreign key to users
    session_token VARCHAR(255) UNIQUE NOT NULL,           -- Session token for authentication
    user_agent TEXT,                                      -- Store user device info
    ip_address VARCHAR(45),                               -- IP address of the user
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL                         -- Session expiration time
);
