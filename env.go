package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var HTTP_LISTEN string

var PG_HOST string
var PG_PORT string
var PG_USER string
var PG_PASSWORD string
var PG_SSLMODE string
var PG_DATABASE string

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

	PG_HOST = os.Getenv("PG_HOST")
	PG_PORT = os.Getenv("PG_PORT")
	PG_USER = os.Getenv("PG_USER")
	PG_PASSWORD = os.Getenv("PG_PASSWORD")
	PG_SSLMODE = os.Getenv("PG_SSLMODE")
	PG_DATABASE = os.Getenv("PG_DATABASE")
}
