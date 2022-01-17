package transfer

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/web/middleware"
	"stoneBanking/app/gateway/web/transfer/vo/input"
	validations "stoneBanking/app/gateway/web/transfer/vo/input/validations/transfer"
	"stoneBanking/app/gateway/web/transfer/vo/output"
)

//@Sumary Create a transfer
//@Description With the data received, validate her and if all is correct, create a new transfer, and update the balance of accounts
//@Accept json
//@Produce json
//@Param authorization header string true "Authorization Token"
//@Param transfer body input.CreateTransferVO true "Transfer Creation Data"
//@Success 200 {object} output.TransferOutputVO
//@Failure	400 {object} output.OutputError
//@Failure 500 {object} output.OutputError
//@Router /transfer [POST]
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
		if errors.Is(err, customError.ErrorTransferCreate) {
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
