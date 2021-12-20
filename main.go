package main

import (
	"fmt"
	"stoneBanking/app/gateway/database"
	"stoneBanking/app/gateway/web/server"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	database.InitiliazeDatabase()
	connection := database.RetrieveConnection()
	repos := server.NewPostgresRepositoryWrapper(connection)

	workspaces := server.NewUseCaseWrapper(repos)

	server.New(workspaces)
}
