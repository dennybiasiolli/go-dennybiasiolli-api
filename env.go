package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var HTTP_LISTEN string

func getEnvVariables() {
	/*
		loading .env files in this order, if a variable is not set in `.env`,
		it's read from `.env.default`
	*/
	errEnv := godotenv.Load(".env")
	errEnvDefault := godotenv.Load(".env.default")
	if errEnvDefault != nil && errEnv != nil {
		log.Fatal("Error loading .env.default or .env file")
	}

	HTTP_LISTEN = os.Getenv("HTTP_LISTEN")
}
