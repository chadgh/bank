package model

import "time"

/*
 Amount:
	type: object
	additionalProperties: false
	required:
	- amount
	- currency
	- debitOrCredit
	properties:
	amount:
		type: string
		description: The amount in the denomination of the currency. For example, $1 = '1.00'
		minLength: 1
	currency:
		type: string
		minLength: 1
	debitOrCredit:
		$ref: '#/components/schemas/DebitCredit'

 DebitCredit:
	type: string
	description: >-
	Debit or Credit flag for the network transaction. A Debit deducts funds from a user. A credit adds funds to a user.
	enum:
	- DEBIT
	- CREDIT
*/

type DebitCredit string

const (
	DEBIT  DebitCredit = "DEBIT"
	CREDIT DebitCredit = "CREDIT"
)

type Amount struct {
	Amount        string      `json:"amount"`
	Currency      string      `json:"currency"`
	DebitOrCredit DebitCredit `json:"debitOfCredit"`
}

/*
 AuthorizationRequest:
      type: object
      additionalProperties: false
      required:
        - userId
        - messageId
        - transactionAmount
      description: Authorization request that needs to be processed.
      properties:
        userId:
          type: string
          minLength: 1
        messageId:
          type: string
          minLength: 1
        transactionAmount:
          $ref: '#/components/schemas/Amount'

    LoadRequest:
      type: object
      additionalProperties: false
      required:
        - userId
        - messageId
        - transactionAmount
      description: Load request that needs to be processed.
      properties:
        userId:
          type: string
          minLength: 1
        messageId:
          type: string
          minLength: 1
        transactionAmount:
          $ref: '#/components/schemas/Amount'
*/

type AuthorizationRequest struct {
	UserID            string `json:"userId"`
	MessageID         string `json:"messageId"`
	TransactionAmount Amount `json:"transactionAmount"`
}

type LoadRequest struct {
	UserID            string `json:"userId"`
	MessageID         string `json:"messageId"`
	TransactionAmount Amount `json:"transactionAmount"`
}

/*
AuthorizationResponse:
      type: object
      additionalProperties: false
      description: The result of an authorization
      required:
        - userId
        - messageId
        - responseCode
        - balance
      properties:
        userId:
          type: string
          minLength: 1
        messageId:
          type: string
          minLength: 1
        responseCode:

          $ref: '#/components/schemas/ResponseCode'
        balance:
          $ref: '#/components/schemas/Amount'

    LoadResponse:
      type: object
      additionalProperties: false
      description: The result of a load.
      required:
        - userId
        - messageId
        - balance
      properties:
        userId:
          type: string
          minLength: 1
        messageId:
          type: string
          minLength: 1
        balance:
          $ref: '#/components/schemas/Amount'
*/

type AuthorizationResponse struct {
	UserID       string       `json:"userId"`
	MessageID    string       `json:"messageId"`
	ResponseCode ResponseCode `json:"responseCode"`
	Balance      Amount       `json:"balance"`
}

type LoadResponse struct {
	UserID    string `json:"userId"`
	MessageID string `json:"messageId"`
	Balance   Amount `json:"balance"`
}

/*
   ResponseCode:
     type: string
     description: >-
       The response code sent back to the network for the merchant. Multiple declines
       reasons may exist but only one will be sent back to the network. Advice messages
       will include the response code that was sent on our behalf.
     enum:
       - APPROVED
       - DECLINED
*/

type ResponseCode string

const (
	APPROVED ResponseCode = "APPROVED"
	DECLINED ResponseCode = "DECLINED"
)

type Ping struct {
	ServerTime string `json:"serverTime"`
}

func NewPing() Ping {
	return Ping{
		ServerTime: time.Now().String(),
	}
}
