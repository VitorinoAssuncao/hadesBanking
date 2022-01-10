package transfer

import (
	"context"
	"encoding/json"
	"net/http"
	"stoneBanking/app/common/utils"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/web/transfer/vo/output"
)

func (controller Controller) GetAllByAccountID(w http.ResponseWriter, r *http.Request) {
	headerToken := r.Header.Get("Token")
	if headerToken == "" {
		http.Error(w, customError.ErrorServerTokenNotFound.Error(), http.StatusBadRequest)
		return
	}

	tokenID, err := utils.ExtractClaims(headerToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transfers, err := controller.usecase.GetAllByAccountID(context.Background(), types.ExternalID(tokenID))

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
