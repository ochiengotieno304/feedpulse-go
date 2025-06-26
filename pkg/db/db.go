package db

import (
	"github.com/ochiengotieno304/feedpulse-go/configs"
	"github.com/ochiengotieno304/feedpulse-go/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config, err := configs.LoadConfig()
	utils.ErrorHandler(err)

	db, err = gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	utils.ErrorHandler(err)
}

func DB() *gorm.DB {
	return db
}
