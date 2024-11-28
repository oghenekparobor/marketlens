-- Migration for password_reset table
CREATE TABLE password_reset (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,  -- Foreign key to users
    reset_token VARCHAR(255) UNIQUE NOT NULL,             -- Token sent to the user for password reset
    expires_at TIMESTAMP NOT NULL,                        -- Token expiration time
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
