-- Migration for addresses table
CREATE TABLE addresses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,  -- Foreign key to users
    address_line_1 VARCHAR(255) NOT NULL,                -- Primary address line
    address_line_2 VARCHAR(255),                         -- Optional secondary address line
    city VARCHAR(100) NOT NULL,                          -- City of the address
    state VARCHAR(100) NOT NULL,                         -- State or region of the address
    postal_code VARCHAR(20),                             -- Postal or ZIP code
    country VARCHAR(100) NOT NULL,                       -- Country of the address
    phone_number VARCHAR(15),                            -- Contact number for the address
    is_default BOOLEAN DEFAULT FALSE,                    -- Mark this address as the default for the user
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
