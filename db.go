package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDb() {
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v dbname=%v sslmode=%v",
		PG_HOST, PG_PORT, PG_USER, PG_DATABASE, PG_SSLMODE,
	)
	if PG_PASSWORD != "" {
		dsn = fmt.Sprint(dsn+" password=%v", PG_PASSWORD)
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the tatabase")
	}
}
