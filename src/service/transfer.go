package service

import (
	"errors"
	"fmt"
	"valsea/src/data"
	"valsea/src/model"

	"go.uber.org/zap"
)

type Transfer struct {
	Repository *data.Repository
}

func NewTransfer(repository *data.Repository) *Transfer {
	return &Transfer{Repository: repository}
}

func (t *Transfer) Transfer(transfer model.Transfer) error {
	accountFrom, err := t.Repository.GetAccountById(transfer.From)
	if err != nil {
		return err
	}

	accountTo, err := t.Repository.GetAccountById(transfer.To)
	if err != nil {
		return err
	}

	if *accountFrom.Balance-transfer.Amount < 0 {
		return fmt.Errorf("insufficient funds in accountId: %v", accountFrom.ID)
	}

	*accountTo.Balance += transfer.Amount
	*accountFrom.Balance -= transfer.Amount

	errorMsg := ""
	if err := t.Repository.UpdateAccount(accountFrom); err != nil {
		zap.S().Errorf("Error updating accountId: %v. %v", accountFrom.ID, zap.Error(err))
		errorMsg = err.Error()
	}

	if err := t.Repository.UpdateAccount(accountTo); err != nil {
		zap.S().Errorf("Error updating account: %v. %v", accountFrom.ID, zap.Error(err))
		errorMsg = fmt.Sprintf("%v. %v", errorMsg, err.Error())
	}

	if errorMsg != "" {
		return errors.New(errorMsg)
	}

	return nil
}
