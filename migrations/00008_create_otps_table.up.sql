-- Migration for otps table
CREATE TABLE otps (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,  -- Foreign key to users
    otp_code VARCHAR(6) NOT NULL,                         -- OTP code (could be a 6-digit code)
    otp_type VARCHAR(50) NOT NULL,                        -- Type of OTP (e.g., phone_verification, email_verification)
    expires_at TIMESTAMP NOT NULL,                        -- OTP expiration time
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
