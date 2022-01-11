package accounts

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/web/account/vo/input"
	validations "stoneBanking/app/gateway/web/account/vo/input/validations"
	"stoneBanking/app/gateway/web/account/vo/output"
)

func (controller *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData input.CreateAccountVO
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &loginData)

	err = validations.ValidateLoginInputData(loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account := account.Account{
		CPF:    types.Document(loginData.CPF).TrimCPF(),
		Secret: types.Password(loginData.Secret),
	}

	token, err := controller.usecase.LoginUser(context.Background(), account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loginOutput := output.LoginOutputVO{
		Token: token,
	}

	json.NewEncoder(w).Encode(loginOutput)
}
