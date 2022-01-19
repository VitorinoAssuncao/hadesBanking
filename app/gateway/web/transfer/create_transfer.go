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
	const operation = "Gateway.Rest.Transfer.Create"
	controller.log.LogInfo(operation, "received request in url: "+r.URL.Path)

	controller.log.LogInfo(operation, "getting the account id from the token in the header")
	accountID, err := middleware.GetAccountIDFromToken(r, controller.tokenRepo)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	var transferData = input.CreateTransferVO{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	controller.log.LogInfo(operation, "unmarshalling the data to a input object")
	json.Unmarshal(reqBody, &transferData)

	transferData.AccountOriginID = accountID

	controller.log.LogInfo(operation, "validating the data")
	transferData, err = validations.ValidateTransferData(transferData)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	controller.log.LogInfo(operation, "transforming in a internal object")
	transfer := transferData.GenerateTransfer()
	newTransfer, err := controller.usecase.Create(context.Background(), transfer)
	if err != nil {
		if errors.Is(err, customError.ErrorTransferCreate) {
			controller.log.LogError(operation, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
			return
		}

		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	transferOutput := output.TransferToTransferOutput(newTransfer)
	controller.log.LogInfo(operation, "transfer created sucessfully")
	json.NewEncoder(w).Encode(transferOutput)
}
