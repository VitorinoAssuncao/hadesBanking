package transfer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/gateway/web/middleware"
	"stoneBanking/app/gateway/web/transfer/vo/input"
	validations "stoneBanking/app/gateway/web/transfer/vo/input/validations/transfer"
	"stoneBanking/app/gateway/web/transfer/vo/output"
)

func (controller Controller) Create(w http.ResponseWriter, r *http.Request) {

	accountID, err := middleware.GetAccountIDFromToken(r, controller.signingKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var transferData = input.CreateTransferVO{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &transferData)

	transferData.AccountOriginID = accountID

	transferData, err = validations.ValidateTransferData(transferData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	transfer := transferData.GenerateTransfer()
	newTransfer, err := controller.usecase.Create(context.Background(), transfer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transferOutput := output.TransferToTransferOutput(newTransfer)
	json.NewEncoder(w).Encode(transferOutput)
}
