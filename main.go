package main

import (
	"log"

	"github.com/mini-eggs/home-server/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.Start()
}
