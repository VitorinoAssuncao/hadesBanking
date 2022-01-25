package main

import (
	"fmt"
	"log"
	"stoneBanking/app/common/utils/config"
	"stoneBanking/app/gateway/database/postgres"
	"stoneBanking/app/gateway/web/server"

	"github.com/joho/godotenv"
)

// @title           stoneBanking API
// @version         1.0
// @description     This is a simples application to create accounts and transfers between valide accounts
// @contact.name   API Support
// @contact.email  vitorinomuller@gmail.com
// @License.name Stone Co®.
// @host      localhost:8000

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("has not possible to load the env file")
	}

	// Load the env variables
	cfg := config.LoadConfig()

	// Initialize the database and return him to a variable
	postgres.InitiliazeDatabase(cfg)
	db := postgres.RetrieveConnection()

	// Create the repositories and usecase repositories
	repository := server.NewPostgresRepositoryWrapper(db, cfg.SigningKey)
	workspaces := server.NewUseCaseWrapper(repository)

	// Initialize the server and host him in localhost:8000
	server.New(workspaces, repository.Token, repository.Log)
}
