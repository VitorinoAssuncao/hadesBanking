package account

import (
	"net/http"
	"stoneBanking/app/gateway/http/account/vo/output"
	"stoneBanking/app/gateway/http/response"
)

//@Summary Get All Accounts
//@Description Get all accounts actually in the system
//@Produce json
//@Success 200 {object} []output.AccountOutputVO
//@Failure	400 {object} response.OutputError
//@Failure 500 {object} response.OutputError
//@Router /accounts [GET]
func (controller *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Account.GetAll"
	resp := response.NewResponse(w)

	accounts, err := controller.usecase.GetAll(r.Context())
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.InternalError(response.NewError(err))
		return
	}

	accountsOutput := output.ToAccountsOutput(accounts)
	controller.log.LogInfo(operation, "accounts created successfully")
	resp.Ok(accountsOutput)
}
