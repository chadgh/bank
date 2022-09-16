-- name: InsertCreditTransaction :one
INSERT INTO account_transactions 
    (message_id, user_id, credit_cents, currency)
    VALUES ($1, $2, $3, $4)
    RETURNING *
;

-- name: InsertDebitTransaction :one
INSERT INTO account_transactions 
    (message_id, user_id, debit_cents, currency)
    VALUES ($1, $2, $3, $4)
    RETURNING *
;