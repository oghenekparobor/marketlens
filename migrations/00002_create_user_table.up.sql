-- Migration for users table
-- Migration for users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone_number VARCHAR(15) UNIQUE,
    -- phone_number is nullable by default
    password_hash VARCHAR(255) NOT NULL,
    role_id UUID REFERENCES user_roles(id),
    -- foreign key reference to user_roles
    is_email_verified BOOLEAN DEFAULT FALSE,
    is_phone_number_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);