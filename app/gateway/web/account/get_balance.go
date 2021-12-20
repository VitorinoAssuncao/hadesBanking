package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (controller *Controller) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["user_id"]
	balance_output, err := controller.usecase.GetBalance(r.Context(), accountId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(balance_output)

}
