package common

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

var SEND_EMAIL_AFTER_CITAZIONE_ADDED bool
var EMAIL_HOST string
var EMAIL_PORT string
var EMAIL_HOST_USER string
var EMAIL_HOST_PASSWORD string
var EMAIL_DEFAULT_FROM string
var EMAIL_DEFAULT_TO string

var JWT_HMAC_SAMPLE_SECRET string

func GetEnvVariables() {
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

	SEND_EMAIL_AFTER_CITAZIONE_ADDED = os.Getenv("SEND_EMAIL_AFTER_CITAZIONE_ADDED") == "1"
	EMAIL_HOST = os.Getenv("EMAIL_HOST")
	EMAIL_PORT = os.Getenv("EMAIL_PORT")
	EMAIL_HOST_USER = os.Getenv("EMAIL_HOST_USER")
	EMAIL_HOST_PASSWORD = os.Getenv("EMAIL_HOST_PASSWORD")
	EMAIL_DEFAULT_FROM = os.Getenv("EMAIL_DEFAULT_FROM")
	EMAIL_DEFAULT_TO = os.Getenv("EMAIL_DEFAULT_TO")

	JWT_HMAC_SAMPLE_SECRET = os.Getenv("JWT_HMAC_SAMPLE_SECRET")
}
