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
	filename := os.Getenv("TASKS_FILE")
	app := webhook.New(filename, authKey)
	_ = app.Listen(":3000")
}
