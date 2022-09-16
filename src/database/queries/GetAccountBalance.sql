-- name: GetAccountBalance :one
SELECT SUM(credit_cents) - SUM(debit_cents) AS total
FROM account_transactions
WHERE user_id = $1;