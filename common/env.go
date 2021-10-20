package common

import (
	"log"
	"os"
	"strconv"

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
var JWT_ACCESS_TOKEN_LIFETIME_SECONDS int = 3600

var GOOGLE_OAUTH2_CLIENT_ID string
var GOOGLE_OAUTH2_CLIENT_SECRET string
var GOOGLE_OAUTH2_DEFAULT_REDIRECT_URL string

func GetEnvVariables(mainFile string, fallbackFile string) {
	/*
		loading .env files in this order, if a variable is not set in `mainFile`,
		it's read from `fallbackFile`
	*/
	errEnv := godotenv.Load(mainFile)
	errEnvDefault := godotenv.Load(fallbackFile)
	if errEnvDefault != nil && errEnv != nil {
		log.Fatalf("Error loading %s or %s file", mainFile, fallbackFile)
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
	if val, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_LIFETIME_SECONDS")); err != nil {
		JWT_ACCESS_TOKEN_LIFETIME_SECONDS = val
	}

	GOOGLE_OAUTH2_CLIENT_ID = os.Getenv("GOOGLE_OAUTH2_CLIENT_ID")
	GOOGLE_OAUTH2_CLIENT_SECRET = os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET")
	GOOGLE_OAUTH2_DEFAULT_REDIRECT_URL = os.Getenv("GOOGLE_OAUTH2_DEFAULT_REDIRECT_URL")
}
