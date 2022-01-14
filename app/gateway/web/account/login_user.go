package account

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/web/account/vo/input"
	validations "stoneBanking/app/gateway/web/account/vo/input/validations"
	"stoneBanking/app/gateway/web/account/vo/output"
)

func (controller *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData input.CreateAccountVO
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	json.Unmarshal(reqBody, &loginData)

	err = validations.ValidateLoginInputData(loginData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	account := account.Account{
		CPF:    types.Document(loginData.CPF).TrimCPF(),
		Secret: types.Password(loginData.Secret),
	}

	token, err := controller.usecase.LoginUser(context.Background(), account)
	if err != nil {
		if err == customError.ErrorAccountTokenGeneration {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(output.OutputError{Error: err.Error()})
		return
	}

	loginOutput := output.LoginOutputVO{
		Authorization: token,
	}

	json.NewEncoder(w).Encode(loginOutput)
}
