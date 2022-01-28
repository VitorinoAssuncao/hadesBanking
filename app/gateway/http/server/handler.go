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
	m := middleware.NewMiddleware(logger, token)
	router.Use(m.LogRoutes)
	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)
	controller_account := accounts.New(usecase.Accounts, token, logger)
	controller_transfer := transfers.New(usecase.Transfer, token, logger)

	account := router.PathPrefix("/accounts").Subrouter()
	account.HandleFunc("", controller_account.GetAll).Methods("GET")
	account.HandleFunc("", controller_account.Create).Methods("POST")

	balance := account.PathPrefix("/{account_id}/balance").Subrouter()
	balance.Use(m.GetAccountIDFromTokenLogRoutes)
	balance.HandleFunc("", controller_account.GetBalance).Methods("GET")

	login := router.PathPrefix("/login").Subrouter()
	login.HandleFunc("", controller_account.LoginUser).Methods("POST")

	transfer := router.PathPrefix("/transfers").Subrouter()
	transfer.Use(m.GetAccountIDFromTokenLogRoutes)
	transfer.HandleFunc("", controller_transfer.Create).Methods("POST")
	transfer.HandleFunc("", controller_transfer.GetAllByAccountID).Methods("GET")

	// Documentation
	doc := router.PathPrefix("/documentation").Subrouter()
	doc.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
	server := Server{Router: *router}
	return &server
}
