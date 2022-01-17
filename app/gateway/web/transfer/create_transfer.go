package transfer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/web/middleware"
	"stoneBanking/app/gateway/web/transfer/vo/input"
	validations "stoneBanking/app/gateway/web/transfer/vo/input/validations/transfer"
	"stoneBanking/app/gateway/web/transfer/vo/output"
)

func (controller Controller) Create(w http.ResponseWriter, r *http.Request) {

	accountID, err := middleware.GetAccountIDFromToken(r, controller.tokenRepo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	var transferData = input.CreateTransferVO{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	json.Unmarshal(reqBody, &transferData)

	transferData.AccountOriginID = accountID

	transferData, err = validations.ValidateTransferData(transferData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	transfer := transferData.GenerateTransfer()
	newTransfer, err := controller.usecase.Create(context.Background(), transfer)
	if err != nil {
		if err != customError.ErrorTransferCreate {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	transferOutput := output.TransferToTransferOutput(newTransfer)
	json.NewEncoder(w).Encode(transferOutput)
}
