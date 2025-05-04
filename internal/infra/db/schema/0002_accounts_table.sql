-- Create accounts table
CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_number VARCHAR(20) NOT NULL UNIQUE,
    customer_id UUID NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    balance DOUBLE PRECISION NOT NULL DEFAULT 0,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_accounts_account_number 
        UNIQUE (account_number)
);

-- Create indexes for accounts table
    CREATE INDEX IF NOT EXISTS idx_accounts_customer_id ON accounts(customer_id);
    CREATE INDEX IF NOT EXISTS idx_accounts_account_number ON accounts(account_number);
