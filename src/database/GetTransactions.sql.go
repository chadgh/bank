// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: GetTransactions.sql

package database

import (
	"context"
)

const getTransactions = `-- name: GetTransactions :many
SELECT message_id, user_id, amount_cents, currency, transaction_type, created FROM account_transactions WHERE user_id = $1 ORDER BY created
`

func (q *Queries) GetTransactions(ctx context.Context, userID string) ([]AccountTransaction, error) {
	rows, err := q.query(ctx, q.getTransactionsStmt, getTransactions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AccountTransaction
	for rows.Next() {
		var i AccountTransaction
		if err := rows.Scan(
			&i.MessageID,
			&i.UserID,
			&i.AmountCents,
			&i.Currency,
			&i.TransactionType,
			&i.Created,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}