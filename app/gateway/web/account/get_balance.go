package accounts

import (
	"encoding/json"
	"net/http"
	"stoneBanking/app/gateway/web/account/vo/output"
	"stoneBanking/app/gateway/web/middleware"
)

func (controller *Controller) GetBalance(w http.ResponseWriter, r *http.Request) {
	tokenID, err := middleware.GetAccountIDFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	balance, err := controller.usecase.GetBalance(r.Context(), tokenID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	balanceOutput := output.AccountBalanceVO{
		Balance: balance,
	}

	json.NewEncoder(w).Encode(balanceOutput)
}
