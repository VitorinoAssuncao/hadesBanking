package account

import (
	"encoding/json"
	"net/http"
	"stoneBanking/app/gateway/web/account/vo/output"
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

	accounts, err := controller.usecase.GetAll(r.Context())
	if err != nil {
		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode([]output.OutputError{{Error: err.Error()}}) //nolint: errorlint
		return
	}

	var accountsOutput = make([]output.AccountOutputVO, 0)
	for _, account := range accounts {
		tempAccount := output.AccountToOutput(account)
		accountsOutput = append(accountsOutput, tempAccount)
	}

	controller.log.LogInfo(operation, "accounts created sucessfully")
	json.NewEncoder(w).Encode(accountsOutput) //nolint: errorlint

}
