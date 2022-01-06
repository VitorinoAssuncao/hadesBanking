package transfer

import (
	"context"
	"encoding/json"
	"net/http"
	"stoneBanking/app/domain/types"

	"github.com/gorilla/mux"
)

func (controller Controller) GetAllByAccountID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	transfersOutput, err := controller.usecase.GetAllByAccountID(context.Background(), types.ExternalID(accountId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(transfersOutput)
}
