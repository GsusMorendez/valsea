package data

import (
	"errors"
	"fmt"
	"sync"
	"valsea/src/model"
)

type Rows struct {
	mu           sync.RWMutex
	Accounts     []model.Account
	Transactions []model.Transaction
}

func createInitialData() *Rows {
	return &Rows{
		Accounts:     []model.Account{},
		Transactions: []model.Transaction{},
	}
}
func (r *Rows) createAccounts(accounts []model.Account) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Accounts = append(r.Accounts, accounts...)
}

func (r *Rows) getAccountById(id string) (model.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, account := range r.Accounts {
		if account.ID == id {
			return account, nil
		}
	}
	return model.Account{}, fmt.Errorf("account not found by id: %s", id)
}

func (r *Rows) listAccounts() []model.Account {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.Accounts
}

func (r *Rows) createTransaction(transaction model.Transaction) (model.Transaction, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Transactions = append(r.Transactions, transaction)
	return transaction, nil
}

func (r *Rows) getTransactionById(id string) (model.Transaction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, transaction := range r.Transactions {
		if transaction.ID == id {
			return transaction, nil
		}
	}
	return model.Transaction{}, errors.New("transaction not found")
}

func (r *Rows) listTransactions() []model.Transaction {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.Transactions
}

func (r *Rows) getTransactionsByAccountId(accountId string) ([]model.Transaction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var transactions []model.Transaction
	for _, transaction := range r.Transactions {
		if transaction.AccountID == accountId {
			transactions = append(transactions, transaction)
		}
	}
	return transactions, nil
}

func (r *Rows) updateAccount(account model.Account) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i := range r.Accounts {
		if r.Accounts[i].ID == account.ID {
			r.Accounts[i].Balance = account.Balance
			return nil
		}
	}

	return fmt.Errorf("account not found by id %s", account.ID)
}
