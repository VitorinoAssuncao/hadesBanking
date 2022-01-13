package account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/web/account/vo/input"
	validations "stoneBanking/app/gateway/web/account/vo/input/validations"
	"stoneBanking/app/gateway/web/account/vo/output"
)

func (controller *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var accountInput = input.CreateAccountVO{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &accountInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountInput.CPF = types.Document(accountInput.CPF.TrimCPF())

	accountInput, err = validations.ValidateAccountInput(accountInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountData := accountInput.GenerateAccount()
	account, err := controller.usecase.Create(r.Context(), accountData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountOutput := output.AccountToOutput(account)

	json.NewEncoder(w).Encode(accountOutput)
}
