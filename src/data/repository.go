package data

import "valsea/src/model"

type Repository struct {
	rows *Rows
}

func NewRepository() *Repository {
	return &Repository{
		rows: createInitialData(),
	}
}

func (r *Repository) CreateAccounts(accounts []model.Account) ([]model.Account, error) {
	r.rows.createAccounts(accounts)
	return accounts, nil
}

func (r *Repository) GetAccountById(id string) (model.Account, error) {
	return r.rows.getAccountById(id)
}

func (r *Repository) ListAccounts() ([]model.Account, error) {
	return r.rows.listAccounts(), nil
}

func (r *Repository) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	return r.rows.createTransaction(transaction)
}

func (r *Repository) GetTransactionById(id string) (model.Transaction, error) {
	return r.rows.getTransactionById(id)
}

func (r *Repository) ListTransactions() ([]model.Transaction, error) {
	return r.rows.listTransactions(), nil
}

func (r *Repository) GetTransactionsByAccountId(id string) ([]model.Transaction, error) {
	return r.rows.getTransactionsByAccountId(id)
}

func (r *Repository) UpdateAccount(account model.Account) error {
	return r.rows.updateAccount(account)
}
