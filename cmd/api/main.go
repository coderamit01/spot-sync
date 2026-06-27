package main

import (
	"spotsync/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDB()
}
