package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sub-rat/MorningContactApi/internals/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	s := server.GetServer()
	s.Run()
}
