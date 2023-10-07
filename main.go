package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := loadEnvs()
	if err != nil {
		log.Fatal("Error while loading .env file: ", err)
	}

	portString := os.Getenv("PORT")

	fmt.Println("Hello world!:", portString)
}

func loadEnvs() error {
	envName := os.Getenv("ENV")
	if envName == "" {
		envName = "development"
	}

	godotenv.Load(".env." + envName)
	err := godotenv.Load(".env")
	return err
}
