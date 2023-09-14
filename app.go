package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"test2/webhook"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	authKey := os.Getenv("AUTH_KEY")
	app := webhook.New(authKey)
	app.Listen(":3000")
}
