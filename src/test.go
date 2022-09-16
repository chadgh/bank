package main

import (
	"context"
	"database/sql"
	"fmt"

	"chadgh.com/bank/model"
	"chadgh.com/bank/repository"
	"chadgh.com/bank/scripts"
	_ "github.com/lib/pq"
)

func runtest(verbose bool) error {
	ctx := context.Background()

	db, err := sql.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		return err
	}
	scripts.TruncateAll(db)
	repo := repository.NewTransactionRepo(db)

	_, err = repo.InsertTransaction(ctx, "1", "1", "100.00", "USD", model.CREDIT)
	if err != nil {
		return err
	}
	_, err = repo.InsertTransaction(ctx, "2", "1", "500.00", "USD", model.CREDIT)
	if err != nil {
		return err
	}
	_, err = repo.InsertTransaction(ctx, "3", "1", "1.01", "USD", model.DEBIT)
	if err != nil {
		return err
	}

	_, err = repo.InsertTransaction(ctx, "4", "2", "100.00", "USD", model.CREDIT)
	if err != nil {
		return err
	}
	_, err = repo.InsertTransaction(ctx, "5", "2", "3.00", "USD", model.DEBIT)
	if err != nil {
		return err
	}

	accountBalance, err := repo.GetAccountBalance(ctx, "1")
	if err != nil {
		return err
	}

	OK := "Error"
	if accountBalance.Balance == "598.99" {
		OK = "OK"
	}
	fmt.Println("expected 598.99 ", OK)
	if OK != "OK" {
		fmt.Println("account balance: ", accountBalance.Balance)
	}
	if verbose {
		fmt.Println("transactions")
		transactions, err := repo.GetTransactions(ctx, "1")
		if err != nil {
			return err
		}

		for _, tr := range transactions {
			fmt.Printf("\t message:%s amount:%s type:%s\n", tr.MessageID, tr.Amount, tr.TransactionType)
		}

	}
	scripts.TruncateAll(db)
	return nil
}
