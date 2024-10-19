package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

var db *pgxpool.Pool

func ConnectDB() *pgxpool.Pool {
	dbURL := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	db = pool
	return db
}

func GetDB() *pgxpool.Pool {
	return db
}
