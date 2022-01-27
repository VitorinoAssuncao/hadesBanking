package account

import (
	"net/http"
	"stoneBanking/app/gateway/http/account/vo/output"
	"stoneBanking/app/gateway/http/response"
)

//@Sumary Get All Accounts
//@Description Get all accounts actually in the system
//@Produce json
//@Success 200 {object} []output.AccountOutputVO
//@Failure	400 {object} output.OutputError
//@Failure 500 {object} output.OutputError
//@Router /accounts [GET]
func (controller *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Account.GetAll"
	resp := response.CustomResponse{
		Writer: w,
	}

	accounts, err := controller.usecase.GetAll(r.Context())
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.InternalError(output.OutputError{Error: err.Error()})
		return
	}

	accountsOutput := output.ToAccountsOutput(accounts)
	controller.log.LogInfo(operation, "accounts created sucessfully")
	resp.Ok(accountsOutput)
}
