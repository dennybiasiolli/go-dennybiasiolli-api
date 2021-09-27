package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v dbname=%v sslmode=%v",
		PG_HOST, PG_PORT, PG_USER, PG_DATABASE, PG_SSLMODE,
	)
	if PG_PASSWORD != "" {
		dsn = fmt.Sprint(dsn+" password=%v", PG_PASSWORD)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the tatabase")
	}

	return db
}

func initDb() {
	db := connectDb()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error accessing the tatabase")
	}
	sqlDB.Close()
}
