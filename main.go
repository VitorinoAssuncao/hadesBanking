package main

import (
	"fmt"
	"stoneBanking/app/gateway/database/postgres"
	"stoneBanking/app/gateway/web/server"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	postgres.InitiliazeDatabase()
	connection := postgres.RetrieveConnection()
	repository := server.NewPostgresRepositoryWrapper(connection)

	workspaces := server.NewUseCaseWrapper(repository)

	server.New(workspaces)
}
