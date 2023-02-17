CREATE TABLE IF NOT EXISTS companies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(15) UNIQUE NOT NULL,
    description VARCHAR(3000),
    employee_count INT NOT NULL,
    registered BOOLEAN NOT NULL,
    type VARCHAR(19) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
