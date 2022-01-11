package transfer

import (
	"context"
	"encoding/json"
	"net/http"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/web/middleware"
	"stoneBanking/app/gateway/web/transfer/vo/output"
)

func (controller Controller) GetAllByAccountID(w http.ResponseWriter, r *http.Request) {
	accountID, err := middleware.GetAccountIDFromToken(r, controller.signingKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	transfers, err := controller.usecase.GetAllByAccountID(context.Background(), types.ExternalID(accountID))

	var transfersOutput = make([]output.TransferOutputVO, 0)
	for _, transfer := range transfers {
		transferOutput := output.TransferToTransferOutput(transfer)
		transfersOutput = append(transfersOutput, transferOutput)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(transfersOutput)
}
