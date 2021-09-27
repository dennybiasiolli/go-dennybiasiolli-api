package main

import (
	"time"
)

type Articolo struct {
	ID                int       `gorm:"AUTO_INCREMENT; primaryKey"`
	DataPubblicazione time.Time ``
	TitoloIt          string    `gorm:"not null; size:255"`
	TitoloEn          string    `gorm:"size:255"`
	TestoIt           string    `gorm:"not null;type:text[]"`
	TestoEn           string    `gorm:"type:text[]"`
	AutoreId          int       `gorm:"not null"`
}

// TableName set Articolo's table name to be `news_articolo`
func (Articolo) TableName() string {
	return "news_articolo"
}
