package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	db_url := os.Getenv("DB_URL")
	var dberr error
	Pool, dberr = pgxpool.New(context.Background(), db_url)
	if dberr != nil {
		log.Println(dberr)
	}
	log.Println("DB is connected")
	query := `
	CREATE TABLE IF NOT EXISTS messages (
		id UUID PRIMARY KEY,
		content TEXT NOT NULL,
		timestamp BIGINT NOT NULL,
		status VARCHAR(10) NOT NULL
	);`
	Pool.Exec(context.Background(), query)
}
