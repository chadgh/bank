DROP TYPE IF EXISTS TRANSACTION_TYPE_ENUM;
CREATE TYPE transaction_type_enum AS ENUM ('CREDIT', 'DEBIT');

CREATE TABLE account_transactions (
    message_id VARCHAR(50) PRIMARY KEY UNIQUE NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    amount_cents INT NOT NULL,
    currency VARCHAR(10),
    transaction_type TRANSACTION_TYPE_ENUM NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_transaction_user_id ON account_transactions(user_id);