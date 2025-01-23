package service

import (
	"time"
	"valsea/src/data"
	"valsea/src/model"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Account struct {
	Repository *data.Repository
}

func NewAccount(repository *data.Repository) *Account {
	return &Account{Repository: repository}
}

func (a *Account) CreateAccounts(accounts []model.Account) ([]model.Account, error) {
	for i := range accounts {
		accounts[i].ID = uuid.New().String()
	}
	accounts, err := a.Repository.CreateAccounts(accounts)
	if err != nil {
		zap.S().Errorf("Error creating accounts: %v", zap.Error(err))
		return nil, err
	}
	return accounts, nil
}

func (a *Account) GetAccountWithTransactionsById(id string) (model.AccountWithTransaction, error) {
	account, err := a.Repository.GetAccountById(id)
	if err != nil {
		return model.AccountWithTransaction{}, err
	}

	transactions, err := a.Repository.GetTransactionsByAccountId(id)
	if err != nil {
		return model.AccountWithTransaction{}, err
	}

	return model.AccountWithTransaction{Account: &account, Transactions: transactions}, nil
}

func (a *Account) GetAccountById(id string) (model.Account, error) {
	account, err := a.Repository.GetAccountById(id)
	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}

func (a *Account) CreateTransaction(t model.Transaction) (model.Transaction, error) {
	t.ID = uuid.New().String()
	t.Timestamp = time.Now()
	t, err := a.Repository.CreateTransaction(t)
	if err != nil {
		return model.Transaction{}, err
	}
	return t, nil
}

func (a *Account) UpdateAccount(account model.Account) error {
	return a.Repository.UpdateAccount(account)
}

func (a *Account) ApplyTransaction(t model.Transaction, account model.Account) error {
	if t.Type == model.TransactionTypeWithdrawal {
		*account.Balance -= t.Amount
	} else {
		*account.Balance += t.Amount
	}

	t.AccountID = account.ID
	if _, err := a.CreateTransaction(t); err != nil {
		return err
	}

	if err := a.UpdateAccount(account); err != nil {
		return err
	}

	return nil
}

func (a *Account) GetTransactionsByAccountId(id string) ([]model.Transaction, error) {
	transactions, err := a.Repository.GetTransactionsByAccountId(id)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
