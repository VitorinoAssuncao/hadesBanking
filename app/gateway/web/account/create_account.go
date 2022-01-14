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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	err = json.Unmarshal(reqBody, &accountInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	accountInput.CPF = types.Document(accountInput.CPF.TrimCPF())

	accountInput, err = validations.ValidateAccountInput(accountInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	accountData := accountInput.GenerateAccount()
	account, err := controller.usecase.Create(r.Context(), accountData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	accountOutput := output.AccountToOutput(account)

	json.NewEncoder(w).Encode(accountOutput)
}
