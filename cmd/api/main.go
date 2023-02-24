package main

import (
	"database/sql"
	"log"
	"weehongayden/finance-flow-app/internal/config"
	"weehongayden/finance-flow-app/internal/database"
	"weehongayden/finance-flow-app/internal/logging"
)

type App struct {
	log *log.Logger
	db  *sql.DB
}

func main() {
	logger, err := logging.Init()
	if err != nil {
		logger.Fatalf("Failed to load log file: %v", err)
	}

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	_, err = database.New(c)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// app := &App{
	// 	db: db,
	// }

	serve(c)
}
