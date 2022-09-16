// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: GetAccountBalance.sql

package database

import (
	"context"
)

const getAccountBalance = `-- name: GetAccountBalance :one
SELECT SUM(credit_cents) - SUM(debit_cents) AS total
FROM account_transactions
WHERE user_id = $1
`

func (q *Queries) GetAccountBalance(ctx context.Context, userID string) (int32, error) {
	row := q.queryRow(ctx, q.getAccountBalanceStmt, getAccountBalance, userID)
	var total int32
	err := row.Scan(&total)
	return total, err
}
