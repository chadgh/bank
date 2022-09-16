package model

import "chadgh.com/bank/scripts"

type Account struct {
	UserID  string
	Balance string
	// Currency string
}

func (a Account) CheckDebit(amount string) bool {
	balance, _ := scripts.ConvertAmountToCents(a.Balance)
	debit, _ := scripts.ConvertAmountToCents(amount)
	if (balance - debit) >= 0 {
		return true
	}
	return false
}
