package repository

import (
	"context"
	"database/sql"

	"chadgh.com/bank/database"
	"chadgh.com/bank/model"
	"chadgh.com/bank/scripts"
)

type TransactionsRepo interface {
	InsertTransaction(
		ctx context.Context,
		messageID string,
		userID string,
		amount string,
		currency string,
		transactionType model.CreditOrDebit,
	) (*model.Account, error)
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
	transactionType model.CreditOrDebit,
) (*model.Account, error) {
	queries, err := database.Prepare(ctx, t.db)
	if err != nil {
		return nil, err
	}
	amount_cents, err := scripts.ConvertAmountToCents(amount)
	if err != nil {
		return nil, err
	}

	if transactionType == model.CREDIT {
		_, err := queries.InsertCreditTransaction(ctx, database.InsertCreditTransactionParams{
			MessageID:   messageID,
			UserID:      userID,
			CreditCents: amount_cents,
			Currency:    sql.NullString{String: currency, Valid: true},
		})
		if err != nil {
			return nil, err
		}
	} else {
		_, err := queries.InsertDebitTransaction(ctx, database.InsertDebitTransactionParams{
			MessageID:  messageID,
			UserID:     userID,
			DebitCents: amount_cents,
			Currency:   sql.NullString{String: currency, Valid: true},
		})
		if err != nil {
			return nil, err
		}
	}

	balance, err := t.GetAccountBalance(ctx, userID)
	if err != nil {
		return nil, err
	}

	return balance, nil
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
		Balance: scripts.ConvertCentsToAmount(int32(balance)),
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
		transactionType := model.CREDIT
		amount := transaction.CreditCents
		if amount == 0 {
			amount = transaction.DebitCents
			transactionType = model.DEBIT
		}
		transactions = append(transactions, &model.Transaction{
			MessageID:       transaction.MessageID,
			UserID:          transaction.UserID,
			Amount:          scripts.ConvertCentsToAmount(amount),
			Currency:        transaction.Currency.String,
			TransactionType: transactionType,
		})
	}

	return transactions, nil
}
