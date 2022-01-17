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
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode([]output.OutputError{{Error: err.Error()}})
		return
	}

	json.NewEncoder(w).Encode(transfersOutput)
}
