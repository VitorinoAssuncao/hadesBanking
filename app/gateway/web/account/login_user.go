package accounts

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/application/vo/input"
)

func (controller *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData input.LoginVO
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &loginData)

	loginOutput, err := controller.usecase.LoginUser(context.Background(), loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(loginOutput)
}
