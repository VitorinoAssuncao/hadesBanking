package transfer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/http/middleware"
	"stoneBanking/app/gateway/http/transfer/vo/output"
)

//@Sumary Get all transfers
//@Description With a valid Authorization Token, get all the transfers that has made or received by the account
//@Accept json
//@Produce json
//@Param authorization header string true "Authorization Token"
//@Success 200 {object} []output.TransferOutputVO
//@Failure	400 {object} output.OutputError
//@Failure 500 {object} output.OutputError
//@Router /transfer [GET]
func (controller Controller) GetAllByAccountID(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Transfer.GetAllByAccountID"

	accountID, err := middleware.GetAccountIDFromToken(r, controller.token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]output.OutputError{{Error: err.Error()}}) //nolint: errorlint
		return
	}
	transfers, err := controller.usecase.GetAllByAccountID(context.Background(), types.ExternalID(accountID))

	controller.log.LogInfo(operation, "creating the objects to by listed")
	var transfersOutput = make([]output.TransferOutputVO, 0)
	for _, transfer := range transfers {
		transferOutput := output.TransferToTransferOutput(transfer)
		transfersOutput = append(transfersOutput, transferOutput)
	}

	if err != nil {
		if errors.Is(err, customError.ErrorTransferAccountNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode([]output.OutputError{{Error: err.Error()}}) //nolint: errorlint
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode([]output.OutputError{{Error: err.Error()}}) //nolint: errorlint
		return
	}

	controller.log.LogInfo(operation, "transfers listed sucessfully")
	json.NewEncoder(w).Encode(transfersOutput) //nolint: errorlint
}
