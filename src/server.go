package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"chadgh.com/bank/api/model"
	rmodel "chadgh.com/bank/model"
	"chadgh.com/bank/repository"
	"chadgh.com/bank/scripts"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func runserver(verbose bool) error {
	if verbose {
		fmt.Println("Starting server...")
	}

	databaseUrl := os.Getenv("POSTGRES_URL")
	// db, err := sql.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return err
	}

	r := repository.NewTransactionRepo(db)
	s := server{
		repo:    r,
		verbose: verbose,
	}

	router := gin.Default()
	router.GET("/ping", s.pingRequest)
	router.PUT("/authorization/:messageId", s.authorizationRequest)
	router.PUT("/load/:messageId", s.loadRequest)

	s.router = router

	s.Run()

	return nil
}

func NewServer(verbose bool) *server {
	databaseUrl := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil
	}

	r := repository.NewTransactionRepo(db)
	s := server{
		repo:    r,
		verbose: verbose,
	}

	router := gin.Default()
	router.GET("/ping", s.pingRequest)
	router.PUT("/authorization/:messageId", s.authorizationRequest)
	router.PUT("/load/:messageId", s.loadRequest)

	s.router = router
	return &s
}

type server struct {
	repo    repository.TransactionsRepo
	router  *gin.Engine
	verbose bool
}

func (s server) Run() {
	s.router.Run("localhost:8080")
}

func (s server) pingRequest(c *gin.Context) {
	c.JSON(http.StatusOK, model.NewPing())
}

func (s server) authorizationRequest(c *gin.Context) {
	messageId := c.Param("messageId")
	arequest := model.AuthorizationRequest{}
	err := c.BindJSON(&arequest)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if messageId != arequest.MessageID {
		if s.verbose {
			fmt.Println("message ids", messageId, arequest.MessageID)
		}
		c.Error(errors.New("bad request, invalid message id"))
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if s.verbose {
		fmt.Println(
			"inserting:",
			arequest.MessageID,
			arequest.UserID,
			arequest.TransactionAmount.Amount,
			arequest.TransactionAmount.Currency,
			model.CREDIT,
		)
	}

	balance, err := s.repo.GetAccountBalance(c, arequest.UserID)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if !balance.CheckDebit(arequest.TransactionAmount.Amount) {
		c.Error(errors.New("unauthorized"))
		c.AbortWithStatus(http.StatusBadRequest)
	}

	_, err = s.repo.InsertTransaction(
		c,
		arequest.MessageID,
		arequest.UserID,
		arequest.TransactionAmount.Amount,
		arequest.TransactionAmount.Currency,
		rmodel.DEBIT,
	)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if s.verbose {
		fmt.Println("got balance:", balance.Balance)
	}
	currentBalance, _ := scripts.ConvertAmountToCents(balance.Balance)
	amountCent, _ := scripts.ConvertAmountToCents(arequest.TransactionAmount.Amount)
	newBalance := scripts.ConvertCentsToAmount(currentBalance - amountCent)
	c.JSON(http.StatusCreated, model.LoadResponse{
		UserID:    arequest.UserID,
		MessageID: arequest.MessageID,
		Balance:   model.Amount{Amount: newBalance, Currency: arequest.TransactionAmount.Currency},
	})
}

func (s server) loadRequest(c *gin.Context) {
	messageId := c.Param("messageId")
	lrequest := model.LoadRequest{}
	err := c.BindJSON(&lrequest)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if messageId != lrequest.MessageID {
		if s.verbose {
			fmt.Println("message ids", messageId, lrequest.MessageID)
		}
		c.Error(errors.New("bad request, invalid message id"))
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if s.verbose {
		fmt.Println(
			"inserting:",
			lrequest.MessageID,
			lrequest.UserID,
			lrequest.TransactionAmount.Amount,
			lrequest.TransactionAmount.Currency,
			model.CREDIT,
		)
	}

	_, err = s.repo.InsertTransaction(
		c,
		lrequest.MessageID,
		lrequest.UserID,
		lrequest.TransactionAmount.Amount,
		lrequest.TransactionAmount.Currency,
		rmodel.CREDIT,
	)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if s.verbose {
		fmt.Println("getting balance:", lrequest.UserID)
	}
	balance, err := s.repo.GetAccountBalance(c, lrequest.UserID)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if s.verbose {
		fmt.Println("got balance:", balance.Balance)
	}
	c.JSON(http.StatusCreated, model.LoadResponse{
		UserID:    lrequest.UserID,
		MessageID: lrequest.MessageID,
		Balance:   model.Amount{Amount: balance.Balance, Currency: lrequest.TransactionAmount.Currency},
	})
}

/*
  /authorization/{messageId}:
    put:
      summary: Removes funds from a user's account if sufficient funds are available.
      parameters:
        - $ref: '#/components/parameters/messageId'
      requestBody:
        $ref: '#/components/requestBodies/AuthorizationRequest'
      responses:
        201:
          $ref: '#/components/responses/AuthorizationResponse'
        default:
          $ref: '#/components/responses/ServerError'

  /load/{messageId}:
    put:
      summary: Adds funds to a user's account.
      parameters:
        - $ref: '#/components/parameters/messageId'
      requestBody:
        $ref: '#/components/requestBodies/LoadRequest'
      responses:
        201:
          $ref: '#/components/responses/LoadResponse'
        default:
          $ref: '#/components/responses/ServerError'
*/
