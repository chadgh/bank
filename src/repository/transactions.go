package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"chadgh.com/bank/database"
	"chadgh.com/bank/model"
)

type TransactionsRepo interface {
	InsertTransaction(
		ctx context.Context,
		messageID string,
		userID string,
		amount string,
		currency string,
		transactionType database.TransactionTypeEnum,
	) (*model.Transaction, error)
	GetAccountBalance(
		ctx context.Context,
		userID string,
	) (*model.Account, error)
	GetTransactions(
		ctx context.Context,
		userID string,
	) ([]*model.Transaction, error)
}

type transactionsImp struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) TransactionsRepo {
	return transactionsImp{
		db: db,
	}
}

func (t transactionsImp) InsertTransaction(
	ctx context.Context,
	messageID string,
	userID string,
	amount string,
	currency string,
	transactionType database.TransactionTypeEnum,
) (*model.Transaction, error) {
	queries, err := database.Prepare(ctx, t.db)
	if err != nil {
		return nil, err
	}
	amount_cents, err := convertAmountToCents(amount)
	if err != nil {
		return nil, err
	}
	transaction, err := queries.InsertTransaction(ctx, database.InsertTransactionParams{
		MessageID:       messageID,
		UserID:          userID,
		AmountCents:     amount_cents,
		Currency:        sql.NullString{String: currency, Valid: true},
		TransactionType: transactionType,
	})
	if err != nil {
		return nil, err
	}
	return &model.Transaction{
		MessageID:       transaction.MessageID,
		UserID:          transaction.UserID,
		Amount:          convertCentsToAmount(transaction.AmountCents),
		Currency:        transaction.Currency.String,
		TransactionType: transaction.TransactionType,
	}, nil
}

func (t transactionsImp) GetAccountBalance(
	ctx context.Context,
	userID string,
) (*model.Account, error) {
	queries, err := database.Prepare(ctx, t.db)
	if err != nil {
		return nil, err
	}
	balance, err := queries.GetAccountBalance(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &model.Account{
		UserID:  userID,
		Balance: convertCentsToAmount(int32(balance)),
	}, nil
}

func (t transactionsImp) GetTransactions(
	ctx context.Context,
	userID string,
) ([]*model.Transaction, error) {
	queries, err := database.Prepare(ctx, t.db)
	if err != nil {
		return nil, err
	}
	dbTransactions, err := queries.GetTransactions(ctx, userID)
	if err != nil {
		return nil, err
	}

	var transactions []*model.Transaction
	for _, transaction := range dbTransactions {
		transactions = append(transactions, &model.Transaction{
			MessageID:       transaction.MessageID,
			UserID:          transaction.UserID,
			Amount:          convertCentsToAmount(transaction.AmountCents),
			Currency:        transaction.Currency.String,
			TransactionType: transaction.TransactionType,
		})
	}

	return transactions, nil
}

func convertAmountToCents(amount string) (int32, error) {
	parts := strings.Split(amount, ".")
	dollars, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	cents, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	total := int32(cents + (dollars * 100))
	return total, nil
}

func convertCentsToAmount(cents int32) string {
	dollars := cents / 100
	remainder := cents % 100
	return fmt.Sprintf("%d.%.2d", dollars, remainder)
}
