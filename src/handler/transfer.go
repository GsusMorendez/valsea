package handler

import (
	"encoding/json"
	"net/http"
	"valsea/src/handler/validate"
	"valsea/src/model"
	"valsea/src/service"

	"go.uber.org/zap"
)

type Transfer struct {
	Service *service.Transfer
}

func NewTransfer(tService *service.Transfer) *Transfer {
	return &Transfer{Service: tService}
}

func (t *Transfer) Transfer(w http.ResponseWriter, r *http.Request) {
	var transfer model.Transfer
	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		zap.S().Errorf("Error decoding transfer from request:  %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	if err := validate.Transfer(transfer); err != nil {
		zap.S().Errorf("Error validating transfer:  %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	if err := t.Service.Transfer(transfer); err != nil {
		zap.S().Errorf("Error making transfer:  %v", zap.Error(err))
		HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	HandleResponse(w, nil, nil, http.StatusCreated)
}
