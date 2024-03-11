package connections

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading env file", err)
	}
}
