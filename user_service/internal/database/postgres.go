package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/olzzhas/edunite-server/user_service/internal/config"
	"log"
	"time"
)

var db *pgxpool.Pool

func ConnectDB(cfg *config.Config) *pgxpool.Pool {
	dbURL := cfg.DatabaseURL

	var pool *pgxpool.Pool
	var err error

	for i := 0; i < 5; i++ {
		pool, err = pgxpool.Connect(context.Background(), dbURL)
		if err == nil {
			log.Println("Connected to database successfully.")
			break
		}

		log.Printf("Failed to connect to database: %v. Retrying in 2 seconds...\n", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Unable to connect to database after retries: %v\n", err)
	}

	db = pool
	return db
}

func GetDB() *pgxpool.Pool {
	return db
}
