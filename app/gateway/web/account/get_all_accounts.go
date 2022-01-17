package account

import (
	"encoding/json"
	"net/http"
	"stoneBanking/app/gateway/web/account/vo/output"
)

//@Sumary Get All Accounts
//@Description Get all accounts actually in the system
//@Produce json
//@Sucess 200 {object} []output.AccountOutputVO
//@Router /accounts [get]
func (controller *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := controller.usecase.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode([]output.OutputError{{Error: err.Error()}})
		return
	}

	var accountsOutput = make([]output.AccountOutputVO, 0)
	for _, account := range accounts {
		tempAccount := output.AccountToOutput(account)
		accountsOutput = append(accountsOutput, tempAccount)
	}

	json.NewEncoder(w).Encode(accountsOutput)

}
