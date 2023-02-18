package main

import (
	"database/sql"
	"fmt"
	"log"
	"weehongayden/finance-flow-app/internal/config"
	"weehongayden/finance-flow-app/internal/database"
)

type App struct {
	db *sql.DB
}

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.New(c)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := &App{
		db: db,
	}

	fmt.Println(app)
}
