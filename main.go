package main

import (
	"fmt"
	"stoneBanking/app/gateway/database/postgres"
	"stoneBanking/app/gateway/web/server"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Não foi possível carregar as varíaveis de ambiente")
	}

	postgres.InitiliazeDatabase()
	db := postgres.RetrieveConnection()
	repository := server.NewPostgresRepositoryWrapper(db)

	workspaces := server.NewUseCaseWrapper(repository)

	server.New(workspaces)
}
