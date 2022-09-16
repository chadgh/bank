package model

type CreditOrDebit string

const (
	CREDIT CreditOrDebit = "CREDIT"
	DEBIT  CreditOrDebit = "DEBIT"
)

type Transaction struct {
	MessageID       string
	UserID          string
	Amount          string
	Currency        string
	TransactionType CreditOrDebit
}
