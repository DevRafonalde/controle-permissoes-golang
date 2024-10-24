package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnection() *pgxpool.Pool {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Monta a URL de conexão
	databaseUrl := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":5432/" + dbName + "?sslmode=disable"

	// Conecta ao banco de dados usando pgxpool
	conn, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatal("falha ao se conectar a DB: ", err)
	}

	// Verifica a conexão
	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal("falha ao verificar a conexão com a DB: ", err)
	}

	return conn
}
