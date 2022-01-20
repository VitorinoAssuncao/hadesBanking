package account

import (
	"encoding/json"
	"errors"
	"net/http"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/web/account/vo/output"
	"stoneBanking/app/gateway/web/middleware"
)

//@Sumary Get the balance of a account
//@Description With a authorization token valid, return the balance of a account
//@Produce json
//@Param authorization header string true "Authorization Token"
//@Success 200 {object} output.AccountBalanceVO
//@Failure	400 {object} output.OutputError
//@Failure 500 {object} output.OutputError
//@Router /account/balance [GET]
func (controller *Controller) GetBalance(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Account.GetBalance"

	controller.log.LogInfo(operation, "take the value from the token")
	tokenID, err := middleware.GetAccountIDFromToken(r, controller.tokenRepo)
	if err != nil {
		if errors.Is(err, customError.ErrorServerTokenNotFound) {
			controller.log.LogError(operation, err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
			return
		}

		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	balance, err := controller.usecase.GetBalance(r.Context(), tokenID)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	balanceOutput := output.AccountBalanceVO{
		Balance: balance,
	}

	json.NewEncoder(w).Encode(balanceOutput)
}
