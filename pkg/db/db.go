package db

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ochiengotieno304/feedpulse-go/configs"
)

var conn *pgx.Conn

func init() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	conn, err = pgx.Connect(context.Background(), config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Setup graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if conn != nil {
			conn.Close(context.Background())
		}
		os.Exit(0)
	}()
}

// Attempt to reconnect a maximum of 3 times with exponential backoff
func reconnect() (*pgx.Conn, error) {
	const maxRetries = 3
	var conn *pgx.Conn
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		config, cfgErr := configs.LoadConfig()
		if cfgErr != nil {
			log.Printf("Error loading config: %s", cfgErr)
			continue
		}

		conn, err = pgx.Connect(context.Background(), config.DatabaseUrl)
		if err == nil {
			return conn, nil
		}

		log.Printf("Failed to connect to database, attempt %d/%d: %v", attempt, maxRetries, err)
		time.Sleep(time.Duration(attempt) * time.Second) // Exponential backoff could be implemented here
	}

	return nil, err // Return the last error encountered
}

func DB() *pgx.Conn {
	if conn.IsClosed() {
		var err error
		conn, err = reconnect()
		if err != nil {
			log.Fatalf("Unable to reconnect to database: %v\n", err)
		}
	}
	return conn
}
