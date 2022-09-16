CREATE TABLE account_transactions (
    message_id VARCHAR(50) PRIMARY KEY UNIQUE NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    credit_cents INT DEFAULT 0 NOT NULL,
    debit_cents INT DEFAULT 0 NOT NULL,
    currency VARCHAR(10),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_transaction_user_id ON account_transactions(user_id);