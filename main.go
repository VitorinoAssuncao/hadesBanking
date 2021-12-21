package main

import (
	"fmt"
	"stoneBanking/app/gateway/database/postgres"
	"stoneBanking/app/gateway/web/server"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	postgres.InitiliazeDatabase()
	db := postgres.RetrieveConnection()
	repository := server.NewPostgresRepositoryWrapper(db)

	workspaces := server.NewUseCaseWrapper(repository)

	server.New(workspaces)
}
