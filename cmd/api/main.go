package main

import (
	"fmt"
	"log"
	"weehongayden/finance-flow-app/internal/config"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Println(c)
}
