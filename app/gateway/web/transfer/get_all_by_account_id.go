package transfer

import (
	"context"
	"encoding/json"
	"net/http"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/web/middleware"
	"stoneBanking/app/gateway/web/transfer/vo/output"
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
	accountID, err := middleware.GetAccountIDFromToken(r, controller.tokenRepo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]output.OutputError{{Error: err.Error()}})
		return
	}

	transfers, err := controller.usecase.GetAllByAccountID(context.Background(), types.ExternalID(accountID))

	var transfersOutput = make([]output.TransferOutputVO, 0)
	for _, transfer := range transfers {
		transferOutput := output.TransferToTransferOutput(transfer)
		transfersOutput = append(transfersOutput, transferOutput)
	}

	if err != nil {
		if err != customError.ErrorTransferListing {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode([]output.OutputError{{Error: err.Error()}})
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode([]output.OutputError{{Error: err.Error()}})
		return
	}

	json.NewEncoder(w).Encode(transfersOutput)
}
