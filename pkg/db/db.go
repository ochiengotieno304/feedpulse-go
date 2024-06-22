package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/ochiengotieno304/feedpulse-go/configs"
)

func DB() *pgx.Conn {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	

	conn, err := pgx.Connect(context.Background(), config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return conn
}
