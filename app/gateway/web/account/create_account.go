package accounts

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/application/vo/input"
)

func (controller *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var accountData = &input.CreateAccountVO{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &accountData)

	accountOutput, err := controller.usecase.Create(r.Context(), *accountData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(accountOutput)
}
