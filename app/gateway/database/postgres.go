package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

var database *pgx.Conn

func InitiliazeDatabase() {
	const user = "postgres"
	const pass = "postgres"
	const base = "bank_stone"

	connection, _ := pgx.Connect(context.Background(), "postgres://"+user+":"+pass+"@0.0.0.0:5432/"+base+"")
	database = connection
}

func RetrieveConnection() *pgx.Conn {
	return database
}
