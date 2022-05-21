package environment

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port       string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
)

func GetEnvironmentConfig() {
	if err := godotenv.Load(".env"); err != nil {
		errMessage := fmt.Sprintf("Failed to load .env %s", err.Error())
		log.Fatal(errMessage)
	}

	fmt.Println("Loading .env")

	Port = os.Getenv("PORT")
	if Port == "" {
		fmt.Println("ERROR PORT")
	} else {
		fmt.Println("SUCCESS PORT")
	}

	DbUser = os.Getenv("DB_USER")
	if DbUser == "" {
		fmt.Println("ERROR DB_USER")
	} else {
		fmt.Println("SUCCESS DB_USER")
	}

	DbPassword = os.Getenv("DB_PASSWORD")
	if DbPassword == "" {
		fmt.Println("ERROR DB_PASSWORD")
	} else {
		fmt.Println("SUCCESS DB_PASSWORD")
	}

	DbHost = os.Getenv("DB_HOST")
	if DbHost == "" {
		fmt.Println("ERROR DB_HOST")
	} else {
		fmt.Println("SUCCESS DB_HOST")
	}

	DbPort = os.Getenv("DB_PORT")
	if DbPort == "" {
		fmt.Println("ERROR DB_PORT")
	} else {
		fmt.Println("SUCCESS DB_PORT")
	}

	DbName = os.Getenv("DB_NAME")
	if DbName == "" {
		fmt.Println("ERROR DB_NAME")
	} else {
		fmt.Println("SUCCESS DB_NAME")
	}

	fmt.Println("Loaded .env")
}
