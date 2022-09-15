-- name: InsertTransaction :one
INSERT INTO account_transactions 
    (message_id, user_id, amount_cents, currency, transaction_type)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
;