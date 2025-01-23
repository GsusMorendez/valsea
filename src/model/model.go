package model

import "time"

type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "deposit"
	TransactionTypeWithdrawal TransactionType = "withdrawal"
)

type Account struct {
	ID      string   `json:"id,omitempty"`
	Owner   string   `json:"owner"`
	Balance *float64 `json:"initial_balance"`
}

type Transaction struct {
	ID        string          `json:"id,omitempty"`
	AccountID string          `json:"account_id"`
	Type      TransactionType `json:"type"`
	Amount    float64         `json:"amount"`
	Timestamp time.Time       `json:"timestamp"`
}

type AccountWithTransaction struct {
	Account      *Account      `json:"account"`
	Transactions []Transaction `json:"transactions"`
}

type Transfer struct {
	From   string  `json:"from_account_id"`
	To     string  `json:"to_account_id"`
	Amount float64 `json:"amount"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}
