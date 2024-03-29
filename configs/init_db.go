package configs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	config, errr := LoadConfig()

	if errr != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	DB, err = gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}
}