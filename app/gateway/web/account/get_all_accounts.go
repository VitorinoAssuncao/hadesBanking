package accounts

import (
	"encoding/json"
	"net/http"
)

func (controller *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	accountsOutput, err := controller.usecase.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(accountsOutput)

}
