package account

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/http/account/vo/output"
	"stoneBanking/app/gateway/http/middleware"
	"stoneBanking/app/gateway/http/response"
)

//@Summary Get the balance of a account
//@Description With a authorization token valid, return the balance of a account
//@Produce json
//@Param authorization header string true "Authorization Token"
//@Success 200 {object} output.AccountBalanceVO
//@Failure	400 {object} response.OutputError
//@Failure 500 {object} response.OutputError
//@Router /accounts/{account_id}/balance [GET]
func (c *Controller) GetBalance(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Account.GetBalance"
	resp := response.NewResponse(w)

	c.log.LogInfo(operation, "take the value from the token")
	accountIDToken, err := middleware.GetAccountIDFromContext(r.Context())
	if err != nil {
		c.log.LogError(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	vars := mux.Vars(r)
	accountIDRoute := vars["account_id"]

	if accountIDToken != accountIDRoute {
		resp.Unauthorized(response.NewError(err))
		return
	}

	balance, err := c.usecase.GetBalance(r.Context(), accountIDToken)
	if err != nil {
		if errors.Is(err, customError.ErrorAccountIDNotFound) {
			c.log.LogError(operation, err.Error())
			resp.BadRequest(response.NewError(err))
			return
		}

		c.log.LogError(operation, err.Error())
		resp.InternalError(response.NewError(err))
		return
	}

	balanceOutput := output.AccountBalanceVO{
		Balance: balance.ToFloat(),
	}

	resp.Ok(balanceOutput)
}
