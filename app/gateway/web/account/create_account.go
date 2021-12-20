package accounts

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/application/vo/input"
)

func (controller *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var account_data = &input.CreateAccountVO{}
	//var connection = database_connector.RetrieveConnection()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &account_data)

	account_output, err := controller.usecase.Create(r.Context(), *account_data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(account_output)

}
