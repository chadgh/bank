-- name: GetAccountBalance :one
SELECT SUM(DISTINCT credits.C) - SUM(DISTINCT debits.D) as total
FROM (SELECT a1.amount_cents AS C FROM account_transactions a1 WHERE a1.user_id = $1 AND a1.transaction_type = 'CREDIT') as credits,
    (SELECT a2.amount_cents AS D FROM account_transactions a2 WHERE a2.user_id = $1 AND a2.transaction_type = 'DEBIT') as debits
;