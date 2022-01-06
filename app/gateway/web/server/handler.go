package server

import (
	"log"
	"net/http"
	accounts "stoneBanking/app/gateway/web/account"
	transfers "stoneBanking/app/gateway/web/transfer"

	"github.com/gorilla/mux"
)

type Server struct {
	Router mux.Router
}

func New(usecase *UseCaseWrapper) *Server {
	router := mux.NewRouter().StrictSlash(true)
	controller_account := accounts.New(usecase.Accounts)
	controller_transfer := transfers.New(usecase.Transfer)
	router.HandleFunc("/account", controller_account.Create).Methods("POST")
	router.HandleFunc("/account/login", controller_account.LoginUser).Methods("POST")
	router.HandleFunc("/accounts", controller_account.GetAll).Methods("GET")
	router.HandleFunc("/account/{user_id}/balance", controller_account.GetBalance).Methods("GET")
	router.HandleFunc("/transfer", controller_transfer.Create).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
	server := Server{Router: *router}
	return &server
}
