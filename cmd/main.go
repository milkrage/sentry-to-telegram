package main

import (
	"log"

	"github.com/milkrage/sentry-to-telegram/internal/app/config"
)

func main() {
	_, err := config.New("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}
