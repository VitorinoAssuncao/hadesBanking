package account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/http/account/vo/input"
	validations "stoneBanking/app/gateway/http/account/vo/input/validations"
	"stoneBanking/app/gateway/http/account/vo/output"
)

//@Sumary Create a account
//@Description With the data received, validate her and if all is correct, and dont exist a account with that document, create a new account
//@Accept json
//@Produce json
//@Param account body input.CreateAccountVO true "Account Creation Data"
//@Success 200 {object} output.AccountOutputVO
//@Failure	400 {object} output.OutputError
//@Failure 500 {object} output.OutputError
//@Router /account [POST]
func (controller *Controller) Create(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Account.Create"

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()}) //nolint: errorlint
		return
	}

	controller.log.LogInfo(operation, "unmarshal the data to a internal object")
	var accountInput = input.CreateAccountVO{}
	err = json.Unmarshal(reqBody, &accountInput)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()}) //nolint: errorlint
		return
	}

	accountInput.CPF = types.Document(accountInput.CPF.TrimCPF())

	controller.log.LogInfo(operation, "begin the validation of the input data")
	accountInput, err = validations.ValidateAccountInput(accountInput)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()}) //nolint: errorlint
		return
	}

	accountData := accountInput.GenerateAccount()
	account, err := controller.usecase.Create(r.Context(), accountData)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()}) //nolint: errorlint
		return
	}

	accountOutput := output.AccountToOutput(account)

	json.NewEncoder(w).Encode(accountOutput) //nolint: errorlint
}
