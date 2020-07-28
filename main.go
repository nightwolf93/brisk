package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nightwolf93/brisk/api"
	"github.com/nightwolf93/brisk/auth"
	"github.com/nightwolf93/brisk/storage"
)

// version is the string ref of the version of brisk
var version = "0.0.1"

func main() {
	log.Printf("Brisk v%s started", version)
	err := godotenv.Load()
	if err != nil {
		log.Print("can't find .env file, using default os env")
	}

	storage.Open()
	auth.InitMasterPair()
	api.Init()
}
