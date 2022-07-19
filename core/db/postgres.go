package db

import (
	"log"
	"rohandhamapurkar/code-executor/core/config"
	"rohandhamapurkar/code-executor/core/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB

func initPostgresDBConn() {
	var err error
	// Connect to database
	Postgres, err = gorm.Open(postgres.Open(config.PostgresDsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to PostgreSQL.")
		initPostgresSchema()
	}

}

func initPostgresSchema() {
	Postgres.AutoMigrate(&models.Snippets{})
	log.Println("Initialized Postgres schema.")
}
