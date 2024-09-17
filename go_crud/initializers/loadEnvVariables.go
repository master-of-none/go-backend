package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading the file")
	}
}
