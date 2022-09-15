package model

import "chadgh.com/bank/database"

type Transaction struct {
	MessageID       string
	UserID          string
	Amount          string
	Currency        string
	TransactionType database.TransactionTypeEnum
}
