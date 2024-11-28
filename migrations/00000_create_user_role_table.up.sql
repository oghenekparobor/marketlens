-- Migration for user_roles table
-- Step 0: Ensure the uuid-ossp extension is enabled (Optional)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);