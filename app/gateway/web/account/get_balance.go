package accounts

import (
	"encoding/json"
	"net/http"
	"stoneBanking/app/common/utils"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/web/account/vo/output"
)

func (controller *Controller) GetBalance(w http.ResponseWriter, r *http.Request) {
	headerToken := r.Header.Get("Authorization")
	if headerToken == "" {
		http.Error(w, customError.ErrorServerTokenNotFound.Error(), http.StatusBadRequest)
		return
	}

	tokenID, err := utils.ExtractClaims(headerToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
