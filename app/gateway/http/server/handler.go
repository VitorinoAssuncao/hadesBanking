package server

import (
	"log"
	"net/http"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	accounts "stoneBanking/app/gateway/http/account"
	"stoneBanking/app/gateway/http/middleware"
	transfers "stoneBanking/app/gateway/http/transfer"

	_ "stoneBanking/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Router mux.Router
}

func New(usecase *UseCaseWrapper, token token.Authenticator, logger logHelper.Logger) *Server {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.LogRoutes)
	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)
	controller_account := accounts.New(usecase.Accounts, token, logger)
	controller_transfer := transfers.New(usecase.Transfer, token, logger)
	router.HandleFunc("/account", controller_account.Create).Methods("POST")
	router.HandleFunc("/account/login", controller_account.LoginUser).Methods("POST")
	router.HandleFunc("/accounts", controller_account.GetAll).Methods("GET")
	router.HandleFunc("/account/balance", controller_account.GetBalance).Methods("GET")
	router.HandleFunc("/transfer", controller_transfer.Create).Methods("POST")
	router.HandleFunc("/transfer", controller_transfer.GetAllByAccountID).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
	server := Server{Router: *router}
	return &server
}
