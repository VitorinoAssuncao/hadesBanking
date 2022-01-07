package accounts

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/common/utils"
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
	}

	accountInput.CPF = utils.TrimCPF(accountInput.CPF)

	accountInput, err = validations.ValidateAccountInput(accountInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	accountData := input.GenerateAccount(accountInput)
	account, err := controller.usecase.Create(r.Context(), accountData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountOutput := output.AccountToOutput(account)

	json.NewEncoder(w).Encode(accountOutput)
}
