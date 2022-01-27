package account

import (
	"errors"
	"net/http"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/http/account/vo/output"
	"stoneBanking/app/gateway/http/middleware"
	"stoneBanking/app/gateway/http/response"
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
	resp := response.CustomResponse{Writer: w}

	controller.log.LogInfo(operation, "take the value from the token")
	tokenID, err := middleware.GetAccountIDFromToken(r, controller.token)
	if err != nil {
		if errors.Is(err, customError.ErrorServerTokenNotFound) {
			controller.log.LogError(operation, err.Error())
			resp.Unauthorized(output.OutputError{Error: err.Error()})
			return
		}

		controller.log.LogError(operation, err.Error())
		resp.InternalError(output.OutputError{Error: err.Error()})
	}

	balance, err := controller.usecase.GetBalance(r.Context(), tokenID)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.InternalError(output.OutputError{Error: err.Error()})
		return
	}

	balanceOutput := output.AccountBalanceVO{
		Balance: balance,
	}

	resp.Ok(balanceOutput)
}
