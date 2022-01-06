package transfer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stoneBanking/app/application/vo/input"
)

func (controller Controller) Create(w http.ResponseWriter, r *http.Request) {
	var transferData = input.CreateTransferVO{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &transferData)

	transferOutput, err := controller.usecase.Create(context.Background(), transferData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transferOutput)
}
