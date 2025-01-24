package validate

import (
	"errors"
	"testing"
	"valsea/src/model"
)

func TestAccounts(t *testing.T) {
	tests := []struct {
		name     string
		input    []model.Account
		expected error
	}{
		{
			"valid accounts",
			[]model.Account{
				{Owner: "Alice", Balance: floatPtr(100)},
				{Owner: "Bob", Balance: floatPtr(200)},
			},
			nil,
		},
		{
			"missing owner",
			[]model.Account{
				{Owner: "", Balance: floatPtr(100)},
			},
			errors.New("account owner is mandatory"),
		},
		{
			"missing balance",
			[]model.Account{
				{Owner: "Alice", Balance: nil},
			},
			errors.New("account balance is mandatory for owner: Alice "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Accounts(tt.input)
			if (err != nil) != (tt.expected != nil) || (err != nil && err.Error() != tt.expected.Error()) {
				t.Errorf("got %v, want %v", err, tt.expected)
			}
		})
	}
}

func TestTransaction(t *testing.T) {
	tests := []struct {
		name     string
		input    model.Transaction
		expected error
	}{
		{
			"valid withdrawal",
			model.Transaction{Amount: 100, Type: "withdrawal"},
			nil,
		},
		{
			"valid deposit",
			model.Transaction{Amount: 50, Type: "deposit"},
			nil,
		},
		{
			"amount is zero",
			model.Transaction{Amount: 0, Type: "deposit"},
			errors.New("transaction amount is mandatory and cannot be 0"),
		},
		{
			"invalid transaction type",
			model.Transaction{Amount: 100, Type: "transfer"},
			errors.New("invalid transaction type"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Transaction(tt.input)
			if (err != nil) != (tt.expected != nil) || (err != nil && err.Error() != tt.expected.Error()) {
				t.Errorf("got %v, want %v", err, tt.expected)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	tests := []struct {
		name     string
		input    model.Transfer
		expected error
	}{
		{
			"valid transfer",
			model.Transfer{Amount: 100, From: "Alice", To: "Bob"},
			nil,
		},
		{
			"amount is zero",
			model.Transfer{Amount: 0, From: "Alice", To: "Bob"},
			errors.New("amount is mandatory and cannot be 0"),
		},
		{
			"missing from",
			model.Transfer{Amount: 100, From: "", To: "Bob"},
			errors.New("from is mandatory"),
		},
		{
			"missing to",
			model.Transfer{Amount: 100, From: "Alice", To: ""},
			errors.New("to is mandatory"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Transfer(tt.input)
			if (err != nil) != (tt.expected != nil) || (err != nil && err.Error() != tt.expected.Error()) {
				t.Errorf("got %v, want %v", err, tt.expected)
			}
		})
	}
}

func floatPtr(f float64) *float64 {
	return &f
}
