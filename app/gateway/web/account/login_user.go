package accounts

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/common/utils"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/gateway/web/account/vo/input"
	validations "stoneBanking/app/gateway/web/account/vo/input/validations"
	"stoneBanking/app/gateway/web/account/vo/output"
)

func (controller *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData input.LoginVO
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &loginData)

	loginData.CPF = utils.TrimCPF(loginData.CPF)
	err = validations.ValidateLoginInputData(loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	account := account.Account{
		CPF:    loginData.CPF,
		Secret: loginData.Secret,
	}

	token, err := controller.usecase.LoginUser(context.Background(), account)
	loginOutput := output.LoginOutputVO{
		Token: token,
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(loginOutput)
}
