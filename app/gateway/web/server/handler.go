package server

import (
	"log"
	"net/http"
	accounts "stoneBanking/app/gateway/web/account"

	"github.com/gorilla/mux"
)

type Server struct {
	Router mux.Router
}

func New(usecase *UseCaseWrapper) *Server {
	router := mux.NewRouter().StrictSlash(true)
	controller := accounts.New(usecase.Accounts)
	router.HandleFunc("/account", controller.Create).Methods("POST")
	router.HandleFunc("/account/{user_id}/balance", controller.GetBalance).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
	server := Server{Router: *router}
	return &server
}
