package accounts

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/application/vo/input"
)

func (controller *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var accountData = &input.CreateAccountVO{}
	//var connection = database_connector.RetrieveConnection()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &accountData)

	account_output, err := controller.usecase.Create(r.Context(), *accountData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(account_output)
}
