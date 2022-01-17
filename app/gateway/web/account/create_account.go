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

//@Sumary Create a account
//@Description With the data received, validate her and if all is correct, and dont exist a account with that document, create a new account
//@Accept json
//@Produce json
//@Param account body input.CreateAccountVO true "Account Creation Data"
//@Success 200 {object} output.AccountOutputVO
//@Failure	400 {object} output.OutputError
//@Failure 500 {object} output.OutputError
//@Router /account [post]
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
