package account

import (
	"encoding/json"
	"net/http"
	"stoneBanking/app/gateway/web/account/vo/output"
)

func (controller *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := controller.usecase.GetAll(r.Context())
	var accountsOutput = make([]output.AccountOutputVO, 0)

	for _, account := range accounts {
		tempAccount := output.AccountToOutput(account)
		accountsOutput = append(accountsOutput, tempAccount)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(accountsOutput)

}
