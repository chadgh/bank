// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: InsertTransaction.sql

package database

import (
	"context"
	"database/sql"
)

const insertTransaction = `-- name: InsertTransaction :one
INSERT INTO account_transactions 
    (message_id, user_id, amount_cents, currency, transaction_type)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING message_id, user_id, amount_cents, currency, transaction_type, created
`

type InsertTransactionParams struct {
	MessageID       string              `json:"message_id"`
	UserID          string              `json:"user_id"`
	AmountCents     int32               `json:"amount_cents"`
	Currency        sql.NullString      `json:"currency"`
	TransactionType TransactionTypeEnum `json:"transaction_type"`
}

func (q *Queries) InsertTransaction(ctx context.Context, arg InsertTransactionParams) (AccountTransaction, error) {
	row := q.queryRow(ctx, q.insertTransactionStmt, insertTransaction,
		arg.MessageID,
		arg.UserID,
		arg.AmountCents,
		arg.Currency,
		arg.TransactionType,
	)
	var i AccountTransaction
	err := row.Scan(
		&i.MessageID,
		&i.UserID,
		&i.AmountCents,
		&i.Currency,
		&i.TransactionType,
		&i.Created,
	)
	return i, err
}