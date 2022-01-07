package accounts

import (
	"encoding/json"
	"net/http"
	"stoneBanking/app/gateway/web/account/vo/output"

	"github.com/gorilla/mux"
)

func (controller *Controller) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["user_id"]
	balance, err := controller.usecase.GetBalance(r.Context(), accountId)
	balanceOutput := output.AccountBalanceVO{
		Balance: balance,
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(balanceOutput)
}
