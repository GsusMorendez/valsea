package validate

import (
	"errors"
	"fmt"
	"strings"
	"valsea/src/model"
)

func Accounts(accounts []model.Account) error {
	for _, a := range accounts {
		if err := Account(a); err != nil {
			return err
		}
	}
	return nil
}

func Account(account model.Account) error {
	if strings.TrimSpace(account.Owner) == "" {
		return errors.New("account owner is mandatory")
	}
	if account.Balance == nil {
		return fmt.Errorf("account balance is mandatory for owner: %s ", account.Owner)
	}
	return nil
}

func Transaction(t model.Transaction) error {
	if t.Amount == 0 {
		return errors.New("transaction amount is mandatory and cannot be 0")
	}
	if t.Type != model.TransactionTypeWithdrawal && t.Type != model.TransactionTypeDeposit {
		return errors.New("invalid transaction type")
	}
	return nil
}

func Transfer(t model.Transfer) error {
	if t.Amount == 0 {
		return errors.New("amount is mandatory and cannot be 0")
	}
	if t.From == "" {
		return errors.New("from is mandatory")
	}
	if t.To == "" {
		return errors.New("to is mandatory")
	}
	return nil
}
