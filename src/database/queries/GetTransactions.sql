-- name: GetTransactions :many
SELECT * FROM account_transactions WHERE user_id = $1 ORDER BY created;