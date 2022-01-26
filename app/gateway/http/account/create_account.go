package account

import (
	"encoding/json"
	"net/http"

	"stoneBanking/app/gateway/http/account/vo/input"
	validations "stoneBanking/app/gateway/http/account/vo/input/validations"
	"stoneBanking/app/gateway/http/account/vo/output"
	"stoneBanking/app/gateway/http/response"
)

//@Summary Create a account
//@Description With the data received, validate her and if all is correct, and don't exist a account with that document, create a new account
//@Accept json
//@Produce json
//@Param account body input.CreateAccountVO true "Account Creation Data"
//@Success 200 {object} output.AccountOutputVO
//@Failure	400 {object} response.OutputError
//@Failure 500 {object} response.OutputError
//@Router /accounts [POST]
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Account.Create"
	c.log.SetRequestIDFromContext(r.Context())
	resp := response.NewResponse(w)

	c.log.LogInfo(operation, "receive the body and unmarshal the data")
	var accountInput input.CreateAccountVO
	if err := json.NewDecoder(r.Body).Decode(&accountInput); err != nil {
		c.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		resp.BadRequest(response.NewError(err))
		return
	}

	accountInput.CPF = accountInput.CPF.TrimCPF()

	c.log.LogDebug(operation, "begin the validation of the input data")
	errs := validations.ValidateAccountInput(accountInput)
	if len(errs) > 0 {
		resp.BadRequest(response.NewErrors(errs))
		return
	}

	accountData := accountInput.ToEntity()
	account, err := c.usecase.Create(r.Context(), accountData)
	if err != nil {
		c.log.LogError(operation, err.Error())
		resp.InternalError(response.NewError(err))
		return
	}

	accountOutput := output.ToAccountOutput(account)

	resp.Created(accountOutput)
}
