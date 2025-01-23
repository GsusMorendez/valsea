package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"valsea/src/handler/validate"
	"valsea/src/model"
	"valsea/src/service"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Account struct {
	Service *service.Account
}

func NewAccount(accountService *service.Account) *Account {
	return &Account{Service: accountService}
}

func (a *Account) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var accountsInput []model.Account
	if err := json.NewDecoder(r.Body).Decode(&accountsInput); err != nil {
		zap.S().Errorf("Error decoding account from request:  %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	if err := validate.Accounts(accountsInput); err != nil {
		zap.S().Errorf("Error validating account:  %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	accounts, err := a.Service.CreateAccounts(accountsInput)
	if err != nil {
		zap.S().Errorf("Error creating account. %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	bodyResponse, err := json.Marshal(accounts)
	if err != nil {
		zap.S().Errorf("Error marshalling response in create accounts. %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	HandleResponse(w, bodyResponse, nil, http.StatusCreated)
}

func (a *Account) GetAccountById(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "id")
	account, err := a.Service.GetAccountWithTransactionsById(accountId)
	if err != nil {
		zap.S().Errorf("Error getting account by id: %v. %v", accountId, zap.Error(err))
		HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	bodyResponse, err := json.Marshal(account)
	if err != nil {
		zap.S().Errorf("Error marshalling response getting account by id: %v. %v", accountId, zap.Error(err))
		HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	HandleResponse(w, bodyResponse, nil, http.StatusCreated)
}

func (a *Account) ListAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := a.Service.Repository.ListAccounts()
	if err != nil {
		zap.S().Errorf("Error listing accounts. %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	bodyResponse, err := json.Marshal(accounts)
	if err != nil {
		zap.S().Errorf("Error marshalling response in list accounts. %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	HandleResponse(w, bodyResponse, nil, http.StatusCreated)
}

func (a *Account) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "id")

	account, err := a.Service.GetAccountById(accountId)
	if err != nil {
		zap.S().Errorf("Error getting account by id: %v.  %v", accountId, zap.Error(err))
		HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	var transaction model.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		zap.S().Errorf("Error decoding account from request:  %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	if err := validate.Transaction(transaction); err != nil {
		zap.S().Errorf("Error validating transaction:  %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	if err := a.Service.ApplyTransaction(transaction, account); err != nil {
		zap.S().Errorf("Error applying transaction for account id.  %v, %v", accountId, zap.Error(err))
		HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	HandleResponse(w, nil, nil, http.StatusCreated)
}

func (a *Account) GetTransactionsByAccountId(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "id")

	if _, err := a.Service.GetAccountById(accountId); err != nil {
		zap.S().Errorf("Error getting account by id: %v.  %v", accountId, zap.Error(err))
		HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	transactions, err := a.Service.GetTransactionsByAccountId(accountId)
	if err != nil {
		zap.S().Errorf("Error getting transactions by account id: %v. %v", accountId, zap.Error(err))
		HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	fmt.Println("hereeeee", transactions)

	bodyResponse, err := json.Marshal(transactions)
	if err != nil {
		zap.S().Errorf("Error marshalling response getting account by id: %v. %v", accountId, zap.Error(err))
		HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	HandleResponse(w, bodyResponse, nil, http.StatusOK)
}

func HandleResponse(w http.ResponseWriter, bodyResponse []byte, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "default-src 'self'")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.WriteHeader(status)

	if err != nil {
		response := model.ErrorResponse{Err: err.Error()}
		bodyResponse, err = json.Marshal(response)
		if err != nil {
			zap.S().Errorf("Error marshalling error response. %v", zap.Error(err))
			return
		}
	}

	if bodyResponse != nil {
		if _, err := w.Write(bodyResponse); err != nil {
			zap.S().Errorf("Error writing response: %v", zap.Error(err))
		}
	}
}
